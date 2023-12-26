package controller

import (
	"server/core/errors"
	"server/infrastructure/logger"

	"github.com/bufbuild/connect-go"
)

func ErrorHandler(domainErr *errors.DomainError) *connect.Error {
	logger.Error(domainErr.Error())

	switch domainErr.ErrType {
	case errors.InvalidParameter:
		return connect.NewError(connect.CodeInvalidArgument, domainErr)
	case errors.UnPemitedOperation:
		return connect.NewError(connect.CodePermissionDenied, domainErr)
	case errors.RepositoryError:
		return connect.NewError(connect.CodeInternal, domainErr)
	case errors.ActionError:
		return connect.NewError(connect.CodeInternal, domainErr)
	case errors.QueryError:
		return connect.NewError(connect.CodeInternal, domainErr)
	case errors.QueryDataNotFoundError:
		return connect.NewError(connect.CodeNotFound, domainErr)
	case errors.ErrorUnknown:
		return connect.NewError(connect.CodeUnknown, domainErr)
	default:
		return connect.NewError(connect.CodeUnknown, domainErr)
	}
}
