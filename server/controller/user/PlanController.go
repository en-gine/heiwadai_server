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

func (ac *PlanController) Search(ctx context.Context, req *connect.Request[user.PlanSearchRequest]) (*connect.Response[user.PlansResponse], error) {
	msg := req.Msg
	var storeUUIDs []uuid.UUID
	for _, storeID := range msg.StoreIDs {
		storeUUID, err := uuid.Parse(storeID)
		if err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("店舗IDが正しい形式ではありません。"))
		}
		storeUUIDs = append(storeUUIDs, storeUUID)
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

	plans, domainErr := ac.planUseCase.Search(
		storeUUIDs,
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

	var displayPlans []*user.DisplayPlan
	for _, plan := range *plans {
		planStore, domainErr := ac.storeUseCase.GetStayableByID(plan.StoreID)
		if domainErr != nil {
			return nil, controller.ErrorHandler(domainErr)
		}
		plans := PlanEntityToResponse(&plan, planStore)
		displayPlans = append(displayPlans, plans)

	}
	return connect.NewResponse(
		&user.PlansResponse{
			Plans: displayPlans,
		}), nil
}
