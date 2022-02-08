package dataobject

/* 原始消息表字段 */
type RawMessageRow struct {
	Id             int32  `db:"id"`
	RawId          int64  `db:"raw_id"`
	MessageBlob    []byte `db:"message_blob"`
	MessageEncrypt string `db:"message_encrypt"`
	FromId         int32  `db:"from_id"`
	PeerType       int32  `db:"peer_type"`
	PeerId         int32  `db:"peer_id"`
	Views          int32  `db:"views"`
	EditDate       int32  `db:"edit_date"`
	AddTime        int32  `db:"add_time"`
}
