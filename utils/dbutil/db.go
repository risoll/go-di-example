package dbutil

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type (
	DB interface {
		Ping() error
		Get(dest interface{}, query string, args ...interface{}) error
		Exec(query string, args ...interface{}) (sql.Result, error)
		Select(dest interface{}, query string, args ...interface{}) error
		Beginx() (Tx, error)
	}

	dbImpl struct {
		db *sqlx.DB
	}

	DBConf struct {
		DriverName string
		Conn string
		MaxIdle int
		MaxOpen int
	}
)

func NewConnection(conf DBConf) (db *dbImpl, err error) {
	dbConn, err := sqlx.Connect(conf.DriverName, conf.Conn)
	if err != nil {
		return &dbImpl{}, err
	}

	dbConn.SetMaxOpenConns(conf.MaxOpen)
	dbConn.SetMaxIdleConns(conf.MaxIdle)
	return &dbImpl{db:dbConn}, err
}

func (db *dbImpl) Ping() error {
	return db.db.Ping()
}

func (db *dbImpl) Beginx() (Tx, error) {
	tx, err := db.db.Beginx()
	return &txImpl{tx: tx}, err
}

func (db *dbImpl) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.db.Exec(query, args)
}

func (db *dbImpl) Get(dest interface{}, query string, args ...interface{}) error {
	return db.db.Get(dest, query, args)
}

func (db *dbImpl) Select(dest interface{}, query string, args ...interface{}) error {
	return db.db.Select(dest, query, args)
}