package admin

import (
	"context"
	"errors"

	"server/api/v1/admin"
	adminv1connect "server/api/v1/admin/adminconnect"
	"server/api/v1/shared"
	"server/controller"
	"server/core/entity"
	"server/core/infra/queryService/types"
	usecase "server/core/usecase/admin"
	"server/router"

	connect "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MailMagazineController struct {
	magazineUseCase usecase.MailMagazineUsecase
}

var _ adminv1connect.MailMagazineControllerClient = &MailMagazineController{}

func NewMailMagazineController(magazineUsecase *usecase.MailMagazineUsecase) *MailMagazineController {
	return &MailMagazineController{
		magazineUseCase: *magazineUsecase,
	}
}

func (uc *MailMagazineController) GetList(ctx context.Context, req *connect.Request[admin.GetMailMagazineListRequest]) (*connect.Response[admin.MailMagazinesResponse], error) {
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
	entities, pageResponse, domaiErr := uc.magazineUseCase.GetList(pager)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	var mailMagazines []*admin.MailMagazine

	for _, entity := range entities {
		var prefs []shared.Prefecture
		for _, pref := range *entity.TargetPrefecture {
			prefs = append(prefs, shared.Prefecture(pref))
		}
		var sentAt *timestamppb.Timestamp
		if entity.SentAt != nil {
			sentAt = timestamppb.New(*entity.SentAt)
		}
		magazine := &admin.MailMagazine{
			ID:                 entity.ID.String(),
			Title:              entity.Title,
			Content:            entity.Content,
			AuthorID:           entity.AuthorID.String(),
			TargetPrefecture:   prefs,
			MailMagazineStatus: admin.MailMagazineStatus(entity.MailMagazineStatus),
			UnsentCount:        int32(entity.UnsentCount),
			SentCount:          int32(entity.SentCount),
			SentAt:             sentAt,
			CreateAt:           timestamppb.New(entity.CreateAt),
		}
		mailMagazines = append(mailMagazines, magazine)
	}
	resPage := &shared.PageResponse{
		TotalCount:  uint32(pageResponse.TotalCount),
		CurrentPage: uint32(pageResponse.CurrentPage),
		PerPage:     uint32(pageResponse.PerPage),
		TotalPage:   uint32(pageResponse.TotalPage),
	}

	result := &admin.MailMagazinesResponse{
		MailMagazines: mailMagazines,
		PageResponse:  resPage,
	}
	return connect.NewResponse(result), nil
}

func (uc *MailMagazineController) GetByID(ctx context.Context, req *connect.Request[admin.MailMagazineIDRequest]) (*connect.Response[admin.MailMagazine], error) {
	entity, domaiErr := uc.magazineUseCase.GetByID(uuid.MustParse(req.Msg.ID))
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	var prefs []shared.Prefecture
	for _, pref := range *entity.TargetPrefecture {
		prefs = append(prefs, shared.Prefecture(pref))
	}

	result := &admin.MailMagazine{
		ID:                 entity.ID.String(),
		Title:              entity.Title,
		Content:            entity.Content,
		AuthorID:           entity.AuthorID.String(),
		TargetPrefecture:   prefs,
		MailMagazineStatus: admin.MailMagazineStatus(entity.MailMagazineStatus),
		UnsentCount:        int32(entity.UnsentCount),
		SentCount:          int32(entity.SentCount),
	}
	return connect.NewResponse(result), nil
}

func (uc *MailMagazineController) SaveDraft(ctx context.Context, req *connect.Request[admin.SaveDraftRequest]) (*connect.Response[admin.MailMagazine], error) {
	if ctx.Value(router.UserIDKey) == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
	}

	adminID, err := uuid.Parse(ctx.Value(router.UserIDKey).(string))
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。UUIDの形式が不正です。"))
	}

	msg := req.Msg
	var prefs []entity.Prefecture
	for _, pref := range msg.TargetPrefectures {
		prefs = append(prefs, entity.Prefecture(pref))
	}
	entity, domainErr := uc.magazineUseCase.SaveDraft(msg.Title, msg.Content, &prefs, adminID)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	var sentAt *timestamppb.Timestamp
	if entity.SentAt != nil {
		sentAt = timestamppb.New(*entity.SentAt)
	}

	return connect.NewResponse(&admin.MailMagazine{
		ID:                 entity.ID.String(),
		Title:              entity.Title,
		Content:            entity.Content,
		AuthorID:           entity.AuthorID.String(),
		TargetPrefecture:   msg.TargetPrefectures,
		MailMagazineStatus: admin.MailMagazineStatus(entity.MailMagazineStatus),
		UnsentCount:        int32(entity.UnsentCount),
		SentCount:          int32(entity.SentCount),
		SentAt:             sentAt,
		CreateAt:           timestamppb.New(entity.CreateAt),
	}), nil
}

func (uc *MailMagazineController) Update(ctx context.Context, req *connect.Request[admin.UpdateMailMagazineRequest]) (*connect.Response[admin.MailMagazine], error) {
	if ctx.Value(router.UserIDKey) == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
	}

	adminID, err := uuid.Parse(ctx.Value(router.UserIDKey).(string))
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。UUIDの形式が不正です。"))
	}

	msg := req.Msg
	var entityPrefs []entity.Prefecture
	for _, pref := range msg.TargetPrefectures {
		entityPrefs = append(entityPrefs, entity.Prefecture(pref))
	}
	entity, domaiErr := uc.magazineUseCase.Update(
		msg.Title,
		msg.Content,
		&entityPrefs,
		adminID,
		uuid.MustParse(msg.ID),
	)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	var reqPrefs []shared.Prefecture

	if entity.TargetPrefecture != nil {
		reqPrefs = make([]shared.Prefecture, len(*entity.TargetPrefecture))
		for _, pref := range *entity.TargetPrefecture {
			reqPrefs = append(reqPrefs, shared.Prefecture(pref))
		}
	}

	magazine := &admin.MailMagazine{
		ID:                 entity.ID.String(),
		Title:              entity.Title,
		Content:            entity.Content,
		AuthorID:           entity.AuthorID.String(),
		TargetPrefecture:   reqPrefs,
		MailMagazineStatus: admin.MailMagazineStatus(entity.MailMagazineStatus),
		UnsentCount:        int32(entity.UnsentCount),
		SentCount:          int32(entity.SentCount),
	}

	return connect.NewResponse(magazine), nil
}

func (uc *MailMagazineController) Delete(ctx context.Context, req *connect.Request[admin.DeleteMailMagazineRequest]) (*connect.Response[emptypb.Empty], error) {
	domainErr := uc.magazineUseCase.Delete(uuid.MustParse(req.Msg.ID))
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (uc *MailMagazineController) Send(ctx context.Context, req *connect.Request[admin.SendMailMagazineRequest]) (*connect.Response[emptypb.Empty], error) {
	domainErr := uc.magazineUseCase.Send(uuid.MustParse(req.Msg.ID))
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
