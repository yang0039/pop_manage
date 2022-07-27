package mysql_dao

import (
	"database/sql"
	"fmt"
)

/*
专门用来处理线上异常数据
*/



func (dao *CommonDAO) GetRepeatPeerData() []map[string]int32 {
	res := make([]map[string]int32, 0)
	var repeatSql = `
	select u1.user_id, u1.peer_id, u1.id id1, u2.id id2, u1.peer_type peer_type1, u2.peer_type peer_type2
	from
	(
	   select * from user_dialog where peer_type != 2
	) u1
	left join
	(
		select * from user_dialog where peer_type != 2
	) u2 on u1.user_id = u2.user_id and u1.peer_id = u2.peer_id
	where u1.peer_type != u2.peer_type and u1.peer_type = 3 order by u1.id desc;
	`
	rows, err := dao.db.Queryx(repeatSql)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var id1, id2, user_id, peer_type1, peer_type2, peer_id int32
		err = rows.Scan(&user_id, &peer_id, &id1, &id2, &peer_type1, &peer_type2)
		raise(err)
		m := map[string]int32{
			"user_id": user_id,
			"peer_id": peer_id,
			"id1": id1,
			"id2": id2,
			"peer_type1": peer_type1,
			"peer_type2": peer_type2,
		}
		res = append(res, m)
	}
	return res
}


func (dao *CommonDAO) DelRepeatPeerData(id int32) {
	sqlStr := "delete from user_dialog where id = ?"
	_, err := dao.db.Exec(sqlStr, id)
	raise(err)
}

func (dao *CommonDAO) GetAllUser(limit, offset int32) []int32 {
	ids := make([]int32, 0)
	var userSql = `select id from user where support = 0 order by id desc limit ? offset ?;`
	rows, err := dao.db.Queryx(userSql, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return ids
	}
	raise(err)
	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		raise(err)
		ids = append(ids, id)
	}
	return ids
}

func (dao *CommonDAO) GetAllChannel(limit, offset int32) []int32 {
	ids := make([]int32, 0)
	var userSql = `select id from chat where type in (2,3) order by id desc limit ? offset ?;`
	rows, err := dao.db.Queryx(userSql, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return ids
	}
	raise(err)
	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		raise(err)
		ids = append(ids, id)
	}
	return ids
}


func (dao *CommonDAO) GetMaxMsgId(userId int32) (int32, int32) {
	var table string
	if userId%2 == 0 {
		table = "user_msgbox2"
	} else {
		table = "user_msgbox1"
	}

	sqlStr := fmt.Sprintf("select msg_id, pts from %s where user_id = ? order by msg_id desc limit 1", table)
	row := dao.db.QueryRow(sqlStr, userId)
	var msgId, pts int32
	err := row.Scan(&msgId, &pts)
	if err == sql.ErrNoRows {
		return 0, 0
	}
	raise(err)
	return msgId, pts
}

func (dao *CommonDAO) GetChannelMaxMsgId(chatId int32) (int32) {
	var table string
	if chatId%2 == 0 {
		table = "channel_msgbox2"
	} else {
		table = "channel_msgbox1"
	}

	sqlStr := fmt.Sprintf("select msg_id from %s where chat_id = ? order by msg_id desc limit 1", table)
	row := dao.db.QueryRow(sqlStr, chatId)
	var msgId int32
	err := row.Scan(&msgId)
	if err == sql.ErrNoRows {
		return 0
	}
	raise(err)
	return msgId
}


