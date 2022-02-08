package mysql_dao

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
	"time"
)

type LoginLogDAO struct {
	db *sqlx.DB
}

func NewLoginLogDAO(db *sqlx.DB) *LoginLogDAO {
	return &LoginLogDAO{db}
}

func (dao *LoginLogDAO) RecordLogin(id int32, ip string ){
	now := time.Now().Unix()
	sqlStr := "insert into manage_account_login (account_id, login_time, login_ip) values (?, ?, ?);"
	_, err := dao.db.Exec(sqlStr, id, now, ip)
	raise(err)
}

func (dao *LoginLogDAO) GetLoginLog(id, limit, offset int32) ([]dataobject.LoginLog, int32) {
	res := make([]dataobject.LoginLog, 0)
	sqlStr := "select id, account_id, login_time, login_ip from manage_account_login where account_id = ? order by login_time desc limit ? offset ?;"

	rows, err := dao.db.Queryx(sqlStr, id, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var log dataobject.LoginLog
		err = rows.StructScan(&log)
		raise(err)
		res = append(res, log)
	}

	qryCount := "select count(*) from manage_account_login where account_id = ?"
	row := dao.db.QueryRow(qryCount, id)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}

func (dao *LoginLogDAO) GetLastLog() map[int32]dataobject.LoginLog {
	res := make(map[int32]dataobject.LoginLog, 0)
	sqlStr := "select account_id, login_time, login_ip from manage_account_login order by login_time desc;"

	rows, err := dao.db.Queryx(sqlStr)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var log dataobject.LoginLog
		err = rows.StructScan(&log)
		raise(err)
		_,ok := res[log.AccountId]
		if !ok {
			res[log.AccountId] = log
		} else {
			if log.LoginTime > res[log.AccountId].LoginTime {
				res[log.AccountId] = log
			}
		}
	}
	return res
}