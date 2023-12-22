package user

import (
	"context"
	"errors"

	"server/api/v1/user"
	userv1connect "server/api/v1/user/userconnect"
	"server/controller"
	usecase "server/core/usecase/user"
	"server/router"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserReportController struct {
	messageUseCase usecase.UserReportUsecase
}

var _ userv1connect.UserReportControllerClient = &UserReportController{}

func NewUserReportController(messageUsecase *usecase.UserReportUsecase) *UserReportController {
	return &UserReportController{
		messageUseCase: *messageUsecase,
	}
}

func (ac *UserReportController) Send(ctx context.Context, req *connect.Request[user.UserReportRequest]) (*connect.Response[emptypb.Empty], error) {
	if ctx.Value(router.UserIDKey) == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
	}

	userID, err := uuid.Parse(ctx.Value(router.UserIDKey).(string))
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。UUIDの形式が不正です。"))
	}

	domaiErr := ac.messageUseCase.Send(req.Msg.Title, req.Msg.Content, userID)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}
