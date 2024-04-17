package controller

import (
	"server/core/errors"
	"server/infrastructure/logger"

	"github.com/bufbuild/connect-go"
)

func ErrorHandler(domainErr *errors.DomainError) *connect.Error {

	switch domainErr.ErrType {
	case errors.InvalidParameter:
		return connect.NewError(connect.CodeInvalidArgument, domainErr)
	case errors.UnPemitedOperation:
		return connect.NewError(connect.CodePermissionDenied, domainErr)
	case errors.AlreadyExist:
		return connect.NewError(connect.CodeAlreadyExists, domainErr)
	case errors.CancelButNeedFeedBack:
		return connect.NewError(connect.CodeCanceled, domainErr)
	case errors.RepositoryError:
		logger.Error(domainErr.Error())
		return connect.NewError(connect.CodeInternal, domainErr)
	case errors.ActionError:
		logger.Error(domainErr.Error())
		return connect.NewError(connect.CodeInternal, domainErr)
	case errors.QueryError:
		logger.Error(domainErr.Error())
		return connect.NewError(connect.CodeInternal, domainErr)
	case errors.QueryDataNotFoundError:
		logger.Error(domainErr.Error())
		return connect.NewError(connect.CodeNotFound, domainErr)
	case errors.ErrorUnknown:
		logger.Error(domainErr.Error())
		return connect.NewError(connect.CodeUnknown, domainErr)
	case errors.InternalError:
		logger.Error(domainErr.Error())
		return connect.NewError(connect.CodeInternal, domainErr)
	default:
		logger.Error(domainErr.Error())
		return connect.NewError(connect.CodeUnknown, domainErr)
	}
}
