package chatapi

import (
	"pop-api/dal/dataobject"
	"pop-api/mtproto"
)

var (
	ChatInfo         func(chatIds7 []int32) []map[string]interface{}
	User_assemble    func(msg *dataobject.UserMsgRow, raw *dataobject.RawMessageRow) *mtproto.Message
	Channel_assemble func(self_id int32, msg *dataobject.ChannelMsgRow, raw *dataobject.RawMessageRow) *mtproto.Message
)
