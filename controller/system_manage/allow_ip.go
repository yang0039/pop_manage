package system_manage

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"pop-api/dal/dao"
	"pop-api/dto"
	"pop-api/middleware"
)

func (service *SystemController) AddAllowIp(c *gin.Context) {
	params := &dto.AllowIp{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Ip == "" {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	userName, _ := c.Get("user_name")
	name, _ := userName.(string)

	allowDao := dao.GetAllowIpDAO()
	id := allowDao.AddAllowIp(name, params.Ip)
	res := map[string]int32{
		"id": id,
	}
	middleware.ResponseSuccess(c, res)

}

func (service *SystemController) DelAllowIp(c *gin.Context) {
	params := &dto.AllowIp{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Id == 0 || params.Pwd == "" {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	accId, _ := c.Get("account_id")
	id, _ := accId.(int32)
	pwdRes := VerifyPwd(id, params.Pwd)
	if pwdRes == 1 {
		middleware.ResponseError(c, 400, "密码错误", errors.New("密码错误"))
		return
	} else if pwdRes == 2 {
		middleware.ResponseError(c, 400, "密码已过期", errors.New("密码已过期"))
		return
	}


	allowDao := dao.GetAllowIpDAO()
	allowDao.DelAllowIp(params.Id)

	res := map[string]int32{
		"id": params.Id,
	}
	middleware.ResponseSuccess(c, res)
}

func (service *SystemController) GetAllowIp(c *gin.Context) {
	allowDao := dao.GetAllowIpDAO()

	ips,count := allowDao.GetAllowIp()

	res := map[string]interface{}{
		"white_list": ips,
		"count": count,
	}
	middleware.ResponseSuccess(c, res)
}
