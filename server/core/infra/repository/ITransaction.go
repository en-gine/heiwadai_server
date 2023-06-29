package repository

type ITransaction interface {
	Commit() error
	Begin() error
	Rollback() error
}
