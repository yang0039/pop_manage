package mysql_dao

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type ChatParticipantDAO struct {
	db *sqlx.DB
}

func NewChatParticipartDAO(db *sqlx.DB) *ChatParticipantDAO {
	return &ChatParticipantDAO{db}
}

// 获取含三人以上群组
func (dao *ChatParticipantDAO) GetChatNumIds(num, limit, offset int32) (map[int32]int32, int32) {
	chatNum := make(map[int32]int32)
	qry := "select * from (select chat_id, count(chat_id) user_num from chat_participant p left join chat c on c.id = p.chat_id where c.type in (1,2) and deactivated = 0 group by chat_id) c where user_num >= ? limit ? offset ?;"
	rows, err := dao.db.Queryx(qry, num, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return chatNum, 0
	}
	raise(err)
	for rows.Next() {
		var chatId, num int32
		err = rows.Scan(&chatId, &num)
		raise(err)
		chatNum[chatId] = num
	}

	qryCount := `
	select count(*) num from (select chat_id, count(chat_id) user_num from chat_participant p left join chat c on c.id = p.chat_id where c.type in (1,2) and deactivated = 0 group by chat_id) c where user_num >= ?;
	`
	row := dao.db.QueryRow(qryCount, num)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return chatNum, count
}

// 通过拥有者id获取群id
func (dao *ChatParticipantDAO) GetChatIdsByCreator(creatId, limit, offset int32, chatType []int32) ([]int32, int32) {
	chatIds := make([]int32, 0)
	qry := `
	select chat_id
	from chat_participant p
	left join chat c on p.chat_id = c.id
	where deactivated = 0 and c.type in (?)
	and user_id = ? and p.type = 1 limit ? offset ?;
	`
	q, args, err := sqlx.In(qry, chatType, creatId, limit, offset)
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

	qryCount := `
	select count(*)
	from chat_participant p
	left join chat c on p.chat_id = c.id
	where deactivated = 0 and c.type in (?)
	and user_id = ? and p.type = 1
	`
	q2, args, err := sqlx.In(qryCount, chatType, creatId)
	raise(err)
	row := dao.db.QueryRow(q2, args...)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return chatIds, count
}

func (dao *ChatParticipantDAO) GetChatByManage(manId, limit, offset int32, chatType []int32) ([]int32, int32) {
	chatIds := make([]int32, 0)
	qry := `
	select chat_id
	from chat_participant p
	left join chat c on p.chat_id = c.id
	where deactivated = 0 and p.type = 2 and c.type in (?) and user_id = ?
	group by chat_id limit ? offset ?;
	`
	q, args, err := sqlx.In(qry, chatType, manId, limit, offset)
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

	qryCount := `
	select count(*) count
	from chat_participant p
	left join chat c on p.chat_id = c.id
	where deactivated = 0 and p.type = 2 and c.type in (?) and user_id = ?;
	`
	q2, args, err := sqlx.In(qryCount, chatType, manId)
	raise(err)
	row := dao.db.QueryRow(q2, args...)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return chatIds, count
}

func (dao *ChatParticipantDAO) GetChatByMember(manId, limit, offset int32, chatType []int32) ([]int32, int32) {
	chatIds := make([]int32, 0)
	qry := `
	select chat_id
	from chat_participant p
	left join chat c on p.chat_id = c.id
	where deactivated = 0 and c.type in (?) and user_id = ?
	group by chat_id limit ? offset ?;
	`
	q, args, err := sqlx.In(qry, chatType, manId, limit, offset)
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

	qryCount := `
	select count(*) count
	from chat_participant p
	left join chat c on p.chat_id = c.id
	where deactivated = 0 and c.type in (?) and user_id = ?;
	`
	q2, args, err := sqlx.In(qryCount, chatType, manId)
	raise(err)
	row := dao.db.QueryRow(q2, args...)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return chatIds, count
}

func (dao *ChatParticipantDAO) GetChatPart(manId, limit, offset int32, chatType []int32) ([]int32, int32) {
	chatIds := make([]int32, 0)
	qry := `
	select chat_id
	from chat_participant p
	left join chat c on p.chat_id = c.id
	where deactivated = 0 and p.type = 0 and c.type in (?) and user_id = ?
	group by chat_id limit ? offset ?;
	`
	q, args, err := sqlx.In(qry, chatType, manId, limit, offset)
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

	qryCount := `
	select count(*) count
	from chat_participant p
	left join chat c on p.chat_id = c.id
	where deactivated = 0 and p.type = 0 and c.type in (?) and user_id = ?;
	`
	q2, args, err := sqlx.In(qryCount, chatType, manId)
	raise(err)
	row := dao.db.QueryRow(q2, args...)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return chatIds, count
}

func (dao *ChatParticipantDAO) GetChatByMemNum(min, max, limit, offset int32, chatType []int32) ([]int32, int32) {
	chatIds := make([]int32, 0)
	qry := `
	select chat_id, count
	from
	(
	  select chat_id, count(*) count
	  from chat_participant p
	  left join chat c on p.chat_id = c.id
	  where deactivated = 0 and c.type in (?)
	  group by chat_id
	) d
	where count between ? and ? 
	order by count limit ? offset ?;
	`
	q, args, err := sqlx.In(qry, chatType, min, max, limit, offset)
	raise(err)
	rows, err := dao.db.Queryx(q, args...)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return chatIds, 0
	}
	raise(err)
	for rows.Next() {
		var chatId, count int32
		err = rows.Scan(&chatId, &count)
		raise(err)
		chatIds = append(chatIds, chatId)
	}

	qryCount := `
	select count(*) num
	from
	(
	  select chat_id, count(*) count
	  from chat_participant p
	  left join chat c on p.chat_id = c.id
	  where deactivated = 0 and c.type in (?)
	  group by chat_id
	) d
	where count between ? and ?;
	`
	q2, args, err := sqlx.In(qryCount, chatType, min, max)
	raise(err)
	row := dao.db.QueryRow(q2, args...)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return chatIds, count
}

// 获取用户所在群组数量，拥有的群数量，管理的群数量
func (dao *ChatParticipantDAO) GetUserChatNum(userIds []int32) map[int32]map[string]int32 {
	res := make(map[int32]map[string]int32, 0)
	if len(userIds) == 0 {
		return res
	}
	queryType := "("
	for i, id := range userIds {
		queryType += strconv.Itoa(int(id))
		if i == len(userIds)-1 {
			queryType += ")"
		} else {
			queryType += ","
		}
	}
	qry := fmt.Sprintf(`
	select
	  user_id,
	  sum(
		case when p.type = 0 then 1
		  else 0 end
	  ) as normal_num,
	  sum(
		case when p.type = 1 then 1
		  else 0 end
	  ) as create_num,
	  sum(
		case when p.type = 2 then 1
		  else 0 end
	  ) as manage_num
	from chat_participant p
	left join chat c on p.chat_id = c.id
	where c.deactivated = 0 and c.type != 4 and user_id in %s
	group by user_id;
	`, queryType)
	rows, err := dao.db.Queryx(qry)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var userId, nonalNum, createNum, manageNum int32
		err = rows.Scan(&userId, &nonalNum, &createNum, &manageNum)
		raise(err)
		m := map[string]int32{
			"normal_num": nonalNum,
			"create_num": createNum,
			"manage_num": manageNum,
		}
		res[userId] = m
	}
	return res
}

func (dao *ChatParticipantDAO) GetCommonChatsCount(a_id, b_id int32) int32 {
	/* 返回两个用户共同群组数 */
	query := `select count(*) from chat_participant p1
	inner join chat_participant p2 on p1.chat_id=p2.chat_id and p2.kicked=0
	inner join chat c on c.id=p1.chat_id
	where p1.user_id=? and p2.user_id=? and p1.kicked=0 and c.deactivated=0
	and c.type != 4`
	row := dao.db.QueryRow(query, a_id, b_id)
	var count int32
	err := row.Scan(&count)
	raise(err)
	return count
}