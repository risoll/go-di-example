package dbutil

import (
	"database/sql"
)

type (
	txFake struct {
		err error
	}
)

func (tx *txFake) Exec(query string, args ...interface{}) (sql.Result, error) {
	return resultFake{}, nil
}

func (tx *txFake) Get(dest interface{}, query string, args ...interface{}) error {
	return nil
}

func (tx *txFake) Select(dest interface{}, query string, args ...interface{}) error {
	return nil
}

func (tx *txFake) Commit() error {
	return nil
}