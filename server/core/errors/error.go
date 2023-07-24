package errors

import "errors"

type DomainError struct {
	ErrType ErrorType
	err     error
}

type ErrorType int

const (
	InvalidParameter ErrorType = iota
	UnPemitedOperation
	RepositoryError
	ActionError
	QueryError
	QueryDataNotFoundError
)

func (e DomainError) Error() string {
	return e.err.Error()
}

func (e DomainError) Unwrap() error {
	return e.err
}

func (e DomainError) GetType() ErrorType {
	return e.ErrType
}

func NewDomainError(errType ErrorType, message string) *DomainError {
	return &DomainError{ErrType: errType, err: errors.New(message)}
}
