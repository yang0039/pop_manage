package util

import (
	"fmt"
	"pop-api/baselib/logger"
	"strconv"
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

func FileType(ext string) string {
	/*
		file      文件
		photo     图片
		audio     音频
		video     视频
		other     其他
	*/
	//-- 文件
	//-- 图片  .jpg  .png .jepg
	//-- 音频  .ogg
	//-- 视频  .mp4  avi
	//-- 其他  ''

	switch ext {
	case ".jpg",".png",".jepg":
		return "photo"
	case ".ogg":
		return "audio"
	case ".mp4", ".avi":
		return "video"
	case "","other":
		return "other"
	default:
		return "file"
	}
}

func Folat4(d float64) float64 {
	f,_ := strconv.ParseFloat(fmt.Sprintf("%.4f", d), 64)
	return f
}

func Folat2(d float64) float64 {
	f,_ := strconv.ParseFloat(fmt.Sprintf("%.2f", d), 64)
	return f
}







