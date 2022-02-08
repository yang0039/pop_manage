package mysql_dao

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
	"time"
)

type AllowIpDAO struct {
	db *sqlx.DB
}

func NewAllowIpDAO(db *sqlx.DB) *AllowIpDAO {
	return &AllowIpDAO{db}
}

func (dao *AllowIpDAO) AddAllowIp(operator, ip string) int32 {
	now := time.Now().Unix()
	sqlStr := "insert into manage_login_ip (operator, ip, add_time) values (?, ?, ?);"
	result, err := dao.db.Exec(sqlStr, operator, ip, now)
	id,_ := result.LastInsertId()
	raise(err)
	return int32(id)
}

func (dao *AllowIpDAO) DelAllowIp(id int32) {
	sqlStr := "delete from manage_login_ip where id = ?"
	_, err := dao.db.Exec(sqlStr, id)
	raise(err)
}

func (dao *AllowIpDAO) GetAllowIp() ([]dataobject.AllowIp, int32){
	res := make([]dataobject.AllowIp, 0)
	sqlStr := "select id, operator, ip, add_time from manage_login_ip order by id desc"
	rows, err  := dao.db.Queryx(sqlStr)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var ip dataobject.AllowIp
		err = rows.StructScan(&ip)
		raise(err)
		res = append(res, ip)
	}

	sqlCount := "select count(*) from manage_login_ip"
	row := dao.db.QueryRowx(sqlCount)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}