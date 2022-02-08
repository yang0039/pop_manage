package mysql_dao

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
)

type UserMsgRowDAO struct {
	db *sqlx.DB
}

func NewUserMsgRowDAO(db *sqlx.DB) *UserMsgRowDAO {
	return &UserMsgRowDAO{db}
}

func (dao *UserMsgRowDAO) GetUserMsgRows(userId, peerId, peerType, msgType int32, minTime, maxTime int64, limit, offset int32) ([]*dataobject.UserMsgRow, int32) {
	if maxTime == 0 {
		maxTime = 2147483647
	}

	var table string
	if Dbindex == 1 {
		table = "user_msgbox"
	} else {
		table = "user_msgbox1"
		if userId %2 == 0 {
			table = "user_msgbox2"
		}
	}

	var typeStr string
	if msgType == -1 {
		// 全部， 不用加条件
	} else if msgType == -2 {
		typeStr = " and type not in (0,1,2,3,4,5,6,7,11,12,13) "
	} else {
		typeStr = fmt.Sprintf(" and type = %d ", msgType)
	}

	res := make([]*dataobject.UserMsgRow, 0, limit)
	sqlStr := fmt.Sprintf(`
	select *
	from %s 
	where user_id = ? and peer_type = ? and peer_id = ? %s and add_time between ? and ?
	order by msg_id desc 
	limit ? offset ?;`, table, typeStr)

	fmt.Println("")

	rows, err := dao.db.Queryx(sqlStr, userId, peerType, peerId, minTime, maxTime, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var m dataobject.UserMsgRow
		err = rows.StructScan(&m)
		raise(err)
		res = append(res, &m)
	}

	sqlCount := fmt.Sprintf(`
	select count(*) count
	from %s
	where user_id = ? and peer_type = ? and peer_id = ?;
	`, table)
	row := dao.db.QueryRowx(sqlCount, userId, peerType, peerId)
	var count int32
	row.Scan(&count)
	return res, count
}


func (dao *UserMsgRowDAO)GetUserMsgRowsById(userId int32, msgIds []int32) []*dataobject.UserMsgRow {
	var table string
	if Dbindex == 1 {
		table = "user_msgbox"
	} else {
		table = "user_msgbox1"
		if userId %2 == 0 {
			table = "user_msgbox2"
		}
	}

	res := make([]*dataobject.UserMsgRow, 0)
	sqlStr := fmt.Sprintf(`
	select
		id, user_id, msg_id, pts, from_msg_id, raw_id, type, from_id, peer_type,
  		peer_id, reply_to_msg_id, mentioned, media_unread, add_time
	from %s 
	where user_id = ? and msg_id in (?);`, table)

	query, args, err := sqlx.In(sqlStr, userId, msgIds)
	raise(err)
	rows, err := dao.db.Queryx(query, args...)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var m dataobject.UserMsgRow
		err = rows.StructScan(&m)
		raise(err)
		res = append(res, &m)
	}
	return res
}