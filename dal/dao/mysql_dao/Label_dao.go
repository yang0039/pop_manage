package mysql_dao

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
	"time"
)

type LabelDAO struct {
	db *sqlx.DB
}

func NewLabelDAO(db *sqlx.DB) *LabelDAO {
	return &LabelDAO{db}
}

func (dao *LabelDAO) AddLabel(name, operator string ) int32 {
	now := time.Now().Unix()
	sqlStr := "insert into manage_label (label_name, operator, add_time) values (?, ?, ?);"
	res, err := dao.db.Exec(sqlStr, name, operator, now)
	raise(err)
	id,_ := res.LastInsertId()
	return int32(id)
}

func (dao *LabelDAO) GetLabels() []*dataobject.Label {
	res := make([]*dataobject.Label, 0)
	sqlStr := "select id, label_name, operator, add_time from manage_label where is_delete = 0;"

	rows, err := dao.db.Queryx(sqlStr)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var log dataobject.Label
		err = rows.StructScan(&log)
		raise(err)
		res = append(res, &log)
	}
	return res
}

func (dao *LabelDAO) DelLabel(id int32) {
	sqlStr := "update manage_label set is_delete = 1 where id = ?;"
	_, err := dao.db.Exec(sqlStr, id)
	raise(err)
}

func (dao *LabelDAO) UpdateLabel(id int32, name string) {
	sqlStr := "update manage_label set label_name = ? where id = ?;"
	_, err := dao.db.Exec(sqlStr, name, id)
	raise(err)
}

func (dao *LabelDAO) GetNames(ids []string) []string {
	res := make([]string, 0)
	sqlStr := "select label_name from manage_label where is_delete = 0 and id in (?);"
	query, args, err := sqlx.In(sqlStr, ids)
	raise(err)
	rows, err := dao.db.Queryx(query, args...)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		raise(err)
		res = append(res, name)
	}
	return res
}