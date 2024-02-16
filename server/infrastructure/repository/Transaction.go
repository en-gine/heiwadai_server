package repository

import (
	"context"
	"database/sql"

	"server/core/infra/repository"
	"server/infrastructure/logger"
)

var _ repository.ITransaction = &Transaction{}

type Transaction struct {
	db  *sql.DB
	ctx *context.Context
	Tx  *sql.Tx
}

func NewTransaction() *Transaction {
	db := InitDB()

	return &Transaction{
		db: db,
	}
}

func (r *Transaction) Begin(ctx *context.Context) error {
	r.ctx = ctx
	tx, err := r.db.BeginTx(*ctx, nil)
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
	return nil
}

func (r *Transaction) Rollback() {
	err := r.Tx.Rollback()
	if err != nil {
		logger.Errorf("rollback error: %v", err)
	}
}

func (r *Transaction) Exec(query string, args ...interface{}) (sql.Result, error) {
	return r.Tx.ExecContext(*r.ctx, query, args...)
}
