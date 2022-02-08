package mysql_dao

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
	"time"
)

type PeerStatusDAO struct {
	db *sqlx.DB
}

func NewPeerStatusDAO(db *sqlx.DB) *PeerStatusDAO {
	return &PeerStatusDAO{db}
}

func (dao *PeerStatusDAO) AddStatus(peerType, peerId, status int32, util int64, note, operator string) {
	dao.DelStatus(peerId)
	sql := `
	insert into manage_peer_status (peer_type, peer_id, status, util, note, operator, add_time)
	values (?, ?, ?, ?, ?, ?, ?);
	`
	now := time.Now().Unix()
	_, err := dao.db.Exec(sql, peerType, peerId, status, util, note, operator, now)
	raise(err)
}

func (dao *PeerStatusDAO) DelStatus(peerId int32) {
	sql := "update manage_peer_status set is_effect = 0, remove_time = ? where peer_id = ? and is_effect = 1;"
	_, err := dao.db.Exec(sql, time.Now().Unix(), peerId)
	raise(err)
}

func (dao *PeerStatusDAO) GetChatStatus(chatIds []int32) map[int32]*dataobject.PeerStatus{
	res := make(map[int32]*dataobject.PeerStatus)
	sqlStr := "select peer_id, status, util, note from manage_peer_status where peer_id in (?) and is_effect = 1;"
	query, args, err := sqlx.In(sqlStr, chatIds)
	raise(err)
	rows, err := dao.db.Queryx(query, args...)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var status dataobject.PeerStatus
		err = rows.StructScan(&status)
		raise(err)
		res[status.PeerId] = &status
	}
	return res
}

func (dao *PeerStatusDAO) QryStatusRecord(peerId, limit, offset int32) ([]*dataobject.PeerStatusRecord, int32) {
	res := make([]*dataobject.PeerStatusRecord, 0)
	sqlStr := `
	select * from (
		SELECT
			id,
			peer_type,
			peer_id,
			status,
			util,
			note,
			'banned' opera,
			add_time opera_time
		FROM manage_peer_status
		WHERE peer_id = ?
		UNION ALL
		SELECT
			id,
			peer_type,
			peer_id,
			status,
			util,
			note,
			'unbanned' opera,
			remove_time opera_time
		FROM manage_peer_status
		WHERE is_effect = 0 and peer_id = ?
	) data
	order by opera_time desc limit ? offset ?;
	`
	rows, err := dao.db.Queryx(sqlStr, peerId, peerId, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var status dataobject.PeerStatusRecord
		err = rows.StructScan(&status)
		raise(err)
		res = append(res, &status)
	}

	sqlCount := `
	select count(*) count from (
		SELECT
			peer_type,
			peer_id,
			status,
			util,
			note,
			'banned' opera,
			add_time opera_time
		FROM manage_peer_status
		WHERE peer_id = ?
		UNION ALL
		SELECT
			peer_type,
			peer_id,
			status,
			util,
			note,
			'unbanned' opera,
			remove_time opera_time
		FROM manage_peer_status
		WHERE is_effect = 0 and peer_id = ?
	) data
	`

	row := dao.db.QueryRow(sqlCount, peerId, peerId)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}