package admin

import (
	"context"
	"errors"

	adminv1 "server/api/v1/admin"
	"server/api/v1/shared"

	adminv1connect "server/api/v1/admin/adminconnect"
	"server/controller"

	"server/controller/util"
	"server/core/infra/queryService/types"
	usecase "server/core/usecase/admin"

	connect "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
)

type UserCheckinController struct {
	checkinUsecase usecase.UserCheckinUsecase
}

var _ adminv1connect.UserCheckinControllerClient = &UserCheckinController{}

func NewUserCheckinController(checkinUsecase *usecase.UserCheckinUsecase) *UserCheckinController {
	return &UserCheckinController{
		checkinUsecase: *checkinUsecase,
	}
}

func (u *UserCheckinController) GetAllRecent(ctx context.Context, req *connect.Request[adminv1.GetRecentAllCheckinRequest]) (*connect.Response[adminv1.CheckinsResponse], error) {
	msg := req.Msg
	checkin, domaiErr := u.checkinUsecase.GetAllRecent(int(msg.Limit))
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	var resCheckins []*adminv1.CheckinData
	for _, checkin := range checkin {
		user := checkin.User
		store := checkin.Store
		resCheckins = append(resCheckins,
			&adminv1.CheckinData{
				ID:              checkin.ID.String(),
				UserID:          user.ID.String(),
				StoreID:         store.ID.String(),
				StoreName:       store.Name,
				StoreBranchName: *store.BranchName,
				CheckinAt:       util.TimePtrToTimeStampPtr(&checkin.CheckInAt),
			},
		)
	}

	res := connect.NewResponse(&adminv1.CheckinsResponse{
		Checkins: resCheckins,
	})
	return res, nil
}

func (u *UserCheckinController) GetUserLog(ctx context.Context, req *connect.Request[adminv1.UserCheckinRequest]) (*connect.Response[adminv1.CheckinsResponse], error) {
	msg := req.Msg
	userID, err := uuid.Parse(msg.UserID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UUIDが正しい形式ではありません。"))
	}

	currentPage := int(*msg.Pager.CurrentPage)
	perPage := int(*msg.Pager.PerPage)
	pageQuery := types.NewPageQuery(&currentPage, &perPage)

	checkin, pageResponse, domaiErr := u.checkinUsecase.GetUserLog(userID, pageQuery)

	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	var resCheckins []*adminv1.CheckinData
	for _, checkin := range checkin {
		user := checkin.User
		store := checkin.Store
		resCheckins = append(resCheckins,
			&adminv1.CheckinData{
				ID:              checkin.ID.String(),
				UserID:          user.ID.String(),
				StoreID:         store.ID.String(),
				StoreName:       store.Name,
				StoreBranchName: *store.BranchName,
				CheckinAt:       util.TimePtrToTimeStampPtr(&checkin.CheckInAt),
			},
		)
	}

	var resPage *shared.PageResponse

	if pageResponse != nil {
		resPage = &shared.PageResponse{
			TotalCount:  uint32(pageResponse.TotalCount),
			CurrentPage: uint32(pageResponse.CurrentPage),
			PerPage:     uint32(pageResponse.PerPage),
			TotalPage:   uint32(pageResponse.TotalPage),
		}
	} else {
		resPage = &shared.PageResponse{
			TotalCount:  0,
			CurrentPage: 0,
			PerPage:     0,
			TotalPage:   0,
		}
	}
	res := &adminv1.CheckinsResponse{
		Checkins:     resCheckins,
		PageResponse: resPage,
	}
	return connect.NewResponse(res), nil
}
