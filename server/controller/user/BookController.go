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
			if stayStore.ID == book.BookPlan.StoreID {
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
	stayStore, domainErr := ac.storeUseCase.GetStayableByID(book.BookPlan.StoreID)
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

	plan := entity.RegenPlan(
		req.Msg.BookPlan.ID,
		req.Msg.BookPlan.Title,
		uint(req.Msg.BookPlan.Price),
		req.Msg.BookPlan.ImageURL,
		entity.RoomType(req.Msg.BookPlan.RoomType),
		entity.MealType{
			Morning: req.Msg.BookPlan.MealType.Morning,
			Dinner:  req.Msg.BookPlan.MealType.Dinner,
		},
		entity.SmokeType(req.Msg.BookPlan.SmokeType),
		req.Msg.BookPlan.OverView,
		uuid.MustParse(req.Msg.BookPlan.StoreID),
	)

	var note string
	if req.Msg.Note == nil {
		note = ""
	} else {
		note = *req.Msg.Note
	}

	domainErr := ac.bookUseCase.Reserve(
		req.Msg.StayFrom.AsTime(),
		req.Msg.StayTo.AsTime(),
		uint(req.Msg.Adult),
		uint(req.Msg.Child),
		uint(req.Msg.RoomCount),
		entity.CheckInTime(req.Msg.CheckInTime),
		uint(req.Msg.TotalCost),
		guest,
		plan,
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

	return &user.BookResponse{
		ID:          entity.ID.String(),
		StayFrom:    timestamppb.New(entity.StayFrom),
		StayTo:      timestamppb.New(entity.StayTo),
		Adult:       uint32(entity.Adult),
		Child:       uint32(entity.Child),
		RoomCount:   uint32(entity.RoomCount),
		CheckInTime: entity.CheckInTime.String(),
		TotalCost:   uint32(entity.TotalCost),
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
		Plan: PlanEntityToResponse(entity.BookPlan, bookstore),
	}
}

func PlanEntityToResponse(plan *entity.Plan, planStore *entity.StayableStore) *user.DisplayPlan {
	return &user.DisplayPlan{
		ID:              plan.ID,
		Title:           plan.Title,
		Price:           uint32(plan.Price),
		ImageURL:        plan.ImageURL,
		RoomTypeName:    plan.RoomType.String(),
		MealTypeName:    plan.MealType.String(),
		SmokeTypeName:   plan.SmokeType.String(),
		OverView:        plan.OverView,
		StoreID:         plan.StoreID.String(),
		StoreName:       planStore.Name,
		StoreBranchName: planStore.BranchName,
	}
}
