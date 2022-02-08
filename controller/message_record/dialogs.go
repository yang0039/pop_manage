package message_record

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	//"pop-api/baselib/util"
	"pop-api/dal/dao"

	"pop-api/baselib/redis_client"
	"pop-api/controller/_dummy/memberapi"
	"pop-api/dto"
	"pop-api/middleware"
	//"strconv"
)

func (service *RecordController) GetDialogs(c *gin.Context) {
	params := &dto.MessageRecord{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}

	if params.FromId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}

	if params.PeerId == 0 {  // 查询全部
		if params.Limit == 0 {
			params.Limit = 20
		}
		dialogs, count := memberapi.GetUserDialog(params.FromId, params.Limit, params.Offset)
		res := make([]map[string]interface{}, 0, len(dialogs))
		for _, d := range dialogs {
			m := make(map[string]interface{})
			//key := fmt.Sprintf("%d:%d:%d", params.FromId, d["peer_type"], d["peer_id"])
			//top, _ := redis_client.RedisCache.HGet(key, "top").Result()
			//topMsgId, _ := strconv.Atoi(top)
			//m["num"] = topMsgId
			m["last_time"] = d["score"]
			m["peer"] = GetPeer(d["peer_type"], d["peer_id"])
			res = append(res, m)
		}

		data := map[string]interface{}{
			"dialogs": res,
			"count":   count,
		}
		middleware.ResponseSuccess(c, data)
	} else {
		m := make(map[string]interface{})
		key := fmt.Sprintf("z:%d", params.FromId)
		kP := 2
		if params.PeerId > 100000 {   // 用户
			peer := GetPeer(2, params.PeerId)
			m["peer"] = peer
		} else {    // 群
			peer := GetPeer(3, params.PeerId)
			chatDao := dao.GetChatDAO()
			c := chatDao.GetChat(params.PeerId)
			if c.Type == 1 || c.Type == 4 {    // 普通群或群发
				kP = 3
			} else {
				kP = 4
			}
			m["peer"] = peer
		}
		lTime,_ := redis_client.RedisCache.ZScore(key, fmt.Sprintf("%d:%d", kP, params.PeerId)).Result()
		m["last_time"] = int64(lTime)

		data := map[string]interface{}{
			"dialogs": []interface{}{m},
			"count":   1,
		}
		middleware.ResponseSuccess(c, data)
	}

}
