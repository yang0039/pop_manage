package data_tool

import (
	"fmt"
	"pop-api/baselib/logger"
	"pop-api/baselib/redis_client"
	"pop-api/dal/dao"
	"strconv"
	"time"
)

func RepairMsgPts() {
	comDao := dao.GetCommonDAO()
	var offset int32
	limit := int32(500)

	c := 0
	for {
		uIds := comDao.GetAllUser(limit, offset)
		for _, uid := range uIds {
			msgId,_ := redis_client.RedisCache.Get(fmt.Sprintf("msg_box:%d", uid)).Int()
			pts,_ := redis_client.RedisCache.Get(fmt.Sprintf("pts:%d", uid)).Int()
			maxMsgId, maxPts := comDao.GetMaxMsgId(uid)
			var info string
			if msgId != 0 && (maxMsgId - int32(msgId)) > 1 {
				redis_client.RedisCache.Set(fmt.Sprintf("msg_box:%d", uid), maxMsgId, 0)
				info += fmt.Sprintf("maxMsgId:%d, msgId:%d", maxMsgId, msgId)
			}
			if pts != 0 && (maxPts - int32(pts)) > 1 {
				redis_client.RedisCache.Set(fmt.Sprintf("pts:%d", uid), maxPts, 0)
				info += fmt.Sprintf(" maxPts:%d, pts:%d", maxPts, pts)
			}
			if info != "" {
				info2 := fmt.Sprintf("=== uid:%d, %s", uid, info)
				logger.LogSugar.Infof(info2)
				c ++
			}
		}
		//if c > 5 {
		//	break
		//}
		if len(uIds) < int(limit) {
			break
		}
		offset += limit
		logger.LogSugar.Infof("RepairMsgPts offset:%d, count:%d", offset, c)
		time.Sleep(1 * time.Second)

	}
	logger.LogSugar.Infof("===============success=================")
}

func RepairChannelMsg() {
	comDao := dao.GetCommonDAO()
	var offset int32
	limit := int32(500)

	c := 0
	for {
		cIds := comDao.GetAllChannel(limit, offset)
		for _, cid := range cIds {
			msgId,_ := redis_client.RedisCache.Get(fmt.Sprintf("channel_msg_id:%d", cid)).Int()
			//pts,_ := redis_client.RedisCache.Get(fmt.Sprintf("pts:%d", uid)).Int()
			maxMsgId := comDao.GetChannelMaxMsgId(cid)
			var info string
			if msgId != 0 && (maxMsgId - int32(msgId)) > 1 {
				redis_client.RedisCache.Set(fmt.Sprintf("channel_msg_id:%d", cid), maxMsgId, 0)
				info += fmt.Sprintf("maxMsgId:%d, msgId:%d", maxMsgId, msgId)
			}
			//if pts != 0 && (maxPts - int32(pts)) > 1 {
			//	redis_client.RedisCache.Set(fmt.Sprintf("pts:%d", uid), maxPts, 0)
			//	info += fmt.Sprintf(" maxPts:%d, pts:%d", maxPts, pts)
			//}
			if info != "" {
				info2 := fmt.Sprintf("=== cid:%d, %s", cid, info)
				logger.LogSugar.Infof(info2)
				c ++
			}
		}
		//if c > 5 {
		//	break
		//}
		if len(cIds) < int(limit) {
			break
		}
		offset += limit
		logger.LogSugar.Infof("RepairChannelMsg offset:%d, count:%d", offset, c)
		time.Sleep(1 * time.Second)

	}
	logger.LogSugar.Infof("===============success=================")
}

func RepairName() {
	comDao := dao.GetCommonDAO()
	// 修复这段时间的数据
	start := int64(1658764800)
	end := int64(1658851200)
	//           1658764805
	users := comDao.GetTimeUser(start, end)
	var c int32
	for _, user := range users {
		//logger.LogSugar.Infof("user:%v", user)
		uidStr, _ := user["id"]
		uid,_ := strconv.Atoi(uidStr)
		if uid == 0 {
			logger.LogSugar.Infof("unknow user:%v", user)
			continue
		}
		f := user["first_name"]
		l := user["last_name"]

		fName,_ := redis_client.RedisCache.HGet(fmt.Sprintf("auth:%d:first_name", uid), "0").Result()
		//lName,_ := redis_client.RedisCache.HGet(fmt.Sprintf("auth:%d:last_name", uid), "0").Result()

		var info string
		if fName == "" && f != "" {
			info += fmt.Sprintf("fname:%s  ", f)
			redis_client.RedisCache.HSet(fmt.Sprintf("auth:%d:first_name", uid), "0", f)
		}

		redis_client.RedisCache.HSet(fmt.Sprintf("auth:%d:last_name", uid), "0", l)
		info += fmt.Sprintf("lname:%s  ", l)
		//if lName == "" && l != "" {
		//	info += fmt.Sprintf("lname:%s  ", l)
		//	redis_client.RedisCache.HSet(fmt.Sprintf("auth:%d:last_name", uid), "0", l)
		//}

		if info != "" {
			info2 := fmt.Sprintf("=== uid:%d, %s", uid, info)
			logger.LogSugar.Infof(info2)
			c ++
		}

	}
	logger.LogSugar.Infof("===============success  count:%d=================", c)
}
