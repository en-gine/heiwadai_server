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
	"server/infrastructure/parser"

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
	books, err := u.bookQuery.GetMyBooking(userID)
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
	NewTlDataID, err := u.bookQuery.GenerateBookDataID()
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}
	if NewTlDataID == nil || *NewTlDataID == "" {
		return errors.NewDomainError(errors.QueryError, "予約番号の生成に失敗しました。")
	}

	domainError, err := u.bookAPI.Cancel(book, *NewTlDataID)
	if err != nil {
		return errors.NewDomainError(errors.CancelButNeedFeedBack, "キャンセル処理がAPIレベルで失敗しました。")
	}
	if domainError != nil {
		return domainError
	}

	err = u.bookRepo.SoftDelete(bookID)
	if err != nil {
		logger.Errorf("キャンセル処理がDBレベルで失敗しました。%s", bookID.String())
		return errors.NewDomainError(errors.CancelButNeedFeedBack, "キャンセル処理は成功しましたが、DBの削除に失敗しました。")
	}

	store, err := u.storeQuery.GetStayableByID(book.BookPlan.Plan.StoreID)
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}

	tmpl, err := template.ParseFiles("core/usecase/template/CancelMail.html")
	if err != nil {
		return errors.NewDomainError(errors.CancelButNeedFeedBack, "予約はキャンセルしましたが、メールテンプレートの取得に失敗しました。")
	}

	// 予約キャンセル完了メールの内容を取得
	content, err := cancelMailContent(tmpl, book, store)
	if err != nil {
		return errors.NewDomainError(errors.CancelButNeedFeedBack, "予約はキャンセルしましたが、お客様へのメール作成に失敗しました。")
	}

	// 予約完了メール送信
	err = u.mailAction.Send(book.GuestData.Mail, "【"+store.Name+*store.BranchName+"】予約をキャンセルしました", *content, action.SendStyleHTML)
	if err != nil {
		return errors.NewDomainError(errors.CancelButNeedFeedBack, "予約はキャンセルしましたが、お客様へのメール送信に失敗しました。")
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
	BookPlan *entity.PlanStayDetail,
	BookUserID uuid.UUID,
	Note string,
) *errors.DomainError {

	store, err := u.storeQuery.GetStayableByID(BookPlan.Plan.StoreID)
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

	newBook := entity.CreateBooking(
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
	tmpl, err := template.ParseFiles("core/usecase/template/ReserveMail.html")
	if err != nil {
		return errors.NewDomainError(errors.CancelButNeedFeedBack, "予約は完了しましたが、メールテンプレートの取得に失敗しました。")
	}
	content, err := reserveMailContent(tmpl, newBook, store)
	if err != nil {
		return errors.NewDomainError(errors.CancelButNeedFeedBack, "予約は完了しましたが、お客様へのメール作成に失敗しました。")
	}

	// 予約完了メール送信
	err = u.mailAction.Send(GuestData.Mail, "【"+store.Name+*store.BranchName+"】宿泊予約完了のお知らせ", *content, action.SendStyleHTML)
	if err != nil {
		return errors.NewDomainError(errors.CancelButNeedFeedBack, "予約は完了しましたが、お客様へのメール送信に失敗しました。: "+err.Error())
	}
	return nil
}

func (u *BookUsecase) GetIsBookingUnderMaintenance() *entity.MaintenanceInfo {
	return entity.GetBookingUnderMaintenance()

}

func reserveMailContent(
	contentTemplate *template.Template,
	bookinfo *entity.Booking,
	store *entity.StayableStore,
) (*string, error) {

	// メールのデータを定義（実際のデータはアプリケーションから取得）
	var people string
	if bookinfo.Child > 0 {
		people = "大人: " + strconv.FormatUint(uint64(bookinfo.Adult), 10) + "名様／子ども：" + strconv.FormatUint(uint64(bookinfo.Child), 10)
	} else {
		people = "大人: " + strconv.FormatUint(uint64(bookinfo.Adult), 10) + "名様"
	}

	address := parser.ParseAddress(store.Address)
	data := map[string]string{
		"GuestName":              bookinfo.GuestData.LastName + bookinfo.GuestData.FirstName,
		"GuestMail":              bookinfo.GuestData.Mail,
		"ReservationNumber":      bookinfo.TlDataID,
		"CheckInDate":            bookinfo.StayFrom.Format("2006年1月2日"),
		"CheckOutDate":           bookinfo.StayTo.Format("2006年1月2日"),
		"CheckInTime":            bookinfo.CheckInTime.String(),
		"CheckInDateTimeFormat":  bookinfo.StayFrom.Format("2006-01-02") + "T00:00:00+09:00",
		"CheckOutDateTimeFormat": bookinfo.StayTo.Format("2006-01-02") + "T00:00:00+09:00",
		"ReservationPlan":        bookinfo.BookPlan.Plan.Title,
		"NumberOfGuests":         people,
		"RoomCount":              strconv.FormatUint(uint64(bookinfo.RoomCount), 10),
		"RoomType":               bookinfo.BookPlan.Plan.TlBookingRoomTypeName,
		"MealType":               bookinfo.BookPlan.Plan.MealType.String(),
		"SmokingType":            bookinfo.BookPlan.Plan.SmokeType.String(),
		"TotalAmount":            strconv.FormatUint(uint64(bookinfo.TotalCost), 10),
		"Note":                   bookinfo.Note,
		"HotelName":              store.Name + *store.BranchName,
		"HotelAddress":           store.Address,
		"HotelAddressLocality":   address[1],
		"HotelAddressRegion":     address[0],
		"HotelAddressStreet":     address[2],
		"HotelPostalCode":        store.ZipCode,
		"HotelPhone":             store.Tel,
		"HotelURL":               store.SiteURL,
	}

	// テンプレートを解析
	contentStr, err := analyzeTemplate(contentTemplate, data)
	if err != nil {
		return nil, err
	}

	return contentStr, nil
}

func cancelMailContent(
	contentTemplate *template.Template,
	bookinfo *entity.Booking,
	store *entity.StayableStore,
) (*string, error) {
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
		"ReservationPlan":   bookinfo.BookPlan.Plan.Title,
		"NumberOfGuests":    people,
		"RoomCount":         strconv.FormatUint(uint64(bookinfo.RoomCount), 10),
		"RoomType":          bookinfo.BookPlan.Plan.TlBookingRoomTypeName,
		"HotelName":         store.Name + *store.BranchName,
		"HotelAddress":      store.Address,
		"HotelPhone":        store.Tel,
		"HotelURL":          store.SiteURL,
	}

	// テンプレートを解析
	contentStr, err := analyzeTemplate(contentTemplate, data)
	if err != nil {
		return nil, err
	}

	return contentStr, nil
}

func analyzeTemplate(templateContent *template.Template, dataMap map[string]string) (*string, error) {

	// テンプレートに値を埋め込む
	var content bytes.Buffer
	err := templateContent.Execute(&content, dataMap)
	if err != nil {
		fmt.Println("テンプレートへの値の埋め込みに失敗しました:", err)
		return nil, err
	}
	str := content.String()
	return &str, nil
}
