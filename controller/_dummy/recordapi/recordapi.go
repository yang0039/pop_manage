package recordapi

import (
	"pop-api/dal/dataobject"
	"pop-api/dto"
)

var (
	UserMsgToMessage func(msgs []*dataobject.UserMsgRow) []*dto.Message
	ChannelMsgToMessage func(msgs []*dataobject.ChannelMsgRow) []*dto.Message
	GetPeer func(peerType, peerId int32) dto.Peer
	GetUser func(userId int32) dto.User
)

