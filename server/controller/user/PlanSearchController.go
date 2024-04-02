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
	for _, storeId := range msg.StoreIDs {
		storeUUID, err := uuid.Parse(storeId)
		if err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("StoreIDが正しい形式ではありません。"))
		}
		storeUUIDs = append(storeUUIDs, storeUUID)
	}

	if msg.RoomCount == 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("RoomCountは1以上で指定してください。"))
	}

	var smokeTypes []entity.SmokeType
	for _, smokeType := range msg.SmokeTypes {
		smokeTypes = append(smokeTypes, entity.SmokeType(smokeType))
	}

	var roomTypes []entity.RoomType
	for _, roomType := range msg.RoomTypes {
		roomTypes = append(roomTypes, entity.RoomType(roomType))
	}

	plans, domainErr := ac.planUseCase.Search(
		storeUUIDs,
		msg.StayFrom.AsTime(),
		msg.StayTo.AsTime(),
		int(msg.Adult),
		int(msg.Child),
		int(msg.RoomCount),
		smokeTypes,
		entity.MealType{
			Morning: msg.MealType.Morning,
			Dinner:  msg.MealType.Dinner,
		},
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
