package mysql_dao

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
)

type UserOperaDAO struct {
	db *sqlx.DB
}

func NewUserOperaDAO(db *sqlx.DB) *UserOperaDAO {
	return &UserOperaDAO{db}
}

func (dao *UserOperaDAO) AddOperaRecords(uo *dataobject.UserOpera) int32 {
	sqlStr := "insert into manage_user_opera (account_id, user_id, opera_type, opera_content, add_time) values (?, ?, ?, ?, ?);"
	res, err := dao.db.Exec(sqlStr,uo.AccountId, uo.UserId, uo.OperaType, uo.OperaContent, uo.AddTime)
	raise(err)
	id,_ := res.LastInsertId()
	return int32(id)
}


func (dao *UserOperaDAO) GetUserOperaRecords(accId, userId, limit, offset int32 ) ([]*dataobject.UserOpera, int32) {
	res := make([]*dataobject.UserOpera, 0)
	var accSql string
	if accId != 1 {   // accId为1的是admin的id，admin可以看所有数据
		accSql = fmt.Sprintf("and account_id=%d", accId)
	}
	sqlStr := "select id, account_id, user_id, opera_type, opera_content, add_time from manage_user_opera where user_id = ? " + accSql + " order by add_time desc limit ? offset ?;"
	rows, err := dao.db.Queryx(sqlStr, userId, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var re dataobject.UserOpera
		err = rows.StructScan(&re)
		raise(err)
		res = append(res, &re)
	}

	qryCount := "select count(*) from manage_user_opera  where user_id = ? " + accSql
	row := dao.db.QueryRow(qryCount, userId)
	var count int32
	err = row.Scan(&count)
	raise(err)

	return res, count
}



