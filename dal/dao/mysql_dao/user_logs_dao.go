package mysql_dao

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
)

type UserLogsDAO struct {
	db *sqlx.DB
}

func NewUserLogsDAO(db *sqlx.DB) *UserLogsDAO {
	return &UserLogsDAO{db}
}

func (dao *UserLogsDAO) GetUerLogs(userId, limit, offset int32) ([]*dataobject.UserLogsDo, int32) {
	res := make([]*dataobject.UserLogsDo, 0)
	var table string
	if Dbindex == 1 {
		table = "user_logs"
	} else {
		table = "user_logs1"
		if userId %2 == 0 {
			table = "user_logs2"
		}
	}
	sqlStr := fmt.Sprintf("select user_id, auth_id, ip, device_model, system_version, app_version, created_at from  %s where user_id = ? order by id desc limit ? offset ?;", table)
	rows, err := dao.db.Queryx(sqlStr, userId, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var info  dataobject.UserLogsDo
		err = rows.StructScan(&info)
		raise(err)
		res = append(res, &info)
	}

	sqlCount := fmt.Sprintf("select count(*) from %s where user_id = ?;", table)
	row := dao.db.QueryRow(sqlCount, userId)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}

func (dao *UserLogsDAO) GetUerLastLogs(userId int32) (string, string) {
	var table string
	if Dbindex == 1 {
		table = "user_logs"
	} else {
		table = "user_logs1"
		if userId %2 == 0 {
			table = "user_logs2"
		}
	}
	sqlStr := fmt.Sprintf("select ip, device_model from  %s where user_id = ? order by id desc limit 1;", table)
	row := dao.db.QueryRowx(sqlStr, userId)
	var ip, device string
	err := row.Scan(&ip, &device)
	raise(err)
	return ip, device
}