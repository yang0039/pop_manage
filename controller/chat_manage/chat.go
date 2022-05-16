package chat_manage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"pop-api/baselib/logger"
	"pop-api/baselib/util"
	"pop-api/dal/dao"
	"pop-api/dto"
	"pop-api/middleware"
	"pop-api/public"
	"strconv"
	"strings"
)

// 通过不通条件查询群信息
// 强制解散群组？？             -- 后面再考虑
// 查看群对话历史信息？？        -- 后面再考虑
// 查看群成员信息
// 编辑标注

// 更新群备注
func (service *ChatController) AddChatNote(c *gin.Context) {
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
	noteDao.AddNote(labelIds, public.PEER_CHAT, params.PeerId, params.Note)
	middleware.ResponseSuccess(c, "")
}

// 1.名称；2.國码；3.群组人数(范围)；4.拥有者ID；5.管理者ID；6.标注关键字
// 7.群id 8.群状态 9.群成员 10.群标签 11.活跃日期 12.创建日期
// 查询群信息
func (service *ChatController) GetChatInfo(c *gin.Context) {
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

	chatDao := dao.GetChatDAO()
	noteDao := dao.GetNoteDAO()
	chatPartDao := dao.GetChatParticipantDAO()
	comDao := dao.GetCommonDAO()
	statusDao := dao.GetPeerStatusDAO()

	queryType := []int32{1,2,3}
	if params.ChatType == 1 {          // 查群
		queryType = []int32{1,2}
	} else if params.ChatType == 2 {   // 查频道
		queryType = []int32{3}
	} else if params.ChatType == 3 {
		queryType = []int32{4}
	}

	var chatIds []int32
	var count int32
	if params.Type == util.QryChatByName {
		if params.Qry == "" {
			middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
			return
		}
		chatIds, count = chatDao.GetChatIdsByTitle(params.Qry, queryType, params.Limit, params.Offset)
	} else if params.Type == util.QryChatByCountry {

	} else if params.Type == util.QryChatByNum { // 人数范围
		nums := strings.Split(params.Qry, ",")
		if len(nums) != 2 {
			middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
			return
		}
		min, _ := strconv.Atoi(nums[0])
		max, _ := strconv.Atoi(nums[1])
		if max == 0 || max < min {
			middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
			return
		}
		chatIds, count = chatPartDao.GetChatByMemNum(int32(min), int32(max), params.Limit, params.Offset, queryType)

	} else if params.Type == util.QryChatByCreator {
		cId, _ := strconv.Atoi(params.Qry)
		if cId == 0 {
			middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
			return
		}
		chatIds, count = chatPartDao.GetChatIdsByCreator(int32(cId), params.Limit, params.Offset, queryType)
	} else if params.Type == util.QryChatByManage { // 管理者id
		manId, _ := strconv.Atoi(params.Qry)
		if manId == 0 {
			middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
			return
		}
		chatIds, count = chatPartDao.GetChatByManage(int32(manId), params.Limit, params.Offset, queryType)
	} else if params.Type == util.QryChatByNote {
		chatIds, count = noteDao.GetChatByNote(params.Qry, params.Limit, params.Offset, queryType)
	} else if params.Type == util.QryChatById {
		cId, _ := strconv.Atoi(params.Qry)
		chatIds = []int32{int32(cId)}
		count = 1
	} else if params.Type == util.QryChatByStatus {   // 8.群状态
		var status []int32
		if params.Qry == "1" {   // 正常的群
			chatIds, count = statusDao.GetStatuNormalChatIds(params.Limit, params.Offset, queryType)
			//status = []int32{2}
		} else if params.Qry == "2" {
			status = []int32{2}
			chatIds, count = statusDao.GetStatuChatIds(status, params.Limit, params.Offset, queryType)
		} else if params.Qry == "3" {
			status = []int32{3}
			chatIds, count = statusDao.GetStatuChatIds(status, params.Limit, params.Offset, queryType)
		} else {
			status = []int32{4}
			chatIds, count = statusDao.GetStatuChatIds(status, params.Limit, params.Offset, queryType)
		}
	} else if params.Type == util.QryChatByUserId {   // 9.群成员
		memId, _ := strconv.Atoi(params.Qry)
		chatIds, count = chatPartDao.GetChatByMember(int32(memId), params.Limit, params.Offset, queryType)
	} else if params.Type == util.QryChatByLabel {  // 10.群标签
		labelIds := strings.Split(params.Qry, ",")
		//labelId, _ := strconv.Atoi(params.Qry)
		chatIds, count = noteDao.GetLabelChatIds(3, params.Limit, params.Offset, labelIds, queryType)
	} else if params.Type == util.QryChatByActiveDate {  // 11.活跃日期
		nums := strings.Split(params.Qry, ",")
		if len(nums) != 2 {
			middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
			return
		}
		start, _ := strconv.Atoi(nums[0])
		end, _ := strconv.Atoi(nums[1])
		chatIds, count = comDao.GetChatByActive(int64(start), int64(end), params.Limit, params.Offset, queryType)
	} else if params.Type == util.QryChatByCreateDate {  // 12.创建日期
		nums := strings.Split(params.Qry, ",")
		if len(nums) != 2 {
			middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
			return
		}
		start, _ := strconv.Atoi(nums[0])
		end, _ := strconv.Atoi(nums[1])
		chatIds, count = chatDao.GetChatIdsByCreate(int64(start), int64(end), params.Limit, params.Offset, queryType)
	} else {
		chatIds, count = chatDao.GetChatIdsDefault(queryType, params.Limit, params.Offset)
	}

	logger.LogSugar.Infof("chatIds:%v, count:%d", chatIds, count)
	chats := ChatInfo(chatIds)

	data := map[string]interface{}{
		"chat":  chats,
		"count": count,
	}
	middleware.ResponseSuccess(c, data)
}

func (service *ChatController) GetChatMembber(c *gin.Context) {
	//type Param struct {
	//	ChatId int32 `form:"chat_id"`
	//}
	//p := &dto.QryChat{}
	//if err := c.ShouldBind(p); err != nil {
	//	middleware.ResponseError(c, 500, "系统错误", err)
	//	return
	//}
	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	p, _ := bindData.(*dto.QryChat)
	if p.ChatId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", p.ChatId)))
		return
	}

	commonDao := dao.GetCommonDAO()

	chatMember := commonDao.GetChatMemberInfo(p.ChatId)
	data := map[string]interface{}{
		"chat_participant": chatMember,
	}
	middleware.ResponseSuccess(c, data)

}
