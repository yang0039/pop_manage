package mysql_dao

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
)

type BannedInfoDAO struct {
	db *sqlx.DB
}

func NewBannedInfoDAO(db *sqlx.DB) *BannedInfoDAO {
	return &BannedInfoDAO{db}
}

func (dao *BannedInfoDAO) AddbannedInfo(info *dataobject.BannedInfo) int32 {
	sqlStr := "insert into manage_banned_info (user_id, auth_id, model, system_version, app_version, system_lang_code, lang_pack, lang_code, ip, layer, date_created, date_activate)" +
		" values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	res, err := dao.db.Exec(sqlStr, info.UserId, info.AuthId, info.Model, info.SystemVersion, info.AppVersion, info.SystemLangCode, info.LangPack, info.LangCode, info.Ip, info.Layer, info.DateCreated, info.DateActivate)
	raise(err)
	id,_ := res.LastInsertId()
	return int32(id)
}

func (dao *BannedInfoDAO) GetUerBannedInfo(userId int32) []*dataobject.BannedInfo {
	res := make([]*dataobject.BannedInfo, 0)
	sqlStr := "select * from manage_banned_info where user_id = ?;"
	rows, err := dao.db.Queryx(sqlStr, userId)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var info  dataobject.BannedInfo
		err = rows.StructScan(&info)
		raise(err)
		res = append(res, &info)
	}
	return res
}