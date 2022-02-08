package mysql_client

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/jmoiron/sqlx"
	"pop-api/baselib/logger"
	"time"
)

type MySQLConfig struct {
	Index int    // distribute index
	Type  string // for distribute type

	Name   string // for trace
	DSN    string // data source name
	Active int    // pool
	Idle   int    // pool
}

func NewSqlxDB(c *MySQLConfig) (db *sqlx.DB) {
	db, err := sqlx.Connect("mysql", c.DSN)
	if err != nil {
		logger.LogSugar.Errorf("Connect db error: %v", err)
	}

	db.SetMaxOpenConns(c.Active)
	db.SetMaxIdleConns(c.Idle)
	db.SetConnMaxLifetime(30 * time.Second)
	return
}
