package dataobject

type ChannelMsgRow struct {
	Id           int32 `db:"id"`
	ChatId       int32 `db:"chat_id"`
	MsgId        int32 `db:"msg_id"`
	RawId        int64 `db:"raw_id"`
	Type         int   `db:"type"`
	FromId       int32 `db:"from_id"`
	ReplyToMsgId int32 `db:"reply_to_msg_id"`
	MentionId    int32 `db:"mention_id"`
	MediaUnread  bool  `db:"media_unread"`
	AddTime      int32 `db:"add_time"`
}
