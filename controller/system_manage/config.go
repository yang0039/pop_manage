package system_manage

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"pop-api/dal/dao"
	"pop-api/dto"
	"pop-api/middleware"
)

func (service *SystemController) GetConfig(c *gin.Context) {
	params := &dto.Config{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Key == "" {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	commomDao := dao.GetCommonDAO()
	value := commomDao.GetConfig(params.Key)
	res := map[string]interface{}{
		"value": value,
	}
	middleware.ResponseSuccess(c, res)
}

func (service *SystemController) EditConfig(c *gin.Context) {
	params := &dto.Config{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Key == "" || params.Value == "" {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}

	commomDao := dao.GetCommonDAO()
	v := commomDao.GetConfig(params.Key)
	if v == "" {
		commomDao.AddConfig(params.Key, params.Value)
	} else {
		commomDao.UpdateConfig(params.Key, params.Value)
	}
	middleware.ResponseSuccess(c, "")
}
