package user

import (
	"bytes"
	"fmt"
	"html/template"
	"strconv"
	"time"

	"server/core/entity"
	"server/core/errors"
	"server/core/infra/action"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"
	"server/infrastructure/logger"

	"github.com/google/uuid"
)

type BookUsecase struct {
	bookQuery  queryservice.IBookQueryService
	bookRepo   repository.IBookRepository
	bookAPI    repository.IBookAPIRepository
	mailAction action.ISendMailAction
	storeQuery queryservice.IStoreQueryService
}

func NewBookUsecase(
	bookQuery queryservice.IBookQueryService,
	bookRepo repository.IBookRepository,
	bookAPI repository.IBookAPIRepository,
	mailAction action.ISendMailAction,
	storeQuery queryservice.IStoreQueryService,
) *BookUsecase {
	return &BookUsecase{
		bookQuery:  bookQuery,
		bookRepo:   bookRepo,
		bookAPI:    bookAPI,
		mailAction: mailAction,
		storeQuery: storeQuery,
	}
}

func (u *BookUsecase) GetMyBook(userID uuid.UUID) ([]*entity.Booking, *errors.DomainError) {
	books, err := u.bookQuery.GetMyBooking(userID, time.Now())
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return books, nil
}

func (u *BookUsecase) Cancel(bookID uuid.UUID) *errors.DomainError {
	book, err := u.bookQuery.GetByID(bookID)
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}

	if book == nil {
		return errors.NewDomainError(errors.QueryDataNotFoundError, "該当の予約が存在しません。")
	}
	// bookID : CCYYMMDD+9桁連番（0埋め、データ毎に+1）
	TlDataID, err := u.bookQuery.GenerateBookDataID()
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}
	if TlDataID == nil || *TlDataID == "" {
		return errors.NewDomainError(errors.QueryError, "予約番号の生成に失敗しました。")
	}

	domainError, err := u.bookAPI.Cancel(book, *TlDataID)
	if err != nil {
		return errors.NewDomainError(errors.CancelButNeedFeedBack, "キャンセル処理がAPIレベルで失敗しました。")
	}
	if domainError != nil {
		return domainError
	}

	err = u.bookRepo.Delete(bookID)
	if err != nil {
		logger.Errorf("キャンセル処理がDBレベルで失敗しました。%s", bookID.String())
		return errors.NewDomainError(errors.CancelButNeedFeedBack, "キャンセル処理は成功しましたが、DBの削除に失敗しました。")
	}
	return nil
}

func (u *BookUsecase) GetByID(bookID uuid.UUID) (*entity.Booking, *errors.DomainError) {
	book, err := u.bookQuery.GetByID(bookID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	if book == nil {
		return nil, errors.NewDomainError(errors.QueryError, "該当の予約が存在しません。")
	}
	return book, nil
}

func (u *BookUsecase) Reserve(
	stayFrom time.Time,
	stayTo time.Time,
	adult uint,
	child uint,
	roomCount uint,
	CheckInTime entity.CheckInTime,
	TotalCost uint,
	GuestData *entity.GuestData,
	BookPlan *entity.Plan,
	BookUserID uuid.UUID,
	Note string,
) *errors.DomainError {

	store, err := u.storeQuery.GetStayableByID(BookPlan.StoreID)
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}
	if store == nil {
		return errors.NewDomainError(errors.QueryDataNotFoundError, "宿泊施設の情報が取得できませんでした。")
	}

	// bookID : CCYYMMDD+9桁連番（0埋め、データ毎に+1）
	TlDataID, err := u.bookQuery.GenerateBookDataID()
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}
	if TlDataID == nil || *TlDataID == "" {
		return errors.NewDomainError(errors.QueryError, "予約番号の生成に失敗しました。")
	}

	newBook, domainErr := entity.CreateBooking(
		stayFrom,
		stayTo,
		adult,
		child,
		roomCount,
		CheckInTime,
		TotalCost,
		GuestData,
		BookPlan,
		BookUserID,
		Note,
		*TlDataID,
		nil)
	if domainErr != nil {
		return domainErr
	}

	tlnumber, domainError, err := u.bookAPI.Reserve(newBook)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	if domainError != nil {
		return domainError
	}

	if tlnumber == nil || *tlnumber == "" {
		return errors.NewDomainError(errors.CancelButNeedFeedBack, "予約は完了しましたが、予約番号が取得できずDB保存に失敗しました。")
	}

	newBook.TlBookingNumber = tlnumber

	err = u.bookRepo.Save(newBook)
	if err != nil {
		logger.Error("予約情報の保存に失敗しました。" + err.Error())
		return errors.NewDomainError(errors.CancelButNeedFeedBack, "予約は完了しましたが、情報の保存に失敗しました。")
	}

	// 予約完了メールの内容を取得
	content, err := reserveMailContent(newBook, store)
	if err != nil {
		return errors.NewDomainError(errors.CancelButNeedFeedBack, "予約は完了しましたが、お客様へのメール作成に失敗しました。")
	}

	// 予約完了メール送信
	err = u.mailAction.Send(GuestData.Mail, "【"+store.Name+*store.BranchName+"】宿泊予約完了のお知らせ", *content)
	if err != nil {
		return errors.NewDomainError(errors.CancelButNeedFeedBack, "予約は完了しましたが、お客様へのメール送信に失敗しました。")
	}
	return nil
}

func reserveMailContent(
	bookinfo *entity.Booking,
	store *entity.StayableStore,
) (*string, error) {
	contentTemplate := `
{{.GuestName}} 様

この度は、当ホテルをご予約いただき、誠にありがとうございます。
以下の内容でご予約が完了しましたので、お知らせいたします。

予約番号: {{.ReservationNumber}}
チェックイン日: {{.CheckInDate}}
チェックアウト日: {{.CheckOutDate}}
チェックイン時間：{{.CheckInTime}}
宿泊プラン: {{.ReservationPlan}}
宿泊人数: {{.NumberOfGuests}}名様
部屋数：{{.RoomCount}}部屋
部屋タイプ：{{.RoomType}}
食事タイプ：{{.MealType}}
禁煙喫煙：{{.SmokingType}}

合計金額: {{.TotalAmount}}円

ご要望: {{.note}}


【ご注意事項】
※チェックイン時間
変更がある場合は、必ず宿泊ホテルまでご連絡ください。

※駐車場について
駐車場は完全予約制です。
宿泊ホテルまで事前にご連絡下さい。
ご連絡がない場合はご利用いただけない場合がございます。

※乳幼児について
乳幼児の添い寝には別途料金を頂戴しております。
事前に施設までご連絡ください。

ご不明な点がございましたら、お気軽にお問い合わせください。
当日は、お客様のお越しを心よりお待ちしております。

敬具

{{.HotelName}}
{{.HotelAddress}}
{{.HotelPhone}}
{{.HotelURL}}
`

	// メールのデータを定義（実際のデータはアプリケーションから取得）
	var people string
	if bookinfo.Child > 0 {
		people = "大人: " + strconv.FormatUint(uint64(bookinfo.Adult), 10) + "名様／子ども：" + strconv.FormatUint(uint64(bookinfo.Child), 10)
	} else {
		people = "大人: " + strconv.FormatUint(uint64(bookinfo.Adult), 10) + "名様"
	}
	data := map[string]string{
		"GuestName":         bookinfo.GuestData.LastName + bookinfo.GuestData.FirstName,
		"ReservationNumber": bookinfo.TlDataID,
		"CheckInDate":       bookinfo.StayFrom.Format("2006年1月2日"),
		"CheckOutDate":      bookinfo.StayTo.Format("2006年1月2日"),
		"CheckInTime":       bookinfo.CheckInTime.String(),
		"ReservationPlan":   bookinfo.BookPlan.Title,
		"NumberOfGuests":    people,
		"RoomCount":         strconv.FormatUint(uint64(bookinfo.RoomCount), 10),
		"RoomType":          bookinfo.BookPlan.RoomType.String(),
		"MealType":          bookinfo.BookPlan.MealType.String(),
		"SmokingType":       bookinfo.BookPlan.SmokeType.String(),
		"TotalAmount":       strconv.FormatUint(uint64(bookinfo.TotalCost), 10),
		"Note":              bookinfo.Note,
		"HotelName":         store.Name + *store.BranchName,
		"HotelAddress":      store.Address,
		"HotelPhone":        store.Tel,
		"HotelURL":          store.SiteURL,
	}

	// テンプレートを解析
	tmpl, err := template.New("email").Parse(contentTemplate)
	if err != nil {
		fmt.Println("テンプレートの解析に失敗しました:", err)
		return nil, err
	}

	// テンプレートに値を埋め込む
	var content bytes.Buffer
	err = tmpl.Execute(&content, data)
	if err != nil {
		fmt.Println("テンプレートへの値の埋め込みに失敗しました:", err)
		return nil, err
	}

	// メールの本文を取得
	contentStr := content.String()
	return &contentStr, nil
}
