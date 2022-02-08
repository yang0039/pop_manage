package mysql_dao

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
)

type ReportDAO struct {
	db *sqlx.DB
}

func NewReportDAO(db *sqlx.DB) *ReportDAO {
	return &ReportDAO{db}
}

func (dao *ReportDAO) GetReports(limit, offset int32) []*dataobject.ReportDO {
	res := make([]*dataobject.ReportDO, 0)

	sqlStr := "select id, user_id, peer_type, peer_id, msg_ids, reason, content, add_time from reports order by add_time desc limit ? offset ?;"
	rows, err := dao.db.Queryx(sqlStr, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var re dataobject.ReportDO
		err = rows.StructScan(&re)
		raise(err)
		res = append(res, &re)
	}
	return res
}

func (dao *ReportDAO) GetReportCount() int32 {
	qryCount := "select count(*) from reports;"
	row := dao.db.QueryRow(qryCount)
	var count int32
	err := row.Scan(&count)
	raise(err)
	return count
}
