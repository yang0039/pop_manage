package member_manage

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ip2location/ip2location-go"
	"pop-api/baselib/redis_client"
	"pop-api/baselib/util"
	"pop-api/controller/_dummy/chatapi"
	"pop-api/dal/dao"
	"pop-api/dal/dataobject"
	"pop-api/dto"
	"pop-api/middleware"
	"pop-api/public"
	"strconv"
	"strings"
	"time"
)

// 更新群备注
func (service *UserController) AddUserNote(c *gin.Context) {
	//params := &dto.NoteObj{}
	//if err := c.ShouldBind(params); err != nil {
	//	middleware.ResponseError(c, 500, "系统错误", err)
	//	return
	//}

	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	params, _ := bindData.(*dto.NoteObj)
	if params.LabelId == "" || params.PeerId == 0 || params.Note == "" {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}

	labelIds := strings.Split(params.LabelId, ",")
	noteDao := dao.GetNoteDAO()
	noteDao.AddNote(labelIds, public.PEER_USER, params.PeerId, params.Note)

	// 记录操作
	accId, _ := c.Get("account_id")
	id, _ := accId.(int32)
	opera := fmt.Sprintf("设置标签:%v", dao.GetLabelDAO().GetNames(labelIds))
	if params.Note  != "" {
		opera += ", 备注:" + params.Note
	}
	uo := &dataobject.UserOpera{
		AccountId:    id,
		UserId:       params.PeerId,
		OperaType:    util.UpdateNote,
		OperaContent: opera,
		AddTime:      time.Now().Unix(),
	}
	dao.GetUserOperaDAO().AddOperaRecords(uo)


	middleware.ResponseSuccess(c, "")
}

// 查询字段：1.名称；2.POP ID；3.绑定电话号；4.注册日期；5.最后活跃日期；6.所在群组数；7.群组拥有者数；8.群组管理员数；9.当前在线状态；10.电子邮箱；11.设定语言；12.标注内容
// 查询条件：1.名称；2. POP ID；3.电话号；4.注册國码；5.电子邮箱 6.语言设定；7.标注关键字; 8.用户id; 9.封禁状态；10.标签；11.活跃日期；12.注册日期；13.设备类型 14.在线状态
func (service *UserController) GetUserInfo(c *gin.Context) {
	//params := &dto.QryType{}
	//if err := c.ShouldBind(params); err != nil {
	//	middleware.ResponseError(c, 500, "系统错误", err)
	//	return
	//}

	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	params, _ := bindData.(*dto.QryType)

	if params.Limit == 0 {
		params.Limit = 20
	}
	userDao := dao.GetUserDAO()
	commonDao := dao.GetCommonDAO()
	noteDao := dao.GetNoteDAO()
	bannedDao := dao.GetBannedDAO()

	var userIds []int32
	var count int32
	if params.Type == util.QryUserByName {
		userIds, count = userDao.GetUserIdsByName(params.Qry, params.Limit, params.Offset)
	} else if params.Type == util.QryUserByUserName {
		userIds, count = userDao.GetUserIdsByUserName(params.Qry, params.Limit, params.Offset)
	} else if params.Type == util.QryUserByPhone {
		userIds, count = userDao.GetUserIdsByPhone(params.Qry, params.Limit, params.Offset)
	} else if params.Type == util.QryUserByCountry {
		userIds, count = userDao.GetUserIdsByCountryCode(params.Qry, params.Limit, params.Offset)
	} else if params.Type == util.QryUserByEmail {
		userIds, count = commonDao.GetUserIdsByEmail(params.Qry, params.Limit, params.Offset)
	} else if params.Type == 6 {

	} else if params.Type == util.QryUserByNote {
		userIds, count = noteDao.GetUserByNote(params.Qry, params.Limit, params.Offset)
	} else if params.Type == util.QryUserById {
		uId, _ := strconv.Atoi(params.Qry)
		userIds = []int32{int32(uId)}
		count = 1
	} else if params.Type == util.QryUserByBanned {
		if params.Qry == "1" {
			userIds, count = bannedDao.GetUserByBanned(params.Limit, params.Offset)
		} else {
			userIds, count = userDao.GetUserIdsNoBanned(params.Limit, params.Offset)
		}
	} else if params.Type == util.QryUserByLabel { // 标签
		//labelId, _ := strconv.Atoi(params.Qry)
		labelIds := strings.Split(params.Qry, ",")
		userIds, count = noteDao.GetLabelUserIds(2, params.Limit, params.Offset, labelIds)
	} else if params.Type == util.QryUserByActiveDate { // 活跃日期
		nums := strings.Split(params.Qry, ",")
		if len(nums) != 2 {
			middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
			return
		}
		start, _ := strconv.Atoi(nums[0])
		end, _ := strconv.Atoi(nums[1])
		userIds, count = commonDao.GetUserByActive(int64(start), int64(end), params.Limit, params.Offset)
	} else if params.Type == util.QryUserByCreateDate { // 注册日期
		nums := strings.Split(params.Qry, ",")
		if len(nums) != 2 {
			middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
			return
		}
		start, _ := strconv.Atoi(nums[0])
		end, _ := strconv.Atoi(nums[1])
		userIds, count = userDao.GetUserIdsByCreate(int64(start), int64(end), params.Limit, params.Offset)
	} else if params.Type == util.QryUserByDevice { // 设备类型
		device, _ := strconv.Atoi(params.Qry)
		if device < 0 || device > 3 {
			middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
			return
		}
		userIds, count = commonDao.GetDeviceUsers(int32(device), params.Limit, params.Offset)
	} else if params.Type == util.QryUserByOnline { // 在线状态
		status, _ := strconv.Atoi(params.Qry)
		onlineUIds := getOnlineUserIds()
		if status == 1 { // 在线
			count = int32(len(onlineUIds))
			if params.Limit*(params.Offset+1) >= count {
				userIds = onlineUIds[params.Limit*params.Offset:]
			} else {
				userIds = onlineUIds[params.Limit*params.Offset : params.Limit*(params.Offset+1)]
			}
		} else { // 离线
			userIds, count = userDao.GetUserIdsNotIn(onlineUIds, params.Limit, params.Offset)
		}
	} else if params.Type == util.QryUserByOfficial { // 查询客服
		userIds, count = userDao.GetUserIdsOfficial(params.Limit, params.Offset)
	} else {
		userIds, count = userDao.GetUserIdsDefault(params.Limit, params.Offset)
	}

	users := QryUserInfos(userIds)

	data := map[string]interface{}{
		"user":  users,
		"count": count,
	}
	middleware.ResponseSuccess(c, data)

}

// 点击栏位动作：7.列出拥有群组的列表；8.列出有管理员权限群组的列表；
func (service *UserController) GetChatByUser(c *gin.Context) {
	//params := &dto.UserChat{}
	//if err := c.ShouldBind(params); err != nil {
	//	middleware.ResponseError(c, 500, "系统错误", err)
	//	return
	//}

	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	params, _ := bindData.(*dto.UserChat)
	if params.UserId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	if params.Limit == 0 {
		params.Limit = 20
	}

	queryType := []int32{1, 2, 3}
	if params.ChatType == 1 { // 查群
		queryType = []int32{1, 2}
	} else if params.ChatType == 2 { // 查频道
		queryType = []int32{3}
	} else if params.ChatType == 3 {
		queryType = []int32{4}
	}

	partDao := dao.GetChatParticipantDAO()

	var chatIds []int32
	var count int32
	if params.Type == 0 { // 查询拥有的群
		chatIds, count = partDao.GetChatIdsByCreator(params.UserId, params.Limit, params.Offset, queryType)
	} else if params.Type == 1 { // 查询管理的群
		chatIds, count = partDao.GetChatByManage(params.UserId, params.Limit, params.Offset, queryType)
	} else if params.Type == 2 { // 所在的群
		chatIds, count = partDao.GetChatPart(params.UserId, params.Limit, params.Offset, queryType)
	} else {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}

	chats := chatapi.ChatInfo(chatIds)
	data := map[string]interface{}{
		"chat":  chats,
		"count": count,
	}
	middleware.ResponseSuccess(c, data)

}

func QryUserInfos(userIds []int32) []map[string]interface{} {
	res := make([]map[string]interface{}, 0, 0)
	if len(userIds) == 0 {
		return res
	}

	userDao := dao.GetUserDAO()
	partDao := dao.GetChatParticipantDAO()
	commonDao := dao.GetCommonDAO()
	noteDao := dao.GetNoteDAO()
	bannedDao := dao.GetBannedDAO()

	//1.名称；2.POP ID；3.绑定电话号；4.注册日期；
	userInfos := userDao.GetUsers(userIds)

	// 获取用户所在群组数量，拥有的群数量，管理的群数量
	userChatNum := partDao.GetUserChatNum(userIds)

	// 最后活跃日期
	userLasttime := commonDao.GetUserActiveTime(userIds)

	// 当前在线状态
	status := userStatus(userIds)

	// 电子邮箱
	userEmail := commonDao.GetUserEmail(userIds)

	// 标注内容
	userNote := noteDao.GetNote(userIds)

	// 封禁内容
	banneds := bannedDao.GetBanneds(userIds)

	// 用户最后登录设备和ip
	deviceInfo := userDevice(userIds)

	for _, u := range userInfos {
		m := make(map[string]interface{})
		m["user_id"] = u.Id
		m["first_name"] = u.FirstName
		m["last_name"] = u.LastName
		m["username"] = u.Username
		m["phone"] = u.Phone
		m["add_time"] = u.AddTime

		m["join_chat_num"] = userChatNum[u.Id]["normal_num"]
		m["create_chat_num"] = userChatNum[u.Id]["create_num"]
		m["manage_chat_num"] = userChatNum[u.Id]["manage_num"]

		m["last_active_time"] = userLasttime[u.Id]
		m["email"] = userEmail[u.Id]
		m["online"] = status[u.Id]

		if u.Support {
			m["is_kefu"] = 1
		} else {
			m["is_kefu"] = 0
		}

		noteM := make(map[string]interface{})
		noteMap, exit := userNote[u.Id]
		if exit {
			labels := strings.Split(noteMap["labels"], ",")
			n := make([]map[string]interface{}, 0, len(labels))
			for _, label := range labels {
				lmap := make(map[string]interface{})
				l := strings.Split(label, "_@_")
				if len(l) == 2 {
					id, _ := strconv.Atoi(l[0])
					name := l[1]
					lmap["id"] = id
					lmap["name"] = name
					n = append(n, lmap)
				}
			}
			noteM["note"] = n
			noteM["content"] = noteMap["note"]
			m["note"] = noteM
		} else {
			m["note"] = nil
		}

		//m["label_id"] = userNote[u.Id]["label_id"]
		//m["note"] = userNote[u.Id]["note"]

		banned := banneds[u.Id]
		m["banned_state"] = banned.State // 1：封禁， 0：未封禁
		if len(banneds) != 0 {
			m["banned"] = banned
		}

		//m["ip"] = device[u.Id]["ip"]
		//m["model"] = device[u.Id]["model"]
		m["device"] = deviceInfo[u.Id]
		res = append(res, m)
	}
	return res
}

func userStatus(userIds []int32) map[int32]bool {
	status := make(map[int32]bool, 0)
	for _, id := range userIds {
		authIds, _ := redis_client.RedisCache.SMembers(fmt.Sprintf("auth:%d:online_set", id)).Result()
		if authIds == nil || len(authIds) == 0 {
			status[id] = false
		} else {
			status[id] = true
		}
	}
	return status
}

func userDevice(userIds []int32) map[int32][]map[string]interface{} {
	res := make(map[int32][]map[string]interface{}, len(userIds))
	bannedInfoDao := dao.GetBannedInfoDAO()
	for _, uId := range userIds {
		var uMap []map[string]interface{}
		exit, _ := redis_client.RedisCache.SIsMember("user_banned", uId).Result()
		if exit {
			// 封号的设备已经清除，直接从数据库取
			bannedInfos := bannedInfoDao.GetUerBannedInfo(uId)
			uMap = make([]map[string]interface{}, 0, len(bannedInfos))
			for _, info := range bannedInfos {
				country, region := GetIpCountryRegion(info.Ip)
				m := map[string]interface{}{
					"ip":            info.Ip,
					"model":         info.Model,
					"app_version":   info.AppVersion,
					"date_created":  info.DateCreated,
					"country":       country,
					"region":        region,
					"date_activate": info.DateActivate,
					"online":        false,
				}
				uMap = append(uMap, m)
			}
		} else {
			// sismember 判断元素是否在集合中 sismember key value
			authIds, _ := redis_client.RedisCache.ZRevRange(fmt.Sprintf("auth:%d:key_zset", uId), 0, -1).Result()
			uMap = make([]map[string]interface{}, 0, len(authIds))
			for _, auId := range authIds {
				lastIp := redis_client.RedisCache.HGet(fmt.Sprintf("auth:key:%s", auId), "ip").Val()
				lastModel := redis_client.RedisCache.HGet(fmt.Sprintf("auth:key:%s", auId), "model").Val()
				appVersion := redis_client.RedisCache.HGet(fmt.Sprintf("auth:key:%s", auId), "app_version").Val()
				date, _ := redis_client.RedisCache.HGet(fmt.Sprintf("auth:key:%s", auId), "date_created").Int()
				active, _ := redis_client.RedisCache.HGet(fmt.Sprintf("auth:key:%s", auId), "date_activate").Int()
				country, region := GetIpCountryRegion(lastIp)

				// gwId先为1和2
				gw1Online := redis_client.RedisCache.SIsMember(fmt.Sprintf("auth:%d:online_set", uId), fmt.Sprintf("1@%s", auId)).Val()
				gw2Online := redis_client.RedisCache.SIsMember(fmt.Sprintf("auth:%d:online_set", uId), fmt.Sprintf("2@%s", auId)).Val()
				var online bool
				if gw1Online || gw2Online {
					online = true
				}
				m := map[string]interface{}{
					"ip":            lastIp,
					"model":         lastModel,
					"app_version":   appVersion,
					"date_created":  date,
					"country":       country,
					"region":        region,
					"date_activate": active,
					"online":        online,
				}

				uMap = append(uMap, m)
			}
		}
		res[uId] = uMap
	}
	return res
}

func GetIpCountryRegion(ip string) (string, string) {
	results := ip2location.Get_all(ip)
	return results.Country_long, results.Region
}

func getOnlineUserIds() []int32 {
	uStrIds, _ := redis_client.RedisCache.SMembers("set:online_set:1").Result()
	uIds := make([]int32, 0, len(uStrIds))
	for _, uStr := range uStrIds {
		uId, _ := strconv.Atoi(uStr)
		uIds = append(uIds, int32(uId))
	}
	return uIds
}

// deivce 0:android 1:ios 2:mac 3:windows
func getDeviceUserIds(deivce int32) []int32 {
	onlineuId := getOnlineUserIds()
	res := make([]int32, 0)
	for _, uId := range onlineuId {
		authIds, _ := redis_client.RedisCache.SMembers(fmt.Sprintf("auth:%d:online_set", uId)).Result()
		for _, authId := range authIds {
			auths := strings.Split(authId, "@")
			if len(auths) != 2 {
				break
			}
			pack := redis_client.RedisCache.HGet(fmt.Sprintf("auth:key:%s", auths[1]), "lang_pack").Val()
			if pack == "android" {
				if deivce == 0 {
					res = append(res, uId)
					break
				}
			} else if pack == "ios" {
				if deivce == 1 {
					res = append(res, uId)
					break
				}
			} else if pack == "tdesktop" {
				version := redis_client.RedisCache.HGet(fmt.Sprintf("auth:key:%s", auths[1]), "sys_version").Val()
				if version[:5] == "macOS" { // mac
					if deivce == 2 {
						res = append(res, uId)
						break
					}
				} else { // windows
					if deivce == 3 {
						res = append(res, uId)
						break
					}
				}
			}
		}

	}
	return res
}

// 用户登录信息
func (service *UserController) QueryUserLogin(c *gin.Context) {
	//params := &dto.QryUser{}
	//if err := c.ShouldBind(params); err != nil {
	//	middleware.ResponseError(c, 500, "系统错误", err)
	//	return
	//}

	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	params, _ := bindData.(*dto.QryUser)
	if params.UserId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	if params.Limit == 0 {
		params.Limit = 20
	}

	logsDao := dao.GetUserLogsDAO()
	logs, count := logsDao.GetUerLogs(params.UserId, params.Limit, params.Offset)
	logInfo := make([]map[string]interface{}, 0, len(logs))
	for _, l := range logs {
		loc, _ := time.LoadLocation("Local")
		var timeInt int64
		if strings.Contains(l.CreatedAt,"+") {
			t, _ := time.ParseInLocation("2006-01-02T15:04:05+08:00", l.CreatedAt, loc)
			timeInt = t.Unix()
		} else {
			t, _ := time.ParseInLocation("2006-01-02 15:04:05", l.CreatedAt, loc)
			timeInt = t.Unix()
		}

		log := map[string]interface{}{
			"user_id":     l.UserId,
			"auth_id":     l.AuthId,
			"ip":          l.Ip,
			"model":       l.DeviceModel,
			"app_version": l.AppVersion,
			"time":        timeInt,
		}
		logInfo = append(logInfo, log)
	}

	data := map[string]interface{}{
		"login_info": logInfo,
		"count":      count,
	}
	middleware.ResponseSuccess(c, data)

}
