package member_manage

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ttacon/libphonenumber"
	"io/ioutil"
	"net/http"
	"net/url"
	"pop-api/baselib/redis_client"
	"pop-api/baselib/util"
	"pop-api/dal/dao"
	"pop-api/dal/dataobject"
	"pop-api/dto"
	"pop-api/middleware"
	"pop-api/public"
	"strconv"
	"strings"
	"time"
)

// 封号
func (service *UserController) BannedUser(c *gin.Context) {
	//banned := &dto.Banned{}
	//if err := c.ShouldBind(banned); err != nil {
	//	middleware.ResponseError(c, 500, "系统错误", err)
	//	return
	//}

	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	banned, _ := bindData.(*dto.Banned)
	if banned.UserId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", banned)))
		return
	}
	if banned.OperaType != 1 && banned.OperaType != 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", banned)))
		return
	}
	bannedInfoDao := dao.GetBannedInfoDAO()

	// 记录封号时的用户设备，不然用户封号会删除设备
	var lastIp, lastModel string
	var lastTime int64
	authIds, _ := redis_client.RedisCache.ZRevRange(fmt.Sprintf("auth:%d:key_zset", banned.UserId), 0, -1).Result()
	for _, auId := range authIds {
		aId, _ := strconv.Atoi(auId)
		if aId == 0 {
			continue
		}
		key := fmt.Sprintf("auth:key:%d", aId)
		var info dataobject.BannedInfo
		info.UserId = banned.UserId
		info.AuthId = int64(aId)
		info.Model, _ = redis_client.RedisCache.HGet(key, "model").Result()
		info.SystemVersion, _ = redis_client.RedisCache.HGet(key, "sys_version").Result()
		info.AppVersion, _ = redis_client.RedisCache.HGet(key, "app_version").Result()
		info.SystemLangCode, _ = redis_client.RedisCache.HGet(key, "system_lang_code").Result()
		info.LangPack, _ = redis_client.RedisCache.HGet(key, "lang_pack").Result()
		info.LangCode, _ = redis_client.RedisCache.HGet(key, "lang_code").Result()
		info.Ip, _ = redis_client.RedisCache.HGet(key, "ip").Result()
		layer, _ := redis_client.RedisCache.HGet(key, "layer").Int()
		info.Layer = int32(layer)
		info.DateCreated, _ = redis_client.RedisCache.HGet(key, "date_created").Int64()
		info.DateActivate, _ = redis_client.RedisCache.HGet(key, "date_activate").Int64()
		redis_client.RedisCache.HGet(key, "ip").Result()
		bannedInfoDao.AddbannedInfo(&info)
		//fmt.Println("info.Ip=", info.Ip)
		//fmt.Println("info.Model=", info.Model)
		if info.DateActivate > lastTime {
			lastTime = info.DateActivate
			lastIp = info.Ip
			lastModel = info.Model
		}
	}

	if lastIp == "" {
		lastIp, lastModel = dao.GetUserLogsDAO().GetUerLastLogs(banned.UserId)
	}
	//fmt.Println("lastIp=", lastIp)
	//fmt.Println("lastModel=", lastModel)

	url := "http://127.0.0.1:9200/bot1614847516:12f9f726d3423000/getWallpaper"
	bannedReq := dto.BannedReq{
		UserId:    banned.UserId,
		OperaType: banned.OperaType,
		Opera:     "admin",
		Reason:    banned.Reason,
	}
	req := dto.ApiReq{
		From:   "manager",
		Method: public.BannedUser,
		Data:   bannedReq,
	}
	m := map[string]interface{}{
		"data": req,
	}
	body, _ := json.Marshal(m)

	err = fetchdata("POST", url, nil, body, nil)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	// 记录操作
	accId, _ := c.Get("account_id")
	id, _ := accId.(int32)
	var opera string
	if banned.OperaType == 1 {
		opera = "封号"
		if lastIp != "" {
			opera += fmt.Sprintf(", 最后ip:%s, 最后设备:%s", lastIp, lastModel)
		}
	} else {
		opera = "解封"
	}
	if banned.Reason != "" {
		opera += ", 原因:" + banned.Reason
	}
	uo := &dataobject.UserOpera{
		AccountId:    id,
		UserId:       banned.UserId,
		OperaType:    util.Updatebanned,
		OperaContent: opera,
		AddTime:      time.Now().Unix(),
	}
	dao.GetUserOperaDAO().AddOperaRecords(uo)
	middleware.ResponseSuccess(c, nil)
}

// 用户设置为客服
func (service *UserController) SetOfficialUser(c *gin.Context) {
	//url := "http://127.0.0.1:9200/bot1614847516:12f9f726d3423000/jsonapi"
	//userSet := &dto.SetUserOfficial{}
	//if err := c.ShouldBind(userSet); err != nil {
	//	middleware.ResponseError(c, 500, "系统错误", err)
	//	return
	//}e

	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	userSet, _ := bindData.(*dto.SetUserOfficial)
	if userSet.UserId == 0 || (userSet.OperaType != 1 && userSet.OperaType != 2) {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", userSet)))
		return
	}

	data := map[string]int32{
		"user_id":    userSet.UserId,
		"opera_type": userSet.OperaType,
	}

	m := map[string]interface{}{
		"cmd":  util.SetUserOfficial,
		"data": data,
	}
	body, _ := json.Marshal(m)
	err = fetchdata("POST", util.Url, nil, body, nil)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}

	// 记录操作
	accId, _ := c.Get("account_id")
	id, _ := accId.(int32)
	var opera string
	if userSet.OperaType == 1 {
		opera = "设置客服"
	} else if userSet.OperaType == 2 {
		opera = "取消客服"
	}
	uo := &dataobject.UserOpera{
		AccountId:    id,
		UserId:       userSet.UserId,
		OperaType:    util.UpdateOfficial,
		OperaContent: opera,
		AddTime:      time.Now().Unix(),
	}
	dao.GetUserOperaDAO().AddOperaRecords(uo)

	middleware.ResponseSuccess(c, nil)
}

// 查询用户封号记录
func (service *UserController) GetUserBanned(c *gin.Context) {
	params := &dto.QryUser{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Limit == 0 {
		params.Limit = 20
	}
	bannedDao := dao.GetBannedDAO()
	banneds, count := bannedDao.GetUserBanneds(params.UserId, params.Limit, params.Offset)
	data := map[string]interface{}{
		"banned": banneds,
		"count":  count,
	}

	middleware.ResponseSuccess(c, data)
}

// 释放popid
func (service *UserController) DelPopId(c *gin.Context) {

}

// 释放电话注册（释放后原账号自动登出，下次登录要重新验证电话）
func (service *UserController) DelPopId2(c *gin.Context) {

}

// 释放电子邮箱

// 账号删除（相当于账号超过期限未登录的处理逻辑）
func (service *UserController) DelAccount(c *gin.Context) {
	type Param struct {
		UserId int32 `json:"user_id" form:"user_id"`
	}
	p := &Param{}
	if err := c.ShouldBind(p); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if p.UserId == 0 || p.UserId == 333000 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", p.UserId)))
		return
	}

	userDao := dao.GetUserDAO()

	// 查看用户是否存在
	if !userDao.CheckUser(p.UserId) {
		middleware.ResponseError(c, 400, "用户不存在", errors.New(fmt.Sprintf("user not find, id:%d", p.UserId)))
		return
	}

	// 查找或创建该用户的机器人
	token := GetBot(p.UserId)

	// 调用删除账号机器人

	// todo 删除群成员

	data := map[string]string{
		"token": token,
	}
	middleware.ResponseSuccess(c, data)

}

// 更新用户username
func (service *UserController) UpdateUserName(c *gin.Context) {
	//url := "http://127.0.0.1:9200/bot1614847516:12f9f726d3423000/jsonapi"
	//userName := &dto.UpdateUserName{}
	//if err := c.ShouldBind(userName); err != nil {
	//	middleware.ResponseError(c, 500, "系统错误", err)
	//	return
	//}

	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}

	userName, _ := bindData.(*dto.UpdateUserName)
	if userName.UserId == 0 || userName.UserName == "" {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", userName)))
		return
	}

	if !IsValidUsername(userName.UserName) {
		middleware.ResponseError(c, 400, "无效的POP ID", errors.New("无效的POP ID"+userName.UserName))
		return
	}
	u := dao.GetUserDAO().GetUserbyUserName(userName.UserName)

	// 检查username是否存在
	if u.Username == userName.UserName {
		middleware.ResponseError(c, 400, "POP ID已存在", errors.New("POP ID已存在"))
		return
	}
	self := dao.GetUserDAO().GetUser(userName.UserId)
	//if dao.GetUserDAO().CheckUserByUsername(userName.UserName) {
	//	middleware.ResponseError(c, 400, "POP ID已存在", errors.New("POP ID已存在"))
	//	return
	//}
	data := map[string]interface{}{
		"user_id":  userName.UserId,
		"username": userName.UserName,
	}

	m := map[string]interface{}{
		"cmd":  util.SetUserUserName,
		"data": data,
	}
	body, _ := json.Marshal(m)
	err = fetchdata("POST", util.Url, nil, body, nil)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}

	// 记录操作
	accId, _ := c.Get("account_id")
	id, _ := accId.(int32)
	uo := &dataobject.UserOpera{
		AccountId:    id,
		UserId:       userName.UserId,
		OperaType:    util.UpdatePopId,
		OperaContent: fmt.Sprintf("原popId:%s, 新popId:%s", self.Username, userName.UserName),
		AddTime:      time.Now().Unix(),
	}
	dao.GetUserOperaDAO().AddOperaRecords(uo)

	middleware.ResponseSuccess(c, nil)
}

// 更新用户手机号
func (service *UserController) UpdateUserPhone(c *gin.Context) {
	//url := "http://127.0.0.1:9200/bot1614847516:12f9f726d3423000/jsonapi"
	//userPhone := &dto.UpdateUserPhone{}
	//if err := c.ShouldBind(userPhone); err != nil {
	//	middleware.ResponseError(c, 500, "系统错误", err)
	//	return
	//}

	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	userPhone, _ := bindData.(*dto.UpdateUserPhone)
	if userPhone.UserId == 0 || userPhone.Phone == "" {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", userPhone)))
		return
	}

	// 检查手机号的有效性
	if !CheckPhone(userPhone.Phone) {
		middleware.ResponseError(c, 400, "无效的手机号", errors.New(fmt.Sprintf("invalid phone:%v", userPhone.Phone)))
		return
	}
	// 检查手机号是否存在
	u := dao.GetUserDAO().GetUserbyPhone(userPhone.Phone)
	if u.Phone == userPhone.Phone {
		middleware.ResponseError(c, 400, "手机号已存在", errors.New("手机号已存在"))
		return
	}
	self := dao.GetUserDAO().GetUser(userPhone.UserId)
	//if dao.GetUserDAO().CheckUserByPhone(userPhone.Phone) {
	//	middleware.ResponseError(c, 400, "手机号已存在", errors.New("手机号已存在"))
	//	return
	//}

	data := map[string]interface{}{
		"user_id": userPhone.UserId,
		"phone":   userPhone.Phone,
	}

	m := map[string]interface{}{
		"cmd":  util.SetUserPhone,
		"data": data,
	}
	body, _ := json.Marshal(m)
	err = fetchdata("POST", util.Url, nil, body, nil)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}

	// 记录操作
	accId, _ := c.Get("account_id")
	id, _ := accId.(int32)
	uo := &dataobject.UserOpera{
		AccountId:    id,
		UserId:       userPhone.UserId,
		OperaType:    util.UpdatePhone,
		OperaContent: fmt.Sprintf("原手机号:%s, 新手机号:%s", self.Phone, userPhone.Phone),
		AddTime:      time.Now().Unix(),
	}
	dao.GetUserOperaDAO().AddOperaRecords(uo)

	middleware.ResponseSuccess(c, nil)
}

// 删除账号
func (service *UserController) DelUser(c *gin.Context) {
	//url := "http://127.0.0.1:9200/bot1614847516:12f9f726d3423000/jsonapi"
	//userPhone := &dto.UpdateUserPhone{}
	//if err := c.ShouldBind(userPhone); err != nil {
	//	middleware.ResponseError(c, 500, "系统错误", err)
	//	return
	//}

	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	userPhone, _ := bindData.(*dto.UpdateUserPhone)
	if userPhone.UserId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", userPhone)))
		return
	}

	//u := dao.GetUserDAO().GetUser(userPhone.UserId)

	data := map[string]int32{
		"user_id": userPhone.UserId,
	}

	m := map[string]interface{}{
		"cmd":  util.DelUser,
		"data": data,
	}
	body, _ := json.Marshal(m)
	err = fetchdata("POST", util.Url, nil, body, nil)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}

	// 记录操作
	accId, _ := c.Get("account_id")
	id, _ := accId.(int32)
	uo := &dataobject.UserOpera{
		AccountId:    id,
		UserId:       userPhone.UserId,
		OperaType:    util.DeleteUser,
		OperaContent: fmt.Sprintf("删除的用户id:%d", userPhone.UserId),
		AddTime:      time.Now().Unix(),
	}
	dao.GetUserOperaDAO().AddOperaRecords(uo)

	middleware.ResponseSuccess(c, nil)
}

func fetchdata(method, u string, param map[string]string, body []byte, res interface{}) error {
	// 跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Timeout: 30 * time.Second, Transport: tr}

	req, err := http.NewRequest(method, u, strings.NewReader(string(body)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	values := url.Values{}
	for k, v := range param {
		values.Set(k, v)
	}
	req.URL.RawQuery = values.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		if res != nil {
			err = json.Unmarshal(body, res)
			if err != nil {
				return err
			}
		}
	} else {
		return errors.New("服务器响应失败:" + resp.Status)
	}
	return nil
}

func CheckPhone(number string) bool {
	if number == "" {
		return false
	}
	var err error
	if number[:1] != "+" {
		number = "+" + number
	}
	var pnumber *libphonenumber.PhoneNumber
	pnumber, err = libphonenumber.Parse(number, "")
	if err != nil {
		return false
	} else {
		if !libphonenumber.IsValidNumber(pnumber) {
			if pnumber.GetCountryCode() == 86 {
				nationalNum := fmt.Sprintf("%d", *pnumber.NationalNumber)
				if len(nationalNum) == 11 && (nationalNum[:3] == "199" || nationalNum[:3] == "166" || nationalNum[:3] == "182") {
					return true
				}
			} else if pnumber.GetCountryCode() == 95 {
				nationalNum := fmt.Sprintf("%d", *pnumber.NationalNumber)
				if len(nationalNum) == 9 && (nationalNum[:2] == "66") {
					return true
				} else if len(nationalNum) == 10 && (nationalNum[:3] == "966") {
					return true
				}
			}
			return false
		} else {
			return true
		}
	}
}

func IsValidUsername(s string) bool {
	if len(s) < 5 || len(s) > 32 {
		return false
	}
	c := 0
	for i, r := range s {
		switch {
		case '_' == r && i != 0 && i != len(s)-1:
			c++
			break
		// case i > 0 && '0' <= r && r <= '9':
		case '0' <= r && r <= '9': //第一个可以为数字
			c++
			break
		case 'a' <= r && r <= 'z':
			c++
			break
		case 'A' <= r && r <= 'Z':
			c++
			break
		}
	}
	return len(s) == c
}

func (service *UserController) GetUserOpera(c *gin.Context) {
	//banned := &dto.Banned{}
	//if err := c.ShouldBind(banned); err != nil {
	//	middleware.ResponseError(c, 500, "系统错误", err)
	//	return
	//}

	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	user, _ := bindData.(*dto.QryUser)
	if user.UserId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", user)))
		return
	}
	if user.Limit == 0 {
		user.Limit = 20
	}

	accId, _ := c.Get("account_id")
	id, _ := accId.(int32)

	account := dao.GetAccountDAO().GetAllAccount()
	accountMap := make(map[int32]dataobject.Account, len(account))
	for _, a := range account {
		accountMap[a.Id] = a
	}

	oData, count := dao.GetUserOperaDAO().GetUserOperaRecords(id, user.UserId, user.Limit, user.Offset)
	opData := make([]map[string]interface{}, 0, len(oData))
	for _, od := range oData {
		d := map[string]interface{}{
			"account_id":   od.AccountId,
			"account_name": accountMap[od.AccountId].AccountName,
			"user_name":    accountMap[od.AccountId].UserName,
			"user_id":      od.UserId,
			"opera_type":   od.OperaType,
			"content":      od.OperaContent,
			"opera_time":   od.AddTime,
		}
		opData = append(opData, d)
	}

	data := map[string]interface{}{
		"opera_data": opData,
		"count":      count,
	}
	middleware.ResponseSuccess(c, data)
}

// 清除两步验证
func (service *UserController) DelUserPassword(c *gin.Context) {
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

	dao.GetCommonDAO().DelUserPassword(params.UserId)
	// 记录操作
	accId, _ := c.Get("account_id")
	id, _ := accId.(int32)
	uo := &dataobject.UserOpera{
		AccountId:    id,
		UserId:       params.UserId,
		OperaType:    util.DeleteUserPwd,
		OperaContent: fmt.Sprintf("删除两步验证, 用户id:%d", params.UserId),
		AddTime:      time.Now().Unix(),
	}
	dao.GetUserOperaDAO().AddOperaRecords(uo)

	middleware.ResponseSuccess(c, nil)
}
