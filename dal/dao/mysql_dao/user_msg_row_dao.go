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
	select 
		id, user_id, msg_id, pts, from_msg_id, raw_id, type, from_id, peer_type,
  		peer_id, reply_to_msg_id, mentioned, media_unread, add_time
	from %s 
	where user_id = ? and peer_type = ? and peer_id = ? %s and add_time between ? and ?
	order by msg_id desc 
	limit ? offset ?;`, table, typeStr)

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

func (dao *UserMsgRowDAO)GetUserMsgRowsByRawId(userId int32, rawIds []int64) []*dataobject.UserMsgRow {
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
	if len(rawIds) == 0 {
		return res
	}
	sqlStr := fmt.Sprintf(`
	select
		id, user_id, msg_id, pts, from_msg_id, raw_id, type, from_id, peer_type,
  		peer_id, reply_to_msg_id, mentioned, media_unread, add_time
	from %s 
	where user_id = ? and raw_id in (?);`, table)

	query, args, err := sqlx.In(sqlStr, userId, rawIds)
	raise(err)
	rows, err := dao.db.Queryx(query, args...)
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
	if len(msgIds) == 0 {
		return res
	}
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
	raise(err)
	for rows.Next() {
		var m dataobject.UserMsgRow
		err = rows.StructScan(&m)
		raise(err)
		res = append(res, &m)
	}
	return res
}

func (dao *UserMsgRowDAO) GetUserFileMsgRows(userId, peerType, peerId int32, start, end int64) []*dataobject.UserMsgRow {
	var table string
	if Dbindex == 1 {
		table = "user_msgbox"
	} else {
		table = "user_msgbox1"
		if userId %2 == 0 {
			table = "user_msgbox2"
		}
	}
	condSql := ""
	if peerType != 0 {     // 个人消息
		condSql += fmt.Sprintf("and peer_type = %d", peerType)
	}
	if peerId != 0 {
		condSql += fmt.Sprintf(" and peer_id = %d", peerId)
	}

	if start > 0 {
		condSql += fmt.Sprintf(" and add_time > %d", start)
	}
	if end > 0 {
		condSql += fmt.Sprintf(" and add_time < %d", end)
	}

	res := make([]*dataobject.UserMsgRow, 0)
	sqlStr := fmt.Sprintf(`
	select 
		id, user_id, msg_id, pts, from_msg_id, raw_id, type, from_id, peer_type,
  		peer_id, reply_to_msg_id, mentioned, media_unread, add_time
	from %s 
	where user_id = ? %s and type in (1,2,3,5,6,7,11);`, table, condSql)
	fmt.Println(sqlStr)
	rows, err := dao.db.Queryx(sqlStr, userId)
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

func (dao *UserMsgRowDAO) GetUserFileMsgById(userId, msgId int32) *dataobject.UserMsgRow{
	var table string
	if Dbindex == 1 {
		table = "user_msgbox"
	} else {
		table = "user_msgbox1"
		if userId %2 == 0 {
			table = "user_msgbox2"
		}
	}
	sqlStr := fmt.Sprintf(`
	select 
		id, user_id, msg_id, pts, from_msg_id, raw_id, type, from_id, peer_type,
  		peer_id, reply_to_msg_id, mentioned, media_unread, add_time
	from %s 
	where user_id = ? and msg_id = ? and type in (1,2,3,5,6,7,11);`, table)
	row := dao.db.QueryRowx(sqlStr, userId, msgId)
	var msg dataobject.UserMsgRow
	err := row.StructScan(&msg)
	if err == sql.ErrNoRows {
		return &msg
	}
	raise(err)
	return &msg
}

func (dao *UserMsgRowDAO) GetUserFileMsgRowsByType(userId, peerType int32, types []int32, start, end int64, offset, limit int32) ([]*dataobject.UserMsgRow, int32) {
	var table string
	if Dbindex == 1 {
		table = "user_msgbox"
	} else {
		table = "user_msgbox1"
		if userId %2 == 0 {
			table = "user_msgbox2"
		}
	}

	var timeSql string
	if start > 0 {
		timeSql = fmt.Sprintf(" and add_time >= %d", start)
	}
	if end > 0 {
		timeSql += fmt.Sprintf(" and add_time <= %d", end)
	}

	res := make([]*dataobject.UserMsgRow, 0)
	sqlStr := fmt.Sprintf(`
	select 
		id, user_id, msg_id, pts, from_msg_id, raw_id, type, from_id, peer_type,
  		peer_id, reply_to_msg_id, mentioned, media_unread, add_time
	from %s 
	where user_id = ?  and peer_type = ?  and from_id = ? %s and type in (?) order by id desc limit ? offset ?;`, table, timeSql)
	query, args, err := sqlx.In(sqlStr, userId, peerType, userId, types, limit, offset)
	raise(err)
	rows, err := dao.db.Queryx(query, args...)
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
	qryCount := fmt.Sprintf("select count(*) from %s where user_id = ? and peer_type = ? and from_id = ? %s and type in (?);", table, timeSql)
	queryCount, args, err := sqlx.In(qryCount, userId, peerType, userId, types)
	raise(err)
	row := dao.db.QueryRow(queryCount, args...)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count

}
