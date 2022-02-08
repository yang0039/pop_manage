package member_manage

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	goredis "github.com/go-redis/redis/v7"
	"pop-api/baselib/redis_client"
	"pop-api/baselib/util"
	"pop-api/dal/dao"
	"pop-api/dto"
	"pop-api/middleware"
	"strconv"
	"strings"
)

// 用户关系
func (service *UserController) UserRelation(c *gin.Context) {
	params := &dto.QryUser{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.UserId == 0 || params.PeerId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}

	commonDao := dao.GetCommonDAO()
	userDao := dao.GetUserDAO()
	partDao := dao.GetChatParticipantDAO()

	//contUserIds, count := commonDao.GetUserContact(params.UserId, params.Limit, params.Offset)

	contUserIds := []int32{params.PeerId}

	userInfos := userDao.GetUsers(contUserIds)
	genderMap := commonDao.GetUserGender(contUserIds) // 0:未知 1：男 2：女
	status := userStatus(contUserIds)
	userLasttime := commonDao.GetUserActiveTime(contUserIds)
	if len(userInfos) == 1 {
		u := userInfos[0]

		uMap := make(map[string]interface{})
		uMap["is_friend"] = isFriend(params.UserId, params.PeerId)
		uMap["user_id"] = u.Id
		uMap["user_name"] = u.Username
		uMap["first_name"] = u.FirstName
		uMap["last_name"] = u.LastName
		uMap["gender"] = genderMap[u.Id]
		uMap["phone"] = u.Phone
		uMap["common_chat_count"] = partDao.GetCommonChatsCount(params.UserId, u.Id)
		uMap["online"] = status[u.Id]
		uMap["last_active_time"] = userLasttime[u.Id]

		data := map[string]interface{}{
			"contact_user": uMap,
		}
		middleware.ResponseSuccess(c, data)

	} else {
		middleware.ResponseError(c, 400, "用户未发现", errors.New(fmt.Sprintf("peer_id:%v", params.PeerId)))
		return
	}
}

func isFriend(aId, bId int32) bool {
	exit1,_ := redis_client.RedisCache.SIsMember(fmt.Sprintf("auth:%d:contact_set", aId), bId).Result()
	exit2,_ := redis_client.RedisCache.SIsMember(fmt.Sprintf("auth:%d:contact_set", bId), aId).Result()
	if exit1 && exit2 {
		return true
	}
	return false
}

// 好友列表
func (service *UserController) UserContact(c *gin.Context) {
	params := &dto.QryUser{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.UserId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	if params.Limit == 0 {
		params.Limit = 20
	}
	commonDao := dao.GetCommonDAO()
	userDao := dao.GetUserDAO()
	partDao := dao.GetChatParticipantDAO()

	contUserIds, count := commonDao.GetUserContact(params.UserId, params.Limit, params.Offset)
	userInfos := userDao.GetUsers(contUserIds)
	genderMap := commonDao.GetUserGender(contUserIds) // 0:未知 1：男 2：女
	status := userStatus(contUserIds)
	userLasttime := commonDao.GetUserActiveTime(contUserIds)
	res := make([]map[string]interface{}, 0, len(contUserIds))
	for _, u := range userInfos {
		uMap := make(map[string]interface{})
		uMap["user_id"] = u.Id
		uMap["user_name"] = u.Username
		uMap["first_name"] = u.FirstName
		uMap["last_name"] = u.LastName
		uMap["gender"] = genderMap[u.Id]
		uMap["phone"] = u.Phone
		uMap["common_chat_count"] = partDao.GetCommonChatsCount(params.UserId, u.Id)
		uMap["online"] = status[u.Id]
		uMap["last_active_time"] = userLasttime[u.Id]
		res = append(res, uMap)
	}

	data := map[string]interface{}{
		"contact_user": res,
		"count":        count,
	}
	middleware.ResponseSuccess(c, data)
}

// 对话列表
func (service *UserController) UserDialogs(c *gin.Context) {
	params := &dto.QryUser{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.UserId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	if params.Limit == 0 {
		params.Limit = 20
	}

	dialogs, count := GetUserDialog(params.UserId, params.Limit, params.Offset)
	userDao := dao.GetUserDAO()
	chatDao := dao.GetChatDAO()
	dialogsRes := make([]map[string]interface{}, 0, len(dialogs))

	for _, d := range dialogs {
		m := make(map[string]interface{})
		m["peer_id"] = d["peer_id"]
		m["time"] = d["score"]
		if d["peer_type"] == 2 {
			users := userDao.GetUsers([]int32{d["peer_id"]})
			if len(users) == 1 {
				m["first_name"] = users[0].FirstName
				m["last_name"] = users[0].LastName
				m["title"] = ""
				m["peer_type"] = util.PeerUser
			}
		} else {
			chat := chatDao.GetChat(d["peer_id"])
			m["first_name"] = ""
			m["last_name"] = ""
			m["title"] = chat.Title
			if chat.Type == 1 || chat.Type == 2 {
				m["peer_type"] = util.PeerChat
			} else if chat.Type == 3 {
				m["peer_type"] = util.PeerChannel
			} else if chat.Type == 4 {
				m["peer_type"] = util.PeerMass
			}
		}
		dialogsRes = append(dialogsRes, m)
	}

	data := map[string]interface{}{
		"dialogs": dialogsRes,
		"count":   count,
	}
	middleware.ResponseSuccess(c, data)
}

func GetUserDialog(uid, limit, offset int32) ([]map[string]int32, int32) {
	res := make([]map[string]int32, 0)
	zKey := fmt.Sprintf("z:%d", uid)
	r := &goredis.ZRangeBy{
		Max:    "2147483647",
		Min:    "0",
		Offset: int64(limit * offset),
		Count:  int64(limit),
	}
	dialogs, err := redis_client.RedisCache.ZRevRangeByScoreWithScores(zKey, r).Result()
	if err != nil {
		return res, 0
	}
	for _, dialog := range dialogs {
		m := make(map[string]int32, 0)
		key, _ := dialog.Member.(string)
		dias := strings.Split(key, ":")
		if len(dias) == 2 {
			pType, _ := strconv.Atoi(dias[0])
			pId, _ := strconv.Atoi(dias[1])
			m["peer_type"] = int32(pType)
			m["peer_id"] = int32(pId)
			m["score"] = int32(dialog.Score)
			res = append(res, m)
		}
	}
	count, _ := redis_client.RedisCache.ZCount(zKey, "0", "2147483647").Result()
	return res, int32(count)
}
