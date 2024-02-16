package repository

import "context"

type ITransaction interface {
	Begin(ctx *context.Context) error
	Commit() error
	Rollback()
}
