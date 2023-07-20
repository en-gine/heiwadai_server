package repository

import (
	"context"
	"database/sql"
	"server/core/infra/repository"
)

var _ repository.ITransaction = &Transaction{}

type Transaction struct {
	db *sql.DB
	tx *sql.Tx
}

func NewTransaction() *Transaction {
	db := InitDB()

	return &Transaction{
		db: db,
	}
}
func (r *Transaction) Begin() error {
	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	r.tx = tx
	return nil
}

func (r *Transaction) Commit() error {
	err := r.tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *Transaction) Rollback() error {
	err := r.tx.Rollback()
	if err != nil {
		return err
	}
	return nil
}
