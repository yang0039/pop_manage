package chat_manage

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"pop-api/baselib/util"
	"pop-api/dal/dao"
	"pop-api/dal/dataobject"
	"pop-api/dto"
	"pop-api/middleware"
	"time"
)

// 删除账号
func (service *ChatController) DelChatHistory(c *gin.Context) {
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
	chat, _ := bindData.(*dto.QryChat)
	if chat.ChatId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", chat)))
		return
	}

	//u := dao.GetUserDAO().GetUser(userPhone.UserId)

	data := map[string]int32{
		"chat_id": chat.ChatId,
	}

	m := map[string]interface{}{
		"cmd":  util.DelChatHistory,
		"data": data,
	}
	body, _ := json.Marshal(m)
	err = util.Fetchdata("POST", util.Url, nil, body, nil)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}

	// 记录操作
	accId, _ := c.Get("account_id")
	id, _ := accId.(int32)
	uo := &dataobject.UserOpera{
		AccountId:    id,
		UserId:       chat.ChatId,
		OperaType:    util.DelChatHistory,
		OperaContent: fmt.Sprintf("删除记录群id:%d", chat.ChatId),
		AddTime:      time.Now().Unix(),
	}
	dao.GetUserOperaDAO().AddOperaRecords(uo)

	middleware.ResponseSuccess(c, nil)
}

// 解散群组
func (service *ChatController) DelChat(c *gin.Context) {
	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	chat, _ := bindData.(*dto.QryChat)
	if chat.ChatId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", chat)))
		return
	}

	data := map[string]int32{
		"chat_id": chat.ChatId,
	}

	m := map[string]interface{}{
		"cmd":  util.DelChat,
		"data": data,
	}
	body, _ := json.Marshal(m)
	err = util.Fetchdata("POST", util.Url, nil, body, nil)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}

	// 记录操作
	accId, _ := c.Get("account_id")
	id, _ := accId.(int32)
	uo := &dataobject.UserOpera{
		AccountId:    id,
		UserId:       chat.ChatId,
		OperaType:    util.DelChat,
		OperaContent: fmt.Sprintf("解散群组,群id:%d", chat.ChatId),
		AddTime:      time.Now().Unix(),
	}
	dao.GetUserOperaDAO().AddOperaRecords(uo)

	middleware.ResponseSuccess(c, nil)
}
