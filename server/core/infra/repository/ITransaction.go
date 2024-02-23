package repository

import (
	"context"
	"database/sql"
)

type ITransaction interface {
	Begin(ctx context.Context) error
	Commit() error
	Rollback()
	Tran() *sql.Tx
	Context() *context.Context
}
