package user

import (
	"context"
	"errors"

	"server/api/v1/user"
	userv1connect "server/api/v1/user/userconnect"
	"server/controller"
	"server/core/entity"
	usecase "server/core/usecase/user"

	connect "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
)

type PlanController struct {
	planUseCase  usecase.PlanUsecase
	storeUseCase usecase.StoreUsecase
}

var _ userv1connect.PlanControllerClient = &PlanController{}

func NewPlanController(bookUsecase *usecase.PlanUsecase, storeUseCase *usecase.StoreUsecase) *PlanController {
	return &PlanController{
		planUseCase:  *bookUsecase,
		storeUseCase: *storeUseCase,
	}
}

func (ac *PlanController) Search(ctx context.Context, req *connect.Request[user.PlanSearchRequest]) (*connect.Response[user.SearchPlanResponse], error) {
	msg := req.Msg
	var storeUUIDs []uuid.UUID
	for _, storeID := range msg.StoreIDs {
		storeUUID, err := uuid.Parse(storeID)
		if err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("店舗IDが正しい形式ではありません。"))
		}
		storeUUIDs = append(storeUUIDs, storeUUID)
	}

	var stayStores []*entity.StayableStore
	if len(storeUUIDs) == 0 {
		strs, err := ac.storeUseCase.GetStayables()
		if err != nil {
			return nil, controller.ErrorHandler(err)
		}
		stayStores = strs
	} else {
		for _, storeID := range storeUUIDs {
			stayStore, err := ac.storeUseCase.GetStayableByID(storeID)
			if err != nil {
				return nil, controller.ErrorHandler(err)
			}
			if stayStore == nil {
				return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("宿泊可能な店舗が見つかりません。"))
			}

			stayStores = append(stayStores, stayStore)
		}
	}

	if msg.RoomCount <= 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("部屋数は1以上で指定してください。"))
	}

	if msg.Adult <= 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("大人の人数を指定してください。"))
	}

	var smokeTypes *[]entity.SmokeType
	if len(msg.SmokeTypes) > 0 {
		var smokeTypeList []entity.SmokeType
		for _, smokeType := range msg.SmokeTypes {
			smokeTypeList = append(smokeTypeList, entity.SmokeType(smokeType))
		}
		smokeTypes = &smokeTypeList
	}

	var roomTypes *[]entity.RoomType
	if len(msg.RoomTypes) > 0 {
		var roomTypeList []entity.RoomType
		for _, roomType := range msg.RoomTypes {
			roomTypeList = append(roomTypeList, entity.RoomType(roomType))
		}
		roomTypes = &roomTypeList
	}

	var morning bool
	var dinner bool
	var mealType *entity.MealType
	for _, mealType := range req.Msg.MealTypes {
		switch mealType {
		case user.MealType_Morning:
			morning = true
		case user.MealType_Dinner:
			dinner = true
		}
	}

	if morning || dinner {
		mealType = &entity.MealType{
			Morning: morning,
			Dinner:  dinner,
		}
	}

	candidates, domainErr := ac.planUseCase.Search(
		stayStores,
		msg.StayFrom.AsTime(),
		msg.StayTo.AsTime(),
		int(msg.Adult),
		int(msg.Child),
		int(msg.RoomCount),
		smokeTypes,
		mealType,
		roomTypes,
	)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}

	var plans []*user.DisplayPlanWithSearchResultOption
	for _, candidate := range *candidates {
		planStore := func() *entity.StayableStore {
			for _, store := range stayStores {
				if store.ID == candidate.Plan.StoreID {
					return store
				}
			}
			return nil
		}()
		if planStore == nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("宿泊可能な店舗が見つかりません。"))
		}
		displayPlan := PlanEntityToResponse(candidate.Plan, planStore)

		resPlan := &user.DisplayPlanWithSearchResultOption{
			Plan:                 displayPlan,
			MinimumPrice:         uint32(candidate.MinimumPrice),
			PricePerCategory:     user.PricePerCategory(candidate.PricePerCategory),
			PricePerCategoryName: candidate.PricePerCategory.String(),
		}
		plans = append(plans, resPlan)

	}
	return connect.NewResponse(
		&user.SearchPlanResponse{
			Plans: plans,
		}), nil
}

func (ac *PlanController) GetDetail(ctx context.Context, req *connect.Request[user.PlanDetailRequest]) (*connect.Response[user.PlanResponse], error) {
	msg := req.Msg
	stayStoreID, err := uuid.Parse(msg.StayStoreID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("プランIDが正しい形式ではありません。"))
	}
	var roomType entity.RoomType = entity.RoomType(msg.RoomType)
	stayStore, err := ac.storeUseCase.GetStayableByID(stayStoreID)

	plan, domainErr := ac.planUseCase.GetDetail(msg.PlanID, msg.StayFrom.AsTime(), msg.StayTo.AsTime(), int(msg.Adult), int(msg.Child), int(msg.RoomCount), &roomType, stayStore)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}

	displayPlan := PlanEntityToResponse(plan, stayStore)

	return connect.NewResponse(
		&user.PlanResponse{
			Plan: displayPlan,
		}), nil
}
