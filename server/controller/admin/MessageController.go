package admin

import (
	"context"
	"errors"

	"server/api/v1/admin"
	adminv1connect "server/api/v1/admin/adminconnect"
	"server/controller"
	"server/controller/util"
	"server/core/entity"
	"server/core/infra/queryService/types"
	usecase "server/core/usecase/admin"

	connect "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MessageController struct {
	messageUseCase usecase.MessageUsecase
}

var _ adminv1connect.MessageControllerClient = &MessageController{}

func NewMessageController(messageUsecase *usecase.MessageUsecase) *MessageController {
	return &MessageController{
		messageUseCase: *messageUsecase,
	}
}

func (uc *MessageController) GetByID(ctx context.Context, req *connect.Request[admin.MessageIDRequest]) (*connect.Response[admin.MessageResponse], error) {
	msg := req.Msg
	messageID, err := uuid.Parse(msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UUIDが正しい形式ではありません。"))
	}

	message, domaiErr := uc.messageUseCase.GetByID(messageID)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}
	resMessage := MessageToResponse(message)

	return connect.NewResponse(resMessage), domaiErr
}

func (uc *MessageController) GetList(ctx context.Context, req *connect.Request[admin.GetMessageRequest]) (*connect.Response[admin.MessagesResponse], error) {
	var currentPage, perPage int
	if req.Msg.Pager.CurrentPage != nil {
		currentPage = int(*req.Msg.Pager.CurrentPage)
	}
	if req.Msg.Pager.PerPage != nil {
		perPage = int(*req.Msg.Pager.PerPage)
	}

	pager := types.NewPageQuery(
		&currentPage,
		&perPage,
	)

	entities, domaiErr := uc.messageUseCase.GetList(pager)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	var msgs []*admin.MessageResponse
	for _, entity := range entities {
		message := MessageToResponse(entity)
		msgs = append(msgs, message)
	}
	result := &admin.MessagesResponse{
		Messages: msgs,
	}
	return connect.NewResponse(result), nil
}

func (uc *MessageController) Create(ctx context.Context, req *connect.Request[admin.MessageCreateRequest]) (*connect.Response[admin.MessageResponse], error) {
	adminID := ctx.Value("userID").(uuid.UUID)

	if adminID == uuid.Nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
	}

	msg := req.Msg

	entity, domainErr := uc.messageUseCase.Create(msg.Title, msg.Content, msg.DisplayDate.AsTime(), adminID)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	message := MessageToResponse(entity)

	return connect.NewResponse(message), nil
}

func (uc *MessageController) Update(ctx context.Context, req *connect.Request[admin.MessageUpdateRequest]) (*connect.Response[admin.MessageResponse], error) {
	adminID := ctx.Value("userID").(uuid.UUID)

	if adminID == uuid.Nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
	}

	msg := req.Msg

	entity, domaiErr := uc.messageUseCase.Update(
		msg.Title,
		msg.Content,
		util.TimeStampPtrToTimePtr(msg.DisplayDate),
		adminID,
		uuid.MustParse(msg.ID),
	)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	message := MessageToResponse(entity)

	return connect.NewResponse(message), nil
}

func (uc *MessageController) Delete(ctx context.Context, req *connect.Request[admin.MessageIDRequest]) (*connect.Response[emptypb.Empty], error) {
	domainErr := uc.messageUseCase.Delete(uuid.MustParse(req.Msg.ID))
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func MessageToResponse(entity *entity.Message) *admin.MessageResponse {
	return &admin.MessageResponse{
		ID:          entity.ID.String(),
		Title:       entity.Title,
		Content:     entity.Content,
		AuthorID:    entity.AuthorID.String(),
		DisplayDate: timestamppb.New(entity.DisplayDate),
	}
}
