package dataobject

/* user_msgbox表的字段 */
type UserMsgRow struct {
	Id        int32 `db:"id"`
	UserId    int32 `db:"user_id"`
	MsgId     int32 `db:"msg_id"`
	Pts       int32 `db:"pts"`
	FromMsgId int32 `db:"from_msg_id"`

	RawId        int64 `db:"raw_id"`
	Type         int   `db:"type"`
	FromId       int32 `db:"from_id"`
	PeerType     int32 `db:"peer_type"`
	PeerId       int32 `db:"peer_id"`
	ReplyToMsgId int32 `db:"reply_to_msg_id"`
	Mentioned    bool  `db:"mentioned"`
	MediaUnread  bool  `db:"media_unread"`
	AddTime      int32 `db:"add_time"`
}

