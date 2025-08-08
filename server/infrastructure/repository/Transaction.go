package repository

import (
	"context"
	"database/sql"
	"time"

	"server/core/infra/repository"
	"server/infrastructure/logger"
)

var _ repository.ITransaction = &Transaction{}

type Transaction struct {
	db         *sql.DB
	ctx        *context.Context
	ctxCancel  context.CancelFunc
	Tx         *sql.Tx
}

func NewTransaction() *Transaction {
	db := InitDB()

	return &Transaction{
		db: db,
	}
}

func (r *Transaction) Begin(ctx context.Context) error {
	// Create context with timeout to prevent long-running transactions
	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	r.ctx = &timeoutCtx
	r.ctxCancel = cancel
	
	tx, err := r.db.BeginTx(timeoutCtx, nil)
	if err != nil {
		logger.Errorf("begin transaction error: %v", err)
		return err
	}
	r.Tx = tx
	return nil
}

func (r *Transaction) Commit() error {
	err := r.Tx.Commit()
	if err != nil {
		logger.Errorf("commit error: %v", err)
		return err
	}
	// Cancel context after successful commit
	if r.ctxCancel != nil {
		r.ctxCancel()
	}
	return nil
}

func (r *Transaction) Rollback() {
	err := r.Tx.Rollback()
	if err != nil {
		logger.Errorf("rollback error: %v", err)
	}
	// Cancel context after rollback
	if r.ctxCancel != nil {
		r.ctxCancel()
	}
}

func (r *Transaction) Exec(query string, args ...interface{}) (sql.Result, error) {
	return r.Tx.ExecContext(*r.ctx, query, args...)
}

func (r *Transaction) Tran() *sql.Tx {
	return r.Tx
}

func (r *Transaction) Context() *context.Context {
	return r.ctx
}
