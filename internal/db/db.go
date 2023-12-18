package db

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

const (
	DefaultDBMS = "mysql"
	MaxConns    = 4
)

var sqlConn *sql.DB

// OpenWithConn
func OpenWithConn(conn *sql.DB) error {
	if conn == nil {
		return fmt.Errorf("*sql.DB is null")
	}

	sqlConn = conn
	return nil
}

// Open
func Open(cfg Config) error {
	if cfg.DBMS == "" {
		cfg.DBMS = DefaultDBMS
	}

	db, err := sql.Open(cfg.DSN(), cfg.DBMS)
	if err != nil {
		return errors.Wrap(err, "sql.Open fail")
	}

	db.SetMaxOpenConns(MaxConns)
	db.SetMaxIdleConns(MaxConns)

	sqlConn = db
	return nil
}

// Close
func Close() {
	if sqlConn != nil {
		sqlConn.Close()
	}
}

// DB
func DB() *sql.DB {
	return sqlConn
}
