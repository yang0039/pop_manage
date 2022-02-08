package impl

import (
	"pop-api/controller/_dummy/chatapi"
	"pop-api/controller/_dummy/memberapi"
	"pop-api/controller/_dummy/recordapi"
	"pop-api/controller/chat_manage"
	"pop-api/controller/member_manage"
	"pop-api/controller/message_record"
)

func init() {
	chatapi.ChatInfo = chat_manage.ChatInfo
	chatapi.User_assemble = chat_manage.User_assemble
	chatapi.Channel_assemble = chat_manage.Channel_assemble

	recordapi.UserMsgToMessage = message_record.UserMsgToMessage
	recordapi.ChannelMsgToMessage = message_record.ChannelMsgToMessage
	recordapi.GetPeer = message_record.GetPeer
	recordapi.GetUser = message_record.GetUser

	memberapi.GetUserDialog = member_manage.GetUserDialog
}
