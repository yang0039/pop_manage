package mysql_dao

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
)

type ReportDAO struct {
	db *sqlx.DB
}

func NewReportDAO(db *sqlx.DB) *ReportDAO {
	return &ReportDAO{db}
}


func (dao *ReportDAO) GetReports(reason, limit, offset int32) ([]*dataobject.ReportDO, int32) {
	// 0:全部 1：垃圾 2：暴力 3：色情 4：虐待 5：版权 6：其他
	res := make([]*dataobject.ReportDO, 0)
	var reasonSql string
	if reason == 0 {
	} else {
		if reason == 6 {
			reason = 0
		}
		reasonSql = fmt.Sprintf("where reason = %d", reason)
	}
	sqlStr := "select id, user_id, peer_type, peer_id, msg_ids, reason, content, add_time from reports " + reasonSql + " order by add_time desc limit ? offset ?;"
	rows, err := dao.db.Queryx(sqlStr, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var re dataobject.ReportDO
		err = rows.StructScan(&re)
		raise(err)
		res = append(res, &re)
	}

	qryCount := "select count(*) from reports " + reasonSql
	row := dao.db.QueryRow(qryCount)
	var count int32
	err = row.Scan(&count)
	raise(err)

	return res, count
}

//func (dao *ReportDAO) GetReportCount() int32 {
//	qryCount := "select count(*) from reports;"
//	row := dao.db.QueryRow(qryCount)
//	var count int32
//	err := row.Scan(&count)
//	raise(err)
//	return count
//}
