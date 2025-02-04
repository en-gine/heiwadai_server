package user

import (
	"context"
	"errors"

	"server/api/v1/shared"
	"server/api/v1/user"
	userv1connect "server/api/v1/user/userconnect"
	"server/controller"
	"server/core/entity"
	usecase "server/core/usecase/user"
	"server/router"

	connect "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookController struct {
	bookUseCase  usecase.BookUsecase
	storeUseCase usecase.StoreUsecase
}

var _ userv1connect.BookControllerClient = &BookController{}

func NewBookController(bookUsecase *usecase.BookUsecase, storeUseCase *usecase.StoreUsecase) *BookController {
	return &BookController{
		bookUseCase:  *bookUsecase,
		storeUseCase: *storeUseCase,
	}
}

func (ac *BookController) GetMyBook(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[user.BooksResponse], error) {
	if ctx.Value(router.UserIDKey) == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
	}

	userID, err := uuid.Parse(ctx.Value(router.UserIDKey).(string))
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。UUIDの形式が不正です。"))
	}

	if userID == uuid.Nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
	}
	books, domainErr := ac.bookUseCase.GetMyBook(userID)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	stayStores, domainErr := ac.storeUseCase.GetStayables()
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}

	var resBooks []*user.BookResponse
	for _, book := range books {
		var bookstore *entity.StayableStore
		for _, stayStore := range stayStores {
			if stayStore.ID == book.BookPlan.Plan.StoreID {
				bookstore = stayStore
				break
			}
		}
		resBooks = append(resBooks, BookEntityToResponse(book, bookstore))
	}

	return connect.NewResponse(&user.BooksResponse{
		Books: resBooks,
	}), nil
}

func (ac *BookController) GetBookByID(ctx context.Context, req *connect.Request[user.BookIDRequest]) (*connect.Response[user.BookResponse], error) {
	msg := req.Msg
	bookID, err := uuid.Parse(msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UUIDが正しい形式ではありません。"))
	}

	book, domainErr := ac.bookUseCase.GetByID(bookID)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	stayStore, domainErr := ac.storeUseCase.GetStayableByID(book.BookPlan.Plan.StoreID)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}

	bookRes := BookEntityToResponse(book, stayStore)
	return connect.NewResponse(bookRes), nil
}

func (ac *BookController) Cancel(ctx context.Context, req *connect.Request[user.BookIDRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	bookID, err := uuid.Parse(msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UUIDが正しい形式ではありません。"))
	}
	domainErr := ac.bookUseCase.Cancel(bookID)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *BookController) Reserve(ctx context.Context, req *connect.Request[user.ReserveRequest]) (*connect.Response[emptypb.Empty], error) {
	if ctx.Value(router.UserIDKey) == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
	}

	userID, err := uuid.Parse(ctx.Value(router.UserIDKey).(string))
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。UUIDの形式が不正です。"))
	}

	var pref *entity.Prefecture = nil
	if req.Msg.GuestData.Prefecture != nil {
		tmp := entity.Prefecture(*req.Msg.GuestData.Prefecture)
		pref = &tmp
	}
	guest := entity.CreateGuestData(
		req.Msg.GuestData.FirstName,
		req.Msg.GuestData.LastName,
		req.Msg.GuestData.FirstNameKana,
		req.Msg.GuestData.LastNameKana,
		req.Msg.GuestData.CompanyName,
		req.Msg.GuestData.ZipCode,
		pref,
		req.Msg.GuestData.City,
		req.Msg.GuestData.Address,
		req.Msg.GuestData.Tel,
		req.Msg.GuestData.Mail,
	)

	var morning bool
	var dinner bool
	for _, mealType := range req.Msg.RequestPlan.MealTypes {
		switch mealType {
		case user.MealType_Morning:
			morning = true
		case user.MealType_Dinner:
			dinner = true
		}
	}

	plan := entity.RegenPlan(
		req.Msg.RequestPlan.ID,
		req.Msg.RequestPlan.Title,
		uint(req.Msg.RequestPlan.Price),
		req.Msg.RequestPlan.ImageURL,
		entity.RoomType(req.Msg.RequestPlan.RoomType),
		entity.MealType{Morning: morning, Dinner: dinner},
		entity.SmokeType(req.Msg.RequestPlan.SmokeType),
		req.Msg.RequestPlan.OverView,
		uuid.MustParse(req.Msg.RequestPlan.StoreID),
		req.Msg.RequestPlan.TlBookingRoomTypeCode,
		req.Msg.RequestPlan.TlBookingRoomTypeName,
	)

	var stayDateInfos []entity.StayDateInfo
	for _, dateInfo := range req.Msg.PlanStayDateInfos {
		stayDateInfos = append(stayDateInfos, entity.StayDateInfo{
			StayDate:           dateInfo.StayDate.AsTime(),
			StayDateTotalPrice: uint(dateInfo.StayDateTotalPrice),
		})
	}

	bookPlan := &entity.PlanStayDetail{
		Plan:          plan,
		StayDateInfos: &stayDateInfos,
	}

	var note string
	if req.Msg.Note == nil {
		note = ""
	} else {
		note = *req.Msg.Note
	}
	checkIntime, domainErr := entity.NewCheckInTime(req.Msg.CheckInTime)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}

	domainErr = ac.bookUseCase.Reserve(
		req.Msg.StayFrom.AsTime(),
		req.Msg.StayTo.AsTime(),
		uint(req.Msg.Adult),
		uint(req.Msg.Child),
		uint(req.Msg.RoomCount),
		*checkIntime,
		uint(req.Msg.TotalCost),
		guest,
		bookPlan,
		userID,
		note,
	)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func BookEntityToResponse(entity *entity.Booking, bookstore *entity.StayableStore) *user.BookResponse {
	var pref *shared.Prefecture = nil
	if entity.GuestData.Prefecture != nil {
		tmp := entity.GuestData.Prefecture.ToInt()
		tempPref := shared.Prefecture(tmp)
		pref = &tempPref
	}

	var infos []*user.PlanStayDateInfo
	for _, info := range *entity.BookPlan.StayDateInfos {
		infos = append(infos, &user.PlanStayDateInfo{
			StayDate:           timestamppb.New(info.StayDate),
			StayDateTotalPrice: uint32(info.StayDateTotalPrice),
		})
	}

	return &user.BookResponse{
		ID:          entity.ID.String(),
		StayFrom:    timestamppb.New(entity.StayFrom),
		StayTo:      timestamppb.New(entity.StayTo),
		Adult:       uint32(entity.Adult),
		Child:       uint32(entity.Child),
		RoomCount:   uint32(entity.RoomCount),
		CheckInTime: entity.CheckInTime.String(),
		TotalCost:   uint32(entity.TotalCost),
		Note:        entity.Note,
		GuestData: &user.GuestData{
			FirstName:     entity.GuestData.FirstName,
			LastName:      entity.GuestData.LastName,
			FirstNameKana: entity.GuestData.FirstNameKana,
			LastNameKana:  entity.GuestData.LastNameKana,
			CompanyName:   entity.GuestData.CompanyName,
			ZipCode:       entity.GuestData.ZipCode,
			Prefecture:    pref,
			City:          entity.GuestData.City,
			Address:       entity.GuestData.Address,
			Tel:           entity.GuestData.Tel,
			Mail:          entity.GuestData.Mail,
		},
		Plan:              PlanEntityToResponse(entity.BookPlan.Plan, bookstore),
		PlanStayDateInfos: infos,
	}
}

func (ac *BookController) GetMentenanceInfo(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[user.BookMentenanceInfoResponse], error) {
	MenternaceInfo := ac.bookUseCase.GetIsBookingUnderMaintenance()
	return connect.NewResponse(&user.BookMentenanceInfoResponse{
		IsMentenance: MenternaceInfo.IsMaintenance,
		Message:      MenternaceInfo.Message,
	}), nil
}

func PlanEntityToResponse(plan *entity.Plan, planStore *entity.StayableStore) *user.DisplayPlan {
	var mealTypes []user.MealType
	if plan.MealType.Morning {
		mealTypes = append(mealTypes, user.MealType_Morning)
	}
	if plan.MealType.Dinner {
		mealTypes = append(mealTypes, user.MealType_Dinner)
	}

	return &user.DisplayPlan{
		ID:                    plan.ID,
		Title:                 plan.Title,
		Price:                 uint32(plan.Price),
		ImageURL:              plan.ImageURL,
		RoomType:              user.RoomType(plan.RoomType),
		RoomTypeName:          plan.RoomType.String(),
		MealTypes:             mealTypes,
		MealTypeName:          plan.MealType.String(),
		SmokeType:             user.SmokeType(plan.SmokeType),
		SmokeTypeName:         plan.SmokeType.String(),
		OverView:              plan.OverView,
		StoreID:               plan.StoreID.String(),
		StoreName:             planStore.Name,
		StoreBranchName:       planStore.BranchName,
		TlBookingRoomTypeCode: plan.TlBookingRoomTypeCode,
		TlBookingRoomTypeName: plan.TlBookingRoomTypeName,
	}
}
