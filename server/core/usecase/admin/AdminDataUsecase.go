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
	storeID uuid.UUID,
) (*entity.Admin, *errors.DomainError) {
	existAdmin, err := u.adminQuery.GetByID(ID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if existAdmin == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "登録されているユーザーが存在しません")
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
		existAdmin.Mail,
		IsActive,
		existAdmin.IsConfirmed,
		belongStore,
	)
	err = u.adminRepository.Update(updateData)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return updateData, nil
}

func (u *AdminDataUsecase) GetByID(ID uuid.UUID) (*entity.Admin, *errors.DomainError) {
	admin, err := u.adminQuery.GetByID(ID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if admin == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "該当のユーザーが見つかりません。")
	}
	return admin, nil
}

func (u *AdminDataUsecase) Delete(ID uuid.UUID) *errors.DomainError {
	admin, err := u.adminQuery.GetByID(ID)
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}
	if admin == nil {
		return errors.NewDomainError(errors.QueryDataNotFoundError, "該当のユーザーが見つかりません。")
	}

	err = u.adminRepository.Delete(ID)
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}
	if admin == nil {
		return errors.NewDomainError(errors.QueryDataNotFoundError, "該当のユーザーが見つかりません。")
	}
	return nil
}

func (u *AdminDataUsecase) GetAll() ([]*entity.Admin, *errors.DomainError) {
	admins, err := u.adminQuery.GetAll()
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if admins == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "該当のユーザーが見つかりません。")
	}
	return admins, nil
}
