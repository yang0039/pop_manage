package chat_manage

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"pop-api/baselib/redis_client"
	"pop-api/dal/dao"
	"pop-api/dto"
	"pop-api/middleware"
	"time"
)

func (service *ChatController) AddChatStatus(c *gin.Context) {
	//params := &dto.ChatStatus{}
	//if err := c.ShouldBind(params); err != nil {
	//	middleware.ResponseError(c, 500, "系统错误", err)
	//	return
	//}

	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	params, _ := bindData.(*dto.ChatStatus)
	if params.ChatId == 0 || params.Status > 4 || params.Status < 1 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	userName, _ := c.Get("user_name")
	name, _ := userName.(string)

	statusDao := dao.GetPeerStatusDAO()
	key := "manage:chat:status"
	if params.Status == 1 { // 正常
		statusDao.DelStatus(params.ChatId)
		exit, _ := redis_client.RedisCache.SIsMember(key, params.ChatId).Result()
		if exit {
			redis_client.RedisCache.SRem(key, params.ChatId)
		}
	} else { // 1:正常 2:警告 3:短期禁言 4:长期禁言
		var utilTime int64
		if params.Status == 3 {
			now := time.Now().Unix()
			utilTime = now + 3600*24*int64(params.Days)
		} else {
			utilTime = -1
		}
		if params.Status == 1 {
			statusDao.DelStatus(params.ChatId)
		} else {
			statusDao.AddStatus(2, params.ChatId, params.Status, utilTime, params.Note, name)
			if params.Status == 3 || params.Status == 4 {
				redis_client.RedisCache.SAdd(key, params.ChatId)
			}
		}
	}
	data := map[string]interface{}{
		"chat_id": params.ChatId,
	}
	middleware.ResponseSuccess(c, data)
}

func (service *ChatController) QryChatStatusRecord(c *gin.Context) {
	//params := &dto.ChatStatus{}
	//if err := c.ShouldBind(params); err != nil {
	//	middleware.ResponseError(c, 500, "系统错误", err)
	//	return
	//}

	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	params, _ := bindData.(*dto.ChatStatus)
	if params.ChatId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	if params.Limit == 0 {
		params.Limit = 20
	}

	statusDao := dao.GetPeerStatusDAO()
	res, count := statusDao.QryStatusRecord(params.ChatId, params.Limit, params.Offset)
	data := map[string]interface{}{
		"record_data": res,
		"count":       count,
	}
	middleware.ResponseSuccess(c, data)

}