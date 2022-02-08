package mysql_dao

import (
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
)

type RawMessageRowDAO struct {
	db *sqlx.DB
}

func NewRawMessageRowDAO(db *sqlx.DB) *RawMessageRowDAO {
	return &RawMessageRowDAO{db}
}

func (dao *RawMessageRowDAO)GetRawMessageRows(rawIds []int64) map[int64]*dataobject.RawMessageRow {
	res := make(map[int64]*dataobject.RawMessageRow, 0)
	if len(rawIds) == 0 {
		return res
	}
	query, args, err := sqlx.In(`select * from raw_message where raw_id in (?)`, rawIds)
	raise(err)
	rows, err := dao.db.Queryx(query, args...)
	defer rows.Close()
	raise(err)
	for rows.Next() {
		var rmr dataobject.RawMessageRow
		err := rows.StructScan(&rmr)
		raise(err)
		res[rmr.RawId] = &rmr
	}
	return res
}
