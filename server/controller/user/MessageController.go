package user

import (
	"context"

	"server/api/v1/user"
	userv1connect "server/api/v1/user/userconnect"
	"server/controller"
	usecase "server/core/usecase/user"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MessageController struct {
	messageUseCase usecase.MessageUsecase
}

var _ userv1connect.MessageControllerClient = &MessageController{}

func NewMessageController(messageUsecase *usecase.MessageUsecase) *MessageController {
	return &MessageController{
		messageUseCase: *messageUsecase,
	}
}

func (ac *MessageController) GetMessagesAfter(ctx context.Context, req *connect.Request[user.MessageRequest]) (*connect.Response[user.MessagesResponse], error) {
	var queryID *uuid.UUID
	uid, err := uuid.Parse(req.Msg.ID)
	if err != nil {
		queryID = nil
	} else {
		queryID = &uid
	}

	messages, domaiErr := ac.messageUseCase.GetAfter(queryID)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	var resMessages []*user.MessageResponse
	for _, message := range messages {
		displayDate := timestamppb.New(message.DisplayDate)
		resMessages = append(resMessages, &user.MessageResponse{
			ID:          message.ID.String(),
			Title:       message.Title,
			Content:     message.Content,
			AuthorID:    message.AuthorID.String(),
			DisplayDate: displayDate,
		})
	}

	return connect.NewResponse(&user.MessagesResponse{
		Messages: resMessages,
	}), nil
}
