package mysql_dao

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
	"strconv"
)

type ChatDAO struct {
	db *sqlx.DB
}

func NewChatDAO(db *sqlx.DB) *ChatDAO {
	return &ChatDAO{db}
}

// 获取总群数
func (dao *ChatDAO) GetChatNum(start, end int64) (num int32) {
	var sql = "select count(*) count from chat where deactivated = 0 and add_time between ? and ?;"
	row := dao.db.QueryRowx(sql, start, end)
	err := row.Scan(&num)
	raise(err)
	return
}

// 获取指定的群
func (dao *ChatDAO) GetChats(ids []int32) []*dataobject.Chat {
	chats := make([]*dataobject.Chat, 0, len(ids))
	qry := `
	select
		id, creator_id, pinned_msg_id, about, title, type, photo_id, admins_enabled, migrated_from_chat_id, migrated_from_max_id, migrated_to,
		username, hidden_prehistory, signatures, verified, rights_mask, until_date, slowmode_seconds, sticker_set_id, add_time
	from chat
	where deactivated = 0 
	`
	queryType := "and id in ("
	for i, id := range ids {
		queryType += strconv.Itoa(int(id))
		if i == len(ids)-1 {
			queryType += ");"
		} else {
			queryType += ","
		}
	}
	rows, err := dao.db.Queryx(qry)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return chats
	}
	raise(err)

	for rows.Next() {
		var chat *dataobject.Chat
		err = rows.StructScan(chat)
		raise(err)
		chats = append(chats, chat)
	}
	return chats
}

// 获取指定的群
func (dao *ChatDAO) GetChat(id int32) *dataobject.Chat {
	qry := `
	select
		id, creator_id, pinned_msg_id, about, title, type, photo_id, admins_enabled, migrated_from_chat_id, migrated_from_max_id, migrated_to,
		username, hidden_prehistory, signatures, verified, rights_mask, until_date, slowmode_seconds, sticker_set_id, add_time
	from chat
	where deactivated = 0 and id = ?
	`
	row := dao.db.QueryRowx(qry, id)
	var c dataobject.Chat
	err := row.StructScan(&c)
	if err == sql.ErrNoRows {
		return &c
	}
	raise(err)
	return &c
}

// 获取特定时间内创建的人数超过指定人数的群
/*
func (dao *ChatDAO) GetMemberNumChats(start, end int64, num int32) []*dataobject.Chat {
	chats := make([]*dataobject.Chat, 0)
	sql := `
	select
	  id, creator_id, pinned_msg_id, about, title, type, photo_id, admins_enabled, migrated_from_chat_id, migrated_from_max_id, migrated_to,
		username, hidden_prehistory, signatures, verified, rights_mask, until_date, slowmode_seconds, sticker_set_id, add_time
	from
	(
	  select chat_id, count(chat_id) user_num
	  from chat_participant group by chat_id
	) cp
	right join
	(
	  select * from chat
	  where deactivated = 0
	  and add_time between ? and ?
	) c on cp.chat_id = c.id
	where user_num >= ?;
	`
	rows, err := dao.db.Queryx(sql, start, end, num)
	raise(err)
	defer rows.Close()
	for rows.Next() {
		var chat *dataobject.Chat
		err = rows.StructScan(chat)
		raise(err)
		chats = append(chats, chat)
	}
	return chats
}
*/

// 通过群名称获取群id
func (dao *ChatDAO) GetChatIdsByTitle(title string, chatType []int32, limit, offset int32) ([]int32, int32){
	chatIds := make([]int32, 0)
	title = "%" + title + "%"
	qry := fmt.Sprintf("select id from chat where deactivated = 0 and type in (?) and title like '%s' limit ? offset ?;", title)
	q, args, err := sqlx.In(qry, chatType, limit, offset)
	raise(err)
	rows, err := dao.db.Queryx(q, args...)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return chatIds, 0
	}
	raise(err)
	for rows.Next() {
		var chatId int32
		err = rows.Scan(&chatId)
		raise(err)
		chatIds = append(chatIds, chatId)
	}

	qryCount := fmt.Sprintf("select count(*) from chat where deactivated = 0 and type in (?) and title like '%s' ;", title)
	q2, args, err := sqlx.In(qryCount, chatType)
	raise(err)
	row := dao.db.QueryRow(q2, args...)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return chatIds, count
}

// 通过拥有者id获取群id
func (dao *ChatDAO) GetChatIdsDefault(chatType []int32, limit, offset int32) ([]int32, int32) {
	chatIds := make([]int32, 0)
	qry :="select c.id from chat c left join chat_participant p on c.id = p.chat_id and c.creator_id = p.user_id where deactivated = 0 and c.type in (?) and p.user_id is not null order by c.add_time desc limit ? offset ?;"
	q, args, err := sqlx.In(qry, chatType, limit, offset)
	raise(err)
	rows, err := dao.db.Queryx(q, args...)
	//rows, err := dao.db.Queryx(qry, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return chatIds, 0
	}
	raise(err)
	for rows.Next() {
		var chatId int32
		err = rows.Scan(&chatId)
		raise(err)
		chatIds = append(chatIds, chatId)
	}

	qryCount := "select count(*) from chat c left join chat_participant p on c.id = p.chat_id and c.creator_id = p.user_id where deactivated = 0 and c.type in (?) and p.user_id is not null;"
	q2, args, err := sqlx.In(qryCount, chatType)
	raise(err)
	row := dao.db.QueryRow(q2, args...)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return chatIds, count
}

// 通过创建时间获取群id
func (dao *ChatDAO) GetChatIdsByCreate(start, end int64, limit, offset int32, chatType []int32) ([]int32, int32) {
	chatIds := make([]int32, 0)
	qry :="select c.id from chat c left join chat_participant p on c.id = p.chat_id and c.creator_id = p.user_id where deactivated = 0 and c.type in (?) and p.user_id is not null and c.add_time between ? and ? limit ? offset ?;"
	q, args, err := sqlx.In(qry, chatType, start, end, limit, offset)
	raise(err)
	rows, err := dao.db.Queryx(q, args...)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return chatIds, 0
	}
	raise(err)
	for rows.Next() {
		var chatId int32
		err = rows.Scan(&chatId)
		raise(err)
		chatIds = append(chatIds, chatId)
	}

	qryCount := "select count(*) from chat c left join chat_participant p on c.id = p.chat_id and c.creator_id = p.user_id where deactivated = 0  and c.type in (?) and p.user_id is not null and c.add_time between ? and ?;"
	q2, args, err := sqlx.In(qryCount, chatType, start, end)
	raise(err)
	row := dao.db.QueryRow(q2, args...)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return chatIds, count
}