package dbutil

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type (
	dbFake struct {
		db  *sqlx.DB
		err error
	}
)

func NewFakeConnection(conf DBConf) (db *dbFake, err error) {
	return &dbFake{}, err
}

func (db *dbFake) Ping() error {
	return db.err
}

func (db *dbFake) Beginx() (Tx, error) {
	return &txFake{}, nil
}

func (db *dbFake) Exec(query string, args ...interface{}) (sql.Result, error) {
	return resultFake{}, nil
}

func (db *dbFake) Get(dest interface{}, query string, args ...interface{}) error {
	return nil
}

func (db *dbFake) Select(dest interface{}, query string, args ...interface{}) error {
	return nil
}
