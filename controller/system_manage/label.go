package system_manage

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"pop-api/dal/dao"
	"pop-api/dto"
	"pop-api/middleware"
)

func (service *SystemController) AddLabel(c *gin.Context) {
	params := &dto.Lable{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Name == "" {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	userName, _ := c.Get("user_name")
	name, _ := userName.(string)

	labbelDao := dao.GetLabelDAO()
	id := labbelDao.AddLabel(params.Name, name)

	res := map[string]int32{
		"id": id,
	}
	middleware.ResponseSuccess(c, res)
}

func (service *SystemController) GetLabel(c *gin.Context) {
	labbelDao := dao.GetLabelDAO()
	labels := labbelDao.GetLabels()

	res := make([]map[string]interface{}, 0, len(labels))
	for _, lab := range labels {
		l := map[string]interface{}{
			"id": lab.Id,
			"name": lab.LabelName,
		}
		res = append(res, l)
	}
	middleware.ResponseSuccess(c, res)
}

func (service *SystemController) DelLabel(c *gin.Context) {
	params := &dto.Lable{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Id == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}

	accId, _ := c.Get("account_id")
	id, _ := accId.(int32)

	// 校验密码
	pwdRes := VerifyPwd(id, params.Pwd)
	if pwdRes == 1 {
		middleware.ResponseError(c, 400, "密码错误", errors.New("密码错误"))
		return
	} else if pwdRes == 2 {
		middleware.ResponseError(c, 400, "密码已过期", errors.New("密码已过期"))
		return
	}


	labbelDao := dao.GetLabelDAO()
	labbelDao.DelLabel(params.Id)

	res := map[string]int32{
		"id": params.Id,
	}
	middleware.ResponseSuccess(c, res)
}

func (service *SystemController) UpdateLabel(c *gin.Context) {
	params := &dto.Lable{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Id == 0 || params.Name == "" {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}

	labbelDao := dao.GetLabelDAO()
	labbelDao.UpdateLabel(params.Id, params.Name)

	res := map[string]int32{
		"id": params.Id,
	}
	middleware.ResponseSuccess(c, res)
}

func (service *SystemController) GetLabbelNote(c *gin.Context) {
	labbelDao := dao.GetLabelDAO()
	daoNote := dao.GetNoteDAO()
	labels := labbelDao.GetLabels()
	lableCount := daoNote.GetLabbelNoteCount()

	res := make([]map[string]interface{}, 0, len(labels))
	for _, lab := range labels {
		l := map[string]interface{}{
			"id": lab.Id,
			"name": lab.LabelName,
			"user_count": lableCount[fmt.Sprintf("%d_2", lab.Id)],
			"chat_count": lableCount[fmt.Sprintf("%d_3", lab.Id)],
		}
		res = append(res, l)
	}
	middleware.ResponseSuccess(c, res)
}