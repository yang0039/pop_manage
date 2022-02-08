package mysql_dao

import (
	"github.com/jmoiron/sqlx"
)

type CallDAO struct {
	db *sqlx.DB
}

func NewCallDAO(db *sqlx.DB) *CallDAO {
	return &CallDAO{db}
}

func (dao *CallDAO) GetCallNum(start, end int64) (num int32) {
	var query = "select count(*) count from phonecall where add_time between ? and ?;"
	row := dao.db.QueryRowx(query, start, end)
	err := row.Scan(&num)
	raise(err)
	return
}
