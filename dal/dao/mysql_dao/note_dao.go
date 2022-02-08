package mysql_dao

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type NoteDAO struct {
	db *sqlx.DB
}

func NewNoteDAO(db *sqlx.DB) *NoteDAO {
	return &NoteDAO{db}
}


func (dao *NoteDAO) AddNote(labelIds []string, peerType, peerId int32, note string) {
	sql := `
	insert into manage_note (label_id, peer_type, peer_id, note, updated_at)
	values (?, ?, ?, ?, ?)
	on DUPLICATE KEY UPDATE note = ?, label_id = ?;
	`
	now := time.Now().Unix()
	for _, labelId := range labelIds {
		_, err := dao.db.Exec(sql, labelId, peerType, peerId, note, now, note, labelId)
		raise(err)
	}
}

func (dao *NoteDAO) GetNote(peerIds []int32) map[int32]map[string]string {
	noteMap := make(map[int32]map[string]string)
	sql2 := "("
	for i, id := range peerIds {
		sql2 = sql2 + fmt.Sprintf("%d", id)
		if i == len(peerIds) - 1 {
			sql2 += ")"
		} else {
			sql2+= ","
		}
	}
	qry := fmt.Sprintf(`
	select
		peer_id,
		ifnull(group_concat(concat(label_id,'_@_', label_name)),'') labels, max(note) note
	from manage_note n
	left join manage_label l on n.label_id = l.id
	where peer_id in %s group by peer_id;
	`, sql2)

	rows, err := dao.db.Queryx(qry)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return noteMap
	}
	raise(err)
	for rows.Next() {
		var peer int32
		var note, labels string
		m := make(map[string]string)
		err = rows.Scan(&peer, &labels, &note)
		raise(err)
		m["labels"] = labels
		m["note"] = note
		noteMap[peer] = m
	}
	return noteMap
}

func (dao *NoteDAO) GetChatByNote(note string, limit, offset int32) ([]int32, int32) {
	chatIds := make([]int32, 0)
	if note == "" {
		return chatIds, 0
	}
	qry := `
	select peer_id
	from manage_note n
	left join chat c on n.peer_id = c.id
	where peer_type in (3, 4) and note = ? and deactivated = 0 and c.type in (1, 2)
	limit ? offset ?;
	`
	rows, err := dao.db.Queryx(qry, note, limit, offset)
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
	from manage_note n
	left join chat c on n.peer_id = c.id
	where peer_type in (3, 4) and note = ? and deactivated = 0 and c.type in (1, 2)
	`
	row := dao.db.QueryRow(qryCount, note)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return chatIds, count
}

func (dao *NoteDAO) GetUserByNote(note string, limit, offset int32) ([]int32, int32) {
	uIds := make([]int32, 0)
	if note == "" {
		return uIds, 0
	}
	qry := `
	select peer_id
	from manage_note n
	where peer_type = 2 and note = ? limit ? offset ?;
	`
	rows, err := dao.db.Queryx(qry, note, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return uIds, 0
	}
	raise(err)
	for rows.Next() {
		var chatId int32
		err = rows.Scan(&chatId)
		raise(err)
		uIds = append(uIds, chatId)
	}

	qryCount := "select count(*) from manage_note where peer_type = 2 and note = ?;"
	row := dao.db.QueryRow(qryCount, note)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return uIds, count
}

func (dao *NoteDAO) GetLabbelNoteCount() map[string]int32 {
	res := make( map[string]int32, 0)
	qry := `
	select
	concat(label_id, '_', peer_type) ptype, count(*) count
	from manage_note
	group by label_id, peer_type;
	`
	rows, err := dao.db.Queryx(qry)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var ptype string
		var count int32
		err = rows.Scan(&ptype, &count)
		raise(err)
		res[ptype] = count
	}
	return res
}


func (dao *NoteDAO) GetLabelChatIds(peerType, limit, offset int32, ids []string) ([]int32, int32) {
	res := make([]int32, 0)

	if len(ids) == 0 {
		return res, 0
	}
	qry := `
	select peer_id
	from manage_note
	where peer_type = ? and label_id in (?)
	group by peer_id limit ? offset ?;
	`
	query, args, err := sqlx.In(qry, peerType, ids, limit, offset)
	raise(err)
	rows, err := dao.db.Queryx(query, args...)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var peerId int32
		err = rows.Scan(&peerId)
		raise(err)
		res = append(res, peerId)
	}

	qryCount := `
	select count(*) c
	FROM (
		SELECT peer_id
		FROM manage_note
		WHERE peer_type = ? AND label_id in (?)
		GROUP BY peer_id
	) c;
	`
	query2, args, err := sqlx.In(qryCount, peerType, ids)
	raise(err)
	row := dao.db.QueryRow(query2, args...)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}