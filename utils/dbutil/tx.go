package dbutil

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type (
	Tx interface {
		Select(dest interface{}, query string, args ...interface{}) error
		Get(dest interface{}, query string, args ...interface{}) error
		Exec(query string, args ...interface{}) (sql.Result, error)
		Commit() error
	}

	txImpl struct {
		tx *sqlx.Tx
	}
)

func (tx *txImpl) Exec(query string, args ...interface{}) (sql.Result, error) {
	return tx.tx.Exec(query, args)
}

func (tx *txImpl) Get(dest interface{}, query string, args ...interface{}) error {
	return tx.tx.Get(dest, query, args)
}

func (tx *txImpl) Select(dest interface{}, query string, args ...interface{}) error {
	return tx.tx.Select(dest, query, args)
}

func (tx *txImpl) Commit() error {
	return tx.tx.Commit()
}