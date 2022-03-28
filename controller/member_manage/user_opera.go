package member_manage

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"pop-api/dal/dao"
	"pop-api/dto"
	"pop-api/middleware"
	"pop-api/public"
	"strings"
	"time"
)

// 封号
func (service *UserController) BannedUser(c *gin.Context) {
	banned := &dto.Banned{}
	if err := c.ShouldBind(banned); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if banned.UserId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", banned)))
		return
	}
	if banned.OperaType != 1 && banned.OperaType != 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", banned)))
		return
	}

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

	err := fetchdata("POST", url, nil, body, nil)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	middleware.ResponseSuccess(c, nil)
}

// 用户设置为客服
func (service *UserController) SetOfficialUser(c *gin.Context) {
	url := "http://127.0.0.1:9200/bot1614847516:12f9f726d3423000/jsonapi"
	userSet := &dto.SetUserOfficial{}
	if err := c.ShouldBind(userSet); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if userSet.UserId == 0 || (userSet.OperaType != 1 && userSet.OperaType != 2) {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", userSet)))
		return
	}

	data := map[string]int32{
		"user_id":    userSet.UserId,
		"opera_type": userSet.OperaType,
	}

	m := map[string]interface{}{
		"cmd":  2001,
		"data": data,
	}
	body, _ := json.Marshal(m)
	err := fetchdata("POST", url, nil, body, nil)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
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
