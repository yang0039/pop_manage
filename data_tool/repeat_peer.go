package data_tool

import(
	"fmt"
	"pop-api/baselib/logger"
	"pop-api/baselib/redis_client"
	"pop-api/dal/dao"
)

/*
有些数据超级群的id，会同时作为超级群和普通群存在，需要删除这些群作为普通群在会话列表里面的数据
*/





func DelRepeatData() {
	comDao := dao.GetCommonDAO()
	repeatData := comDao.GetRepeatPeerData()
	chatDao := dao.GetChatDAO()

	successCount := 0
	for _, d := range repeatData {

		cId := d["peer_id"]
		uId := d["user_id"]
		if cId == 0 {
			continue
		}
		//pType1 := d["peer_type1"]  // 普通群
		//pType2 := d["peer_type2"]  // 超级群
		id1 := d["id1"]
		id2 := d["id2"]

		var delId int32
		var delPeer string
		c := chatDao.GetChat(cId)
		if c.Type == 1 || c.Type == 4 {         // 该群为普通群或者群发,需要将超级群数据删除
			delId = id2
			delPeer = fmt.Sprintf("4:%d", cId)
		} else if c.Type == 2 || c.Type == 3 {  // 该群为超级群，需要将普通群数据删除
			delId = id1
			delPeer = fmt.Sprintf("3:%d", cId)
		} else {
			logger.LogSugar.Infof("continue del peer uid:%d, type:%d, cId:%d", uId, c.Type, cId)
			continue
		}

		comDao.DelRepeatPeerData(delId)
		redis_client.RedisCache.ZRem(fmt.Sprintf("z:%d", uId), delPeer)
		redis_client.RedisCache.ZRem(fmt.Sprintf("zp:%d", uId), delPeer)
		logger.LogSugar.Infof("del peer uid:%d, cId:%d, delId:%d, delPeer:%s", uId, cId, delId, delPeer)
		successCount ++
	}
	logger.LogSugar.Infof("DelRepeatData success len:%d", successCount)

}

