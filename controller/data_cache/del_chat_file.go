package data_cache

import (
	"pop-api/baselib/logger"
	"pop-api/baselib/redis_client"
	"pop-api/baselib/util"
	"pop-api/controller/message_record"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

// 定时从缓存中取出需要删除文件资源的群，并删除文件
func DelChatFileData() {
	defer func() {
		if err := recover(); err != nil {
			logger.LogSugar.Errorf("recover error DelChatFileData panic: %v\n%s", err, string(debug.Stack()))
		}
	}()

	loopTimer := time.NewTimer(time.Minute * 1)
	defer loopTimer.Stop()

	for {
		loopTimer.Reset(time.Minute * 1)

		select {
		case <-loopTimer.C:
			delChatFile()
		}
	}
}

func delChatFile() {
	chatIds,_ := redis_client.RedisCache.SMembers("del_chat_file").Result()
	for _, typeCid := range chatIds {
		pc := strings.Split(typeCid, ":")
		if len(pc) != 2 {
			logger.LogSugar.Errorf("delChatFile unknow typeCid:%s", typeCid)
		}
		pType, _ := strconv.Atoi(pc[0])
		chatId, _ := strconv.Atoi(pc[1])
		if chatId == 0 {
			continue
		}
		logger.LogSugar.Infof("delChatFile typeCid:%s", typeCid)
		// 删除文件
		var peerType int32
		if pType == 3 {
			peerType = util.PeerChat
		} else if pType == 4 {
			peerType = util.PeerChannel
		}
		message_record.DelFile(peerType, int32(chatId), 0, 0, false)
		// 成功后删除缓存中该群
		redis_client.RedisCache.SRem("del_chat_file", typeCid)
	}
}

