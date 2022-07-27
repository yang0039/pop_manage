package mysql_dao

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"pop-api/baselib/logger"
	"pop-api/baselib/util"
	"pop-api/dal/dataobject"
	"strconv"
	"time"
)

var Dbindex int

func raise(err error) {
	if err == nil {
		return
	}
	logger.Logger.Error(err.Error())
	panic(err.Error())
}

func raiseTx(err error, tx *sql.Tx) {
	if err == nil {
		return
	}
	tx.Rollback()
	logger.Logger.Error(err.Error())
	panic(err.Error())
}

type CommonDAO struct {
	db *sqlx.DB
}

func NewCommonDAO(db *sqlx.DB) *CommonDAO {
	return &CommonDAO{db}
}

func (dao *CommonDAO) GetActiveUserCount(start, end int64) int32 {
	qry := `select count(*) count from (select from_id from raw_message where add_time between ? and ? group by from_id) d;`
	row2 := dao.db.QueryRow(qry, start, end)
	var count int32
	err := row2.Scan(&count)
	raise(err)
	return count

	/*
	qry1 := `select count(*) count from (select user_id from user_msgbox1 where add_time between ? and ? group by user_id) d;`
	row := dao.db.QueryRow(qry1, start, end)
	var count1 int32
	err := row.Scan(&count1)
	raise(err)

	qry2 := `select count(*) count from (select user_id from user_msgbox2 where add_time between ? and ? group by user_id) d;`
	row2 := dao.db.QueryRow(qry2, start, end)
	var count2 int32
	err = row2.Scan(&count2)
	raise(err)
	*/

	//return count1 + count2
}

func (dao *CommonDAO) GetSendMsgCount() int32 {
	qry := `select count(*) count from raw_message;`
	row := dao.db.QueryRow(qry)
	var count int32
	err := row.Scan(&count)
	raise(err)
	return count
}

func (dao *CommonDAO) GetSendMsgCountWithTime(start, end int64) int32 {
	qry := `select count(*) count from raw_message where add_time between ? and ?;`
	row := dao.db.QueryRow(qry, start, end)
	var count int32
	err := row.Scan(&count)
	raise(err)
	return count
}

// 查找含特定数目以上人群数
func (dao *CommonDAO) GetMemberChatNum(start, end int64, num int32) int32 {
	sql := `
	select
	  count(*) count
	from
	(
	  select chat_id, count(chat_id) user_num
	  from chat_participant group by chat_id
	) cp
	right join
	(
	  select id from chat
	  where deactivated = 0
	  and add_time between ? and ?
	) c on cp.chat_id = c.id
	where user_num >= ?;
	`
	var n int32
	row := dao.db.QueryRowx(sql, start, end, num)
	err := row.Scan(&n)
	raise(err)
	return n
}

func (dao *CommonDAO) GetChatIsActivve(chatId int32, start, end int64) bool {
	qry := "select peer_id from raw_message where peer_type in (3,4) and peer_id = ? and add_time between ? and ? order by id desc limit 1;"
	var peerId int32
	row := dao.db.QueryRowx(qry, chatId, start, end)
	err := row.Scan(&peerId)
	if err == sql.ErrNoRows {
		return false
	}
	return peerId == chatId
}

func (dao *CommonDAO) GetActiveChatIds(start, end int64) []int32 {
	ids := make([]int32, 0)
	qry := "select peer_id from raw_message where peer_type in (3,4) and add_time between ? and ? group by peer_id;"
	rows, err := dao.db.Queryx(qry, start, end)
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

func (dao *CommonDAO) GetChatByActive(start, end int64, limit, offset int32, chatType []int32) ([]int32, int32) {
	ids := make([]int32, 0)
	qry := "select peer_id from raw_message r left join chat c on r.peer_id = c.id where peer_type in (3,4) and c.type in (?) and r.add_time between ? and ? group by peer_id limit ? offset ?;"
	q, args, err := sqlx.In(qry, chatType, start, end, limit, offset)
	raise(err)
	rows, err := dao.db.Queryx(q, args...)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return ids, 0
	}
	raise(err)
	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		raise(err)
		ids = append(ids, id)
	}

	qryCount := "select count(*) c from(select peer_id from raw_message r left join chat c on r.peer_id = c.id where peer_type in (3,4) and c.type in (?) and r.add_time between ? and ? group by peer_id) d;"
	q2, args, err := sqlx.In(qryCount, chatType, start, end)
	raise(err)
	row := dao.db.QueryRow(q2, args...)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return ids, count
}


func (dao *CommonDAO) GetUserByActive(start, end int64, limit, offset int32) ([]int32, int32) {
	ids := make([]int32, 0)
	qry := "select from_id from raw_message where add_time between ? and ? group by from_id limit ? offset ?;"
	rows, err := dao.db.Queryx(qry, start, end, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return ids, 0
	}
	raise(err)
	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		raise(err)
		ids = append(ids, id)
	}

	qryCount := "select count(*) c from(select from_id from raw_message where add_time between ? and ? group by from_id) d;"
	row := dao.db.QueryRow(qryCount, start, end)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return ids, count
}

// 指定群基本信息
func (dao *CommonDAO) GetChatBaseInfo(chatIds []int32) []map[string]interface{} {
	chats := make([]map[string]interface{}, 0)
	if len(chatIds) == 0 {
		return chats
	}
	sql2 := "and c.id in ("
	for i, id := range chatIds {
		sql2 = sql2 + fmt.Sprintf("%d", id)
		if i == len(chatIds)-1 {
			sql2 += ")"
		} else {
			sql2 += ","
		}
	}
	qry := fmt.Sprintf(`
	select
	  c.id chat_id, count(p.id) num,
      sum(
		  case ifnull(p.type,0)
		  when  0 then 0
		  else 1 end
      ) manage_num,
      c.title, c.type, c.creator_id, c.add_time
	from chat_participant p
	right join chat c on p.chat_id = c.id
	where c.deactivated = 0 and c.type in (1, 2, 3, 4)
	%s
	group by c.id order by num;
	`, sql2)
	rows, err := dao.db.Queryx(qry)
	defer rows.Close()


	if err == sql.ErrNoRows {
		return chats
	}
	raise(err)
	for rows.Next() {
		var chatId, num, manageNum, chatType, creatorId, addTime int32
		var title string
		m := make(map[string]interface{})
		err = rows.Scan(&chatId, &num, &manageNum, &title, &chatType, &creatorId, &addTime)
		raise(err)
		m["chat_id"] = chatId
		m["member_num"] = num
		m["manage_num"] = manageNum
		m["title"] = title
		m["chat_type"] = util.DbToApiChatType(chatType)
		m["creator_id"] = creatorId
		m["add_time"] = addTime
		chats = append(chats, m)
	}
	return chats
}

func (dao *CommonDAO) GetChatLastActiveDate(chatIds []int32) map[int32]int32 {
	chatDate := make(map[int32]int32)
	if len(chatIds) == 0 {
		return chatDate
	}
	sql2 := "("
	for i, id := range chatIds {
		sql2 = sql2 + fmt.Sprintf("%d", id)
		if i == len(chatIds)-1 {
			sql2 += ")"
		} else {
			sql2 += ","
		}
	}

	qry := fmt.Sprintf(`
	select peer_id chat_id, max(add_time) add_time
	from raw_message
	where peer_type in (3, 4) and peer_id in %s
	group by peer_id;
	`, sql2)
	rows, err := dao.db.Queryx(qry)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return chatDate
	}
	raise(err)

	for rows.Next() {
		var chatId, date int32
		err = rows.Scan(&chatId, &date)
		raise(err)
		chatDate[chatId] = date
	}
	return chatDate
}

func (dao *CommonDAO) GetChatMemberInfo(chatId int32) []*dataobject.ChatParticipantInfo {
	qry := `
	select
	  chat_id, user_id, type, p.add_time,
	  first_name, last_name, username, phone,
	  country_code, u.add_time register_time
	from chat_participant p
	left join user u on p.user_id = u.id
	where kicked = 0 and chat_id = ?;
	`
	rows, err := dao.db.Queryx(qry, chatId)
	defer rows.Close()
	chatMembers := make([]*dataobject.ChatParticipantInfo, 0)
	if err == sql.ErrNoRows {
		return chatMembers
	}
	raise(err)
	for rows.Next() {
		chatMember := &dataobject.ChatParticipantInfo{}
		err = rows.StructScan(chatMember)
		raise(err)
		chatMembers = append(chatMembers, chatMember)
	}
	return chatMembers
}

func (dao *CommonDAO) Get100ChatIds(limit, offset int32) []int32 {
	chatIds := make([]int32, 0)
	qry := `
	select
	  chat_id, count(*) num
	from chat_participant p
	left join chat c on p.chat_id = c.id
	where c.deactivated = 0 and c.type in (1, 2)
	group by chat_id order by num desc limit ? offset ?;
	`
	rows, err := dao.db.Queryx(qry, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return chatIds
	}
	raise(err)
	for rows.Next() {
		var chatId, num int32
		err = rows.Scan(&chatId, &num)
		raise(err)
		chatIds = append(chatIds, chatId)
	}
	return chatIds
}

func (dao *CommonDAO) GetUserActiveTime(userIds []int32) map[int32]int64 {
	res := make(map[int32]int64, 0)
	if len(userIds) == 0 {
		return res
	}
	sql2 := "("
	for i, id := range userIds {
		sql2 = sql2 + fmt.Sprintf("%d", id)
		if i == len(userIds)-1 {
			sql2 += ")"
		} else {
			sql2 += ","
		}
	}
	qry := fmt.Sprintf(`
	select peer_id user_id, max(add_time) add_time
	from raw_message
	where peer_type = 2 and peer_id in %s
	group by peer_id;
	`, sql2)

	logger.LogSugar.Infof("GetUserActiveTime sql:%s", qry)

	rows, err := dao.db.Queryx(qry)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var userId int32
		var lastTime int64
		err = rows.Scan(&userId, &lastTime)
		raise(err)
		res[userId] = lastTime
	}
	return res
}

func (dao *CommonDAO) GetUserEmail(userIds []int32) map[int32]string {
	res := make(map[int32]string, 0)
	if len(userIds) == 0 {
		return res
	}

	sql2 := "("
	for i, id := range userIds {
		sql2 = sql2 + fmt.Sprintf("%d", id)
		if i == len(userIds)-1 {
			sql2 += ")"
		} else {
			sql2 += ","
		}
	}

	qry := fmt.Sprintf(`
	select user_id, email
	from user_passwords
	where email != ''
	and user_id in %s;
	`, sql2)
	rows, err := dao.db.Queryx(qry)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var userId int32
		var email string
		err = rows.Scan(&userId, &email)
		raise(err)
		res[userId] = email
	}
	return res
}

func (dao *CommonDAO) GetUserIdsByEmail(email string, limit, offset int32) ([]int32, int32) {
	res := make([]int32, 0)
	if email == "" {
		return res, 0
	}
	qry := "select user_id from user_passwords where email = ? limit ? offset ?;"
	rows, err := dao.db.Queryx(qry, email, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)

	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		raise(err)
		res = append(res, id)
	}

	qryCount := "select count(*) from user_passwords where email = ?;"
	row := dao.db.QueryRow(qryCount, email)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}

func (dao *CommonDAO) GetAuthIds(userId int32) []int32 {
	res := make([]int32, 0)
	if userId == 0 {
		return res
	}
	qry := "select auth_id from auth_users where user_id = ?;"
	rows, err := dao.db.Queryx(qry, userId)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		raise(err)
		res = append(res, id)
	}
	return res
}


func (dao *CommonDAO) DeleteUser(userId int32, authIds []int32) (succ bool) {
	if userId == 0 || len(authIds) == 0 {
		return true
	}

	sql2 := "("
	for i, id := range authIds {
		sql2 = sql2 + fmt.Sprintf("%d", id)
		if i == len(authIds)-1 {
			sql2 += ")"
		} else {
			sql2 += ","
		}
	}

	tx, err := dao.db.Begin()
	raise(err)

	delAuthKey := fmt.Sprintf("delete from auth_keys where auth_id in %s", sql2)
	_, err = tx.Exec(delAuthKey)
	raiseTx(err, tx)

	delAuthPhone := fmt.Sprintf("delete from auth_phone_transactions where auth_key_id in %s", sql2)
	_, err = tx.Exec(delAuthPhone)
	raiseTx(err, tx)

	delbotCommand := "delete from bot_command where user_id = ?;"
	_, err = tx.Exec(delbotCommand, userId)
	raiseTx(err, tx)

	delchannelMen := "delete from channel_mention where user_id = ?;"
	_, err = tx.Exec(delchannelMen, userId)
	raiseTx(err, tx)

	delChatInvi := "delete from chat_invite where user_id = ?;"
	_, err = tx.Exec(delChatInvi, userId)
	raiseTx(err, tx)

	delChatPart := "delete from chat_participant where user_id = ?;"
	_, err = tx.Exec(delChatPart, userId)
	raiseTx(err, tx)

	delDevices := "delete from devices where user_id = ?;"
	_, err = tx.Exec(delDevices, userId)
	raiseTx(err, tx)

	delPhoneCall := "delete from phonecall where a_id = ? or b_id = ?;"
	_, err = tx.Exec(delPhoneCall, userId, userId)
	raiseTx(err, tx)

	delSecretChat := "delete from secret_chat where a_id = ? or b_id = ?;"
	_, err = tx.Exec(delSecretChat, userId, userId)
	raiseTx(err, tx)

	delUser := "delete from user where id = ?;"
	_, err = tx.Exec(delUser, userId)
	raiseTx(err, tx)

	delContact1 := "delete from user_contacts where owner_user_id = ?;"
	_, err = tx.Exec(delContact1, userId)
	raiseTx(err, tx)

	delContact2 := "delete from user_contacts where contact_user_id = ?;"
	_, err = tx.Exec(delContact2, userId)
	raiseTx(err, tx)

	delUserDialog := "delete from user_dialog where user_id = ?;"
	_, err = tx.Exec(delUserDialog, userId)
	raiseTx(err, tx)

	delUserFaved := "delete from user_faved_stickers where user_id = ?;"
	_, err = tx.Exec(delUserFaved, userId)
	raiseTx(err, tx)

	delImportContact := "delete from user_import_contacts where user_id = ?;"
	_, err = tx.Exec(delImportContact, userId)
	raiseTx(err, tx)

	delUserLogs := "delete from user_logs where user_id = ?;"
	_, err = tx.Exec(delUserLogs, userId)
	raiseTx(err, tx)

	delUserNotify := "delete from user_notify_settings where user_id = ?;"
	_, err = tx.Exec(delUserNotify, userId)
	raiseTx(err, tx)

	delUserPassword := "delete from user_passwords where user_id = ?;"
	_, err = tx.Exec(delUserPassword, userId)
	raiseTx(err, tx)

	delUserPhotoHistory := "delete from user_photo_history where user_id = ?;"
	_, err = tx.Exec(delUserPhotoHistory, userId)
	raiseTx(err, tx)

	delUserPresence := "delete from user_presences where user_id = ?;"
	_, err = tx.Exec(delUserPresence, userId)
	raiseTx(err, tx)

	delUserPrivacy := "delete from user_privacys where user_id = ?;"
	_, err = tx.Exec(delUserPrivacy, userId)
	raiseTx(err, tx)

	delUserReadSticker := "delete from user_read_sticker_sets where user_id = ?;"
	_, err = tx.Exec(delUserReadSticker, userId)
	raiseTx(err, tx)

	delUserRecentSticker := "delete from user_recent_stickers where user_id = ?;"
	_, err = tx.Exec(delUserRecentSticker, userId)
	raiseTx(err, tx)

	delUserStickerSet := "delete from user_sticker_sets where user_id = ?;"
	_, err = tx.Exec(delUserStickerSet, userId)
	raiseTx(err, tx)

	delAuthUser := "delete from auth_users where user_id = ?;"
	_, err = tx.Exec(delAuthUser, userId)
	raiseTx(err, tx)

	if Dbindex == 1 {
		delAuthSeq := "delete from auth_seq_updates where user_id = ?;"
		_, err = tx.Exec(delAuthSeq, userId)
		raiseTx(err, tx)

		delUserMsgbox := "delete from user_msgbox where user_id = ?;"
		_, err = tx.Exec(delUserMsgbox, userId)
		raiseTx(err, tx)

		delUserPts := "delete from user_pts_updates where user_id = ?;"
		_, err = tx.Exec(delUserPts, userId)
		raiseTx(err, tx)

		delUserQts := "delete from user_qts_updates where user_id = ?;"
		_, err = tx.Exec(delUserQts, userId)
		raiseTx(err, tx)
	} else {
		if userId%2 == 1 {
			delAuthSeq := "delete from auth_seq_updates1 where user_id = ?;"
			_, err = tx.Exec(delAuthSeq, userId)
			raiseTx(err, tx)

			delUserMsgbox := "delete from user_msgbox1 where user_id = ?;"
			_, err = tx.Exec(delUserMsgbox, userId)
			raiseTx(err, tx)

			delUserPts := "delete from user_pts_updates1 where user_id = ?;"
			_, err = tx.Exec(delUserPts, userId)
			raiseTx(err, tx)

			delUserQts := "delete from user_qts_updates1 where user_id = ?;"
			_, err = tx.Exec(delUserQts, userId)
			raiseTx(err, tx)
		} else {
			delAuthSeq := "delete from auth_seq_updates2 where user_id = ?;"
			_, err = tx.Exec(delAuthSeq, userId)
			raiseTx(err, tx)

			delUserMsgbox := "delete from user_msgbox2 where user_id = ?;"
			_, err = tx.Exec(delUserMsgbox, userId)
			raiseTx(err, tx)

			delUserPts := "delete from user_pts_updates2 where user_id = ?;"
			_, err = tx.Exec(delUserPts, userId)
			raiseTx(err, tx)

			delUserQts := "delete from user_qts_updates2 where user_id = ?;"
			_, err = tx.Exec(delUserQts, userId)
			raiseTx(err, tx)
		}
	}

	err = tx.Commit()
	raise(err)
	return true
}

func (dao *CommonDAO) GetBotToken(userId int32) string {
	qry := "select token from bot where creator_id = ? limit 1"
	row := dao.db.QueryRow(qry, userId)
	var token string
	err := row.Scan(&token)
	if err == sql.ErrNoRows {
		return ""
	}
	raise(err)
	return token
}

func (dao *CommonDAO) GetPageName(ids []int32) []string {
	res := make([]string, 0)
	if len(ids) == 0 {
		return res
	}
	qryType := "("

	for i, id := range ids {
		qryType += strconv.Itoa(int(id))
		if i == len(ids) - 1 {
			qryType += ")"
		} else {
			qryType += ","
		}
	}

	sqlStr := fmt.Sprintf("select name1, name2 from manage_permission_page where permissions_id in %s;", qryType)
	rows, err := dao.db.Queryx(sqlStr)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	m := make(map[string]bool, 0)
	for rows.Next() {
		var name1, name2 string
		err = rows.Scan(&name1, &name2)
		raise(err)
		m[name1] = true
		m[name2] = true
	}
	for k, _ := range m {
		res = append(res, k)
	}
	return res
}

func (dao *CommonDAO) GetConfig(key string) string {
	sqlStr := "select value from manage_config where `config_key` = ? limit 1;"
	row := dao.db.QueryRowx(sqlStr, key)
	var value string
	err := row.Scan(&value)
	if err == sql.ErrNoRows {
		return ""
	}
	raise(err)
	return value
}

func (dao *CommonDAO) UpdateConfig(key, value string) {
	sqlStr := "update manage_config set value = ? where `config_key` = ?;"
	_, err := dao.db.Exec(sqlStr, value, key)
	raise(err)
}

func (dao *CommonDAO) AddConfig(key, value string) {
	sqlStr := "insert into manage_config (config_key, value) values (?, ?);"
	_, err := dao.db.Exec(sqlStr, key, value)
	raise(err)
}

func (dao *CommonDAO) GetUserContact(id int32, limit, offset int32) ([]int32, int32) {
	res := make([]int32, 0)
	qry := "select contact_user_id from user_contacts where owner_user_id = ? and is_deleted = 0 order by date2 desc limit ? offset ?;"
	rows, err := dao.db.Queryx(qry, id, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		raise(err)
		res = append(res, id)
	}

	qryCount := "select count(*) from user_contacts where owner_user_id = ? and is_deleted = 0;"
	row := dao.db.QueryRow(qryCount, id)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}

func (dao *CommonDAO) GetUserGender(ids []int32) map[int32]int32 {
	res := make(map[int32]int32)
	if len(ids) == 0 {
		return res
	}
	sqlStr := "select user_id, gender from user_info where user_id in (?);"
	query, args, err := sqlx.In(sqlStr, ids)
	raise(err)
	rows, err := dao.db.Queryx(query, args...)
	raise(err)
	for rows.Next() {
		var user_id, gender int32
		err = rows.Scan(&user_id, &gender)
		raise(err)
		res[user_id] = gender
	}
	return res
}

func (dao *CommonDAO) GetDayActiveUser(start, end int64) int32 {
	defer func() {
		if err := recover(); err != nil {
			logger.LogSugar.Errorf("GetDayActiveUser err:%v", err)
			return
		}
	}()
	sqlStr :=`
	select
	  count(from_id) count
	from
	(
	  SELECT
	  from_id
	  FROM raw_message
	  WHERE add_time between ? and ?
	  group by from_id
	) d;
	`
	row := dao.db.QueryRowx(sqlStr, start, end)
	var count int32
	err := row.Scan(&count)
	if err == sql.ErrNoRows {
		return 0
	}
	raise(err)
	return count
}

func (dao *CommonDAO) GetDayActiveChat(start, end int64) int32 {
	defer func() {
		if err := recover(); err != nil {
			logger.LogSugar.Errorf("GetDayActiveChat err:%v", err)
			return
		}
	}()
	sqlStr :=`
	select
	  count(peer_id) count
	from
	(
	  SELECT
	  peer_id
	  FROM raw_message
	  WHERE add_time between ? and ? and peer_type in (3, 4)
	  group by peer_id
	) d;
	`
	row := dao.db.QueryRowx(sqlStr, start, end)
	var count int32
	err := row.Scan(&count)
	if err == sql.ErrNoRows {
		return 0
	}
	raise(err)
	return count
}

func (dao *CommonDAO) AddActieCache(uCount, cCount, uCount5, cCount5, msgCount, phoneCount int32, date string) {
	sqlStr := "insert into manage_active_data (date, user_count, chat_count, user_count5, chat_count5, msg_count, call_count, add_time) values (?, ?, ?, ?, ?, ?, ?, ?);"
	_, err := dao.db.Exec(sqlStr, date, uCount, cCount, uCount5, cCount5, msgCount, phoneCount, time.Now().Unix())
	raise(err)
}

func (dao *CommonDAO) ActieCacheExit(date string) bool {
	sqlStr := "select date from manage_active_data where date = ? limit 1"
	row := dao.db.QueryRowx(sqlStr, date)
	var value string
	err := row.Scan(&value)
	if err == sql.ErrNoRows {
		return false
	}
	raise(err)
	return value == date
}

func (dao *CommonDAO) GetActiveData5(date string) (int32, int32) {
	qry := `select user_count5, chat_count5 count from manage_active_data where date <= ? order by date desc limit 1;`
	row := dao.db.QueryRow(qry, date)
	var uCount, cCount int32
	err := row.Scan(&uCount, &cCount)
	if err == sql.ErrNoRows {
		return uCount, cCount
	}
	raise(err)
	return uCount, cCount
}

func (dao *CommonDAO) GetMsgPhoneCount() (int32, int32) {
	qry := `select sum(msg_count) mCount, sum(call_count) cCount from manage_active_data;`
	row := dao.db.QueryRow(qry)
	var mCount, cCount int32
	err := row.Scan(&mCount, &cCount)
	raise(err)
	return mCount, cCount
}

func (dao *CommonDAO) Get30DaysActiveData(start, end string) ([]map[string]interface{}, []map[string]interface{}) {
	userNums := make([]map[string]interface{}, 0)
	chatNums := make([]map[string]interface{}, 0)
	qry := "select date, user_count, chat_count from manage_active_data where date between ? and ?;"
	rows, err := dao.db.Queryx(qry, start, end)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return userNums, chatNums
	}
	raise(err)
	for rows.Next() {
		var date string
		var uNum, cNum int32
		err = rows.Scan(&date, &uNum, &cNum)
		raise(err)
		m := map[string]interface{}{
			"date": date,
			"num":  uNum,
		}
		userNums = append(userNums, m)

		m2 := map[string]interface{}{
			"date": date,
			"num":  cNum,
		}
		chatNums = append(chatNums, m2)
	}

	timeStr := time.Now().Add(24 * time.Hour).Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	e := t.Unix() - 8 *3600
	s := e - 3600*24

	u := dao.GetDayActiveUser(s, e)
	c := dao.GetDayActiveChat(s, e)
	todayUMap := map[string]interface{} {
		"date": t.Add(-16 * time.Hour).Format("2006-01-02"),
		"num":  u,
	}
	todayCMap := map[string]interface{} {
		"date": t.Add(-16 * time.Hour).Format("2006-01-02"),
		"num":  c,
	}
	userNums = append(userNums, todayUMap)
	chatNums = append(chatNums, todayCMap)
	return userNums,chatNums
}

//func (dao *CommonDAO) GetDayActiveChat() []map[string]interface{} {
//	dataNums := make([]map[string]interface{}, 0)
//	qry := `
//		select
//		  date_time, count(from_id) member_num
//		from
//		(
//		  SELECT
//		  from_id,
//		  from_unixtime(add_time, '%Y-%m-%d') date_time
//		  FROM raw_message
//		  WHERE add_time >= ? and peer_type in (3, 4)
//		  group by from_id, date_time
//		) d
//		group by date_time order by date_time;
//		`
//	rows, err := dao.db.Queryx(qry, start)
//	defer rows.Close()
//	if err == sql.ErrNoRows {
//		return dataNums
//	}
//	raise(err)
//	for rows.Next() {
//		var date string
//		var num int32
//		err = rows.Scan(&date, &num)
//		raise(err)
//		m := map[string]interface{}{
//			"date": date,
//			"num":  num,
//		}
//		dataNums = append(dataNums, m)
//	}
//	return dataNums
//}


func (dao *CommonDAO) GetDeviceUsers(device, limit, offset int32) ([]int32, int32) {
	// device 0:android 1:ios 2:mac 3:windows
	/*
	ios      iPhone 8 Plus  device_model like 'iPhone%'
	android  other     system_version like 'SDK%'
	windows  PC 64bit    device_model like 'PC%'
	mac      MacBook Pro  system_version like 'macOS%'
	*/
	var uIds []int32
	var whereSql string
	if device == 0 {
		whereSql = "where system_version like 'SDK%'"
	} else if device == 1 {
		whereSql = "where device_model like 'iPhone%'"
	} else if device == 2 {
		whereSql = "where system_version like 'macOS%'"
	} else if device == 3 {
		whereSql = "where device_model like 'PC%'"
	}

	sqlStr :=fmt.Sprintf(`
	select user_id
	from (
		select user_id, device_model from user_logs1 %s
		union all 
		select user_id, device_model from user_logs2 %s
	) d
	group  by user_id order by user_id desc limit ? offset ?
	`, whereSql, whereSql)
	rows, err := dao.db.Queryx(sqlStr, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return uIds, 0
	}
	raise(err)
	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		raise(err)
		uIds = append(uIds, id)
	}

	sqlConut :=fmt.Sprintf(`
	select count(*) count
	from (
		select user_id from user_logs1 %s group by user_id
		union all 
		select user_id from user_logs2 %s group by user_id
	) d
	`, whereSql, whereSql)
	row := dao.db.QueryRowx(sqlConut)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return uIds, count
}

func (dao *CommonDAO) DelUserPassword(userId int32) {
	sqlStr := "delete from user_passwords where user_id = ?;"
	_, err := dao.db.Exec(sqlStr, userId)
	raise(err)
}