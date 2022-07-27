package mysql_dao

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
)

type ChannelMsgRowDAO struct {
	db *sqlx.DB
}

func NewChannelMsgRowDAO(db *sqlx.DB) *ChannelMsgRowDAO {
	return &ChannelMsgRowDAO{db}
}

func (dao *ChannelMsgRowDAO)GetChannelMsgRows(chatId, minTime, maxTime int64, limit, offset int32) ([]*dataobject.ChannelMsgRow, int32) {
	if maxTime == 0 {
		maxTime = 2147483647
	}
	var table string
	if Dbindex == 1 {
		table = "channel_msgbox"
	} else {
		table = "channel_msgbox1"
		if chatId %2 == 0 {
			table = "channel_msgbox2"
		}
	}

	res := make([]*dataobject.ChannelMsgRow, 0, limit)
	sqlStr := fmt.Sprintf(`
	select
		chat_id, msg_id, raw_id, type, from_id,
		reply_to_msg_id, media_unread, add_time
	from %s 
	where chat_id = ? and add_time between ? and ?
	order by msg_id desc 
	limit ? offset ?;`, table)
	rows, err := dao.db.Queryx(sqlStr, chatId, minTime, maxTime, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var m dataobject.ChannelMsgRow
		err = rows.StructScan(&m)
		raise(err)
		res = append(res, &m)
	}

	sqlCount := fmt.Sprintf(`
	select count(*) count
	from %s 
	where chat_id = ?;`, table)
	row := dao.db.QueryRowx(sqlCount, chatId)
	var count int32
	row.Scan(&count)
	return res, count
}


func (dao *ChannelMsgRowDAO)GetChannelMsgRowsByFrom(chatId , msgType int32, minTime, maxTime int64, limit, offset int32) ([]*dataobject.ChannelMsgRow, int32) {
	if maxTime == 0 {
		maxTime = 2147483647
	}
	var table string
	if Dbindex == 1 {
		table = "channel_msgbox"
	} else {
		table = "channel_msgbox1"
		if chatId %2 == 0 {
			table = "channel_msgbox2"
		}
	}

	var typeStr string
	if msgType == -1 {
		// 全部， 不用加条件
	} else if msgType == -2 {
		typeStr = " and type not in (0,1,2,3,4,5,6,7,11,12,13) "
	} else {
		typeStr = fmt.Sprintf(" and type = %s ", msgType)
	}

	res := make([]*dataobject.ChannelMsgRow, 0, limit)
	sqlStr := fmt.Sprintf(`
	select
		chat_id, msg_id, raw_id, type, from_id,
		reply_to_msg_id, media_unread, add_time
	from %s 
	where chat_id = ? %s and add_time between ? and ?
	order by msg_id desc 
	limit ? offset ?;`, table, typeStr)
	rows, err := dao.db.Queryx(sqlStr, chatId, minTime, maxTime, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var m dataobject.ChannelMsgRow
		err = rows.StructScan(&m)
		raise(err)
		res = append(res, &m)
	}

	sqlCount := fmt.Sprintf(`
	select count(*) count
	from %s 
	where chat_id = ?;`, table)
	row := dao.db.QueryRowx(sqlCount, chatId)
	var count int32
	row.Scan(&count)
	return res, count
}

func (dao *ChannelMsgRowDAO)GetChannelMsgRowsById(chatId int32, msgIds []int32) []*dataobject.ChannelMsgRow {
	var table string
	if Dbindex == 1 {
		table = "channel_msgbox"
	} else {
		table = "channel_msgbox1"
		if chatId %2 == 0 {
			table = "channel_msgbox2"
		}
	}

	res := make([]*dataobject.ChannelMsgRow, 0)
	if len(msgIds) == 0 {
		return res
	}
	sqlStr := fmt.Sprintf(`
	select
		chat_id, msg_id, raw_id, type, from_id,
		reply_to_msg_id, media_unread, add_time
	from %s 
	where chat_id = ? and msg_id in (?);`, table)

	query, args, err := sqlx.In(sqlStr, chatId, msgIds)
	raise(err)
	rows, err := dao.db.Queryx(query, args...)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var m dataobject.ChannelMsgRow
		err = rows.StructScan(&m)
		raise(err)
		res = append(res, &m)
	}
	return res
}

func (dao *ChannelMsgRowDAO)GetFileChannelMsgRows(userId, chatId int32, start, end int64) []*dataobject.ChannelMsgRow {
	var table string
	if Dbindex == 1 {
		table = "channel_msgbox"
	} else {
		table = "channel_msgbox1"
		if chatId %2 == 0 {
			table = "channel_msgbox2"
		}
	}
	condSql := ""
	if userId != 0 {
		condSql += fmt.Sprintf("and from_id = %d ", userId)
	}
	if start > 0 {
		condSql += fmt.Sprintf("and add_time > %d ", start)
	}
	if end > 0 {
		condSql += fmt.Sprintf("and add_time < %d", end)
	}

	res := make([]*dataobject.ChannelMsgRow, 0)
	sqlStr := fmt.Sprintf(`
	select
		chat_id, msg_id, raw_id, type, from_id,
		reply_to_msg_id, media_unread, add_time
	from %s 
	where chat_id = ? and type in (1,2,3,5,6,7,11) %s;`, table, condSql)

	rows, err := dao.db.Queryx(sqlStr, chatId)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var m dataobject.ChannelMsgRow
		err = rows.StructScan(&m)
		raise(err)
		res = append(res, &m)
	}
	return res
}

func (dao *ChannelMsgRowDAO)GetChannelFileMsgById(chatId, msgId int32) *dataobject.ChannelMsgRow {
	var table string
	if Dbindex == 1 {
		table = "channel_msgbox"
	} else {
		table = "channel_msgbox1"
		if chatId %2 == 0 {
			table = "channel_msgbox2"
		}
	}
	sqlStr := fmt.Sprintf(`
	select
		chat_id, msg_id, raw_id, type, from_id,
		reply_to_msg_id, media_unread, add_time
	from %s 
	where chat_id = ? and msg_id = ? and type in (1,2,3,5,6,7,11);`, table)
	row := dao.db.QueryRowx(sqlStr, chatId, msgId)
	var msg dataobject.ChannelMsgRow
	err := row.StructScan(&msg)
	if err == sql.ErrNoRows {
		return &msg
	}
	raise(err)
	return &msg
}

func (dao *ChannelMsgRowDAO) GetChannelFileMsgRowsByType(userId int32, chatIds, types []int32, start, end int64, offset, limit int32) ([]*dataobject.ChannelMsgRow, int32) {
	res := make([]*dataobject.ChannelMsgRow, 0)
	var timeSql string
	if start > 0 {
		timeSql = fmt.Sprintf(" and add_time >= %d", start)
	}
	if end > 0 {
		timeSql += fmt.Sprintf(" and add_time <= %d", end)
	}
	if len(chatIds) == 0 {
		return res, 0
	}
	var query string
	var args []interface{}
	var err error
	if Dbindex == 1 {
		sqlStr := `
		select * 
		from (
			select chat_id, msg_id, raw_id, type, from_id, reply_to_msg_id, media_unread, add_time from channel_msgbox where chat_id in (?) and from_id = ? and type in (?)` + timeSql + `
		) data order by add_time desc limit ? offset ?;
		`
		query, args, err = sqlx.In(sqlStr, chatIds, userId, types, limit, offset)
		raise(err)
	} else {
		sqlStr := `
		select * 
		from (
			select chat_id, msg_id, raw_id, type, from_id, reply_to_msg_id, media_unread, add_time from channel_msgbox1 where chat_id in (?) and from_id = ? and type in (?)` + timeSql + `
			union all
			select chat_id, msg_id, raw_id, type, from_id, reply_to_msg_id, media_unread, add_time from channel_msgbox2 where chat_id in (?) and from_id = ? and type in (?)` + timeSql + `
		) data order by add_time desc limit ? offset ?;
		`
		query, args, err = sqlx.In(sqlStr, chatIds, userId, types, chatIds, userId, types, limit, offset)
		raise(err)
	}

	rows, err := dao.db.Queryx(query, args...)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var m dataobject.ChannelMsgRow
		err = rows.StructScan(&m)
		raise(err)
		res = append(res, &m)
	}

	// 查找数量
	var count int32
	if Dbindex == 1 {
		sqlCount := `select count(*) from channel_msgbox where chat_id in (?) and from_id = ? and type in (?)` + timeSql
		queryCount, args, err := sqlx.In(sqlCount, chatIds, userId, types)
		raise(err)
		row := dao.db.QueryRow(queryCount, args...)
		err = row.Scan(&count)
	} else {
		sqlCount := `
		select sum(c) count 
		from (
			select count(*) c from channel_msgbox1 where chat_id in (?) and from_id = ? and type in (?)` + timeSql + `
			union all 
			select count(*) c from channel_msgbox2 where chat_id in (?) and from_id = ? and type in (?)` + timeSql + `
		) cc
		`

		queryCount, args, err := sqlx.In(sqlCount, chatIds, userId, types, chatIds, userId, types)
		raise(err)
		row := dao.db.QueryRow(queryCount, args...)
		err = row.Scan(&count)
	}
	return res, count
}


