package admin

import (
	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"

	"github.com/google/uuid"
)

type AdminDataUsecase struct {
	adminRepository repository.IAdminRepository
	adminQuery      queryservice.IAdminQueryService
	storeQuery      queryservice.IStoreQueryService
}

func NewAdminDataUsecase(adminRepository repository.IAdminRepository, adminQuery queryservice.IAdminQueryService, storeQuery queryservice.IStoreQueryService) *AdminDataUsecase {
	return &AdminDataUsecase{
		adminRepository: adminRepository,
		adminQuery:      adminQuery,
		storeQuery:      storeQuery,
	}
}

func (u *AdminDataUsecase) Update(
	ID uuid.UUID,
	Name string,
	IsActive bool,
	Mail string,
	storeID uuid.UUID,
) (*entity.Admin, *errors.DomainError) {

	existAdmin, err := u.adminQuery.GetByID(ID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if existAdmin == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "登録されているユーザーが存在しません")
	}

	existAdmin, err = u.adminQuery.GetByMail(Mail)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if existAdmin == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "このアドレスで登録されているユーザーが存在しません")
	}

	belongStore, err := u.storeQuery.GetByID(storeID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "店舗の検索に失敗しました")
	}

	if belongStore == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "IDで指定された店舗が見つかりません")

	}

	updateData := entity.RegenAdmin(
		ID,
		Name,
		Mail,
		IsActive,
		belongStore,
	)
	err = u.adminRepository.Update(updateData)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return updateData, nil
}
