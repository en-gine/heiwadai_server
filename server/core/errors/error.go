package errors

type Error struct {
	ErrType ErrorType
	Message string
}

type ErrorType int

const (
	InvalidValue ErrorType = iota
	UnPemitedOperation
)
