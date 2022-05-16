package mysql_dao

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
	"time"
)

type RequestRecoreDAO struct {
	db *sqlx.DB
}

func NewRequestRecoreDAO(db *sqlx.DB) *RequestRecoreDAO {
	return &RequestRecoreDAO{db}
}

func (dao *RequestRecoreDAO) GetRequestRecords(accId, limit, offset int32) ([]*dataobject.RequestRecordDO, int32) {
	var accSql string
	if accId != 0 {
		accSql = fmt.Sprintf(" where account_id = %d", accId)
	}
	res := make([]*dataobject.RequestRecordDO, 0)
	sqlStr := "select id, account_id, url, method, client_ip, req_data, is_success, reason, add_time from manage_request_record" + accSql + " order by add_time desc limit ? offset ?;"
	rows, err := dao.db.Queryx(sqlStr, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var re dataobject.RequestRecordDO
		err = rows.StructScan(&re)
		raise(err)
		res = append(res, &re)
	}

	qryCount := "select count(*) from manage_request_record" + accSql
	row := dao.db.QueryRow(qryCount)
	var count int32
	err = row.Scan(&count)
	raise(err)

	return res, count
}

func (dao *RequestRecoreDAO) AddRequestRecords(accId int32, url, method, client_ip string) int32 {
	now := time.Now().Unix()
	sqlStr := "insert into manage_request_record (account_id, url, method, client_ip, add_time) values (?, ?, ?, ?, ?);"
	res, err := dao.db.Exec(sqlStr, accId, url, method, client_ip, now)
	raise(err)
	id,_ := res.LastInsertId()
	return int32(id)
}

func (dao *RequestRecoreDAO) AddRequestData(id int32, reqData string)  {
	sqlStr := "update manage_request_record set req_data = ? where id = ?"
	_, err := dao.db.Exec(sqlStr, reqData, id)
	raise(err)
}

func (dao *RequestRecoreDAO) AddRequestResult(id, isSuccess int32, reason string)  {
	sqlStr := "update manage_request_record set is_success = ?, reason = ? where id = ?"
	_, err := dao.db.Exec(sqlStr, isSuccess, reason, id)
	raise(err)
}

func (dao *RequestRecoreDAO) AddRequestAccountId(id, accId int32)  {
	sqlStr := "update manage_request_record set account_id = ? where id = ?"
	_, err := dao.db.Exec(sqlStr, accId, id)
	raise(err)
}

