package dataobject

type ReportDO struct {
	Id       int32  `db:"id"`
	UserId   int32  `db:"user_id"`
	PeerType int32  `db:"peer_type"`
	PeerId   int32  `db:"peer_id"`
	MsgIds   string `db:"msg_ids"`
	Reason   int32  `db:"reason"`
	Content  string `db:"content"`
	AddTime  int64  `db:"add_time"`
}
