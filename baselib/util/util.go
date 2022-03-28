package util

import (
	"pop-api/baselib/logger"
	"time"
)

const (
	PeerUser = 1       // 个人
	PeerChat = 2       // 群
	PeerChannel = 3    // 频道
	PeerMass = 4       // 群发
)

func DbToApiChatType(peerType int32) int32 {
	if peerType == 1 || peerType == 2 {
		return PeerChat
	} else if peerType == 3 {
		return PeerChannel
	} else if peerType == 4 {
		return PeerMass
	}
	return 0
}

func GetTodayUnix() int64 {
	y, m, d := time.Now().Date()
	today := time.Date(y, m, d, 0, 0, 0, 0, time.FixedZone("CST", 3600*8)).Unix()
	return today
}

func RaiseDBERR(err error) {
	if err == nil {
		return
	}
	logger.Logger.Error(err.Error())
	panic(err.Error())
}
