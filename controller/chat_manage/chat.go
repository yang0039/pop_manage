package chat_manage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"pop-api/baselib/logger"
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
	params := &dto.NoteObj{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
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
	params := &dto.QryType{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Limit == 0 {
		params.Limit = 20
	}

	chatDao := dao.GetChatDAO()
	noteDao := dao.GetNoteDAO()
	chatPartDao := dao.GetChatParticipantDAO()
	comDao := dao.GetCommonDAO()

	var chatIds []int32
	var count int32
	if params.Type == 1 {
		if params.Qry == "" {
			middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
			return
		}
		chatIds, count = chatDao.GetChatIdsByTitle(params.Qry, params.Limit, params.Offset)
	} else if params.Type == 2 {

	} else if params.Type == 3 { // 人数范围
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
		chatIds, count = chatPartDao.GetChatByMemNum(int32(min), int32(max), params.Limit, params.Offset)

	} else if params.Type == 4 {
		cId, _ := strconv.Atoi(params.Qry)
		if cId == 0 {
			middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
			return
		}
		chatIds, count = chatDao.GetChatIdsByCreator(int32(cId), params.Limit, params.Offset)
	} else if params.Type == 5 { // 管理者id
		manId, _ := strconv.Atoi(params.Qry)
		if manId == 0 {
			middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
			return
		}
		chatIds, count = chatPartDao.GetChatByManage(int32(manId), params.Limit, params.Offset)
	} else if params.Type == 6 {
		chatIds, count = noteDao.GetChatByNote(params.Qry, params.Limit, params.Offset)
	} else if params.Type == 7 {
		cId, _ := strconv.Atoi(params.Qry)
		chatIds = []int32{int32(cId)}
		count = 1
	} else if params.Type == 8 {   // 8.群状态
		// todo
	} else if params.Type == 9 {   // 9.群成员
		memId, _ := strconv.Atoi(params.Qry)
		chatIds, count = chatPartDao.GetChatByMember(int32(memId), params.Limit, params.Offset)
	} else if params.Type == 10 {  // 10.群标签
		labelIds := strings.Split(params.Qry, ",")
		//labelId, _ := strconv.Atoi(params.Qry)
		chatIds, count = noteDao.GetLabelChatIds(3, params.Limit, params.Offset, labelIds)
	} else if params.Type == 11 {  // 11.活跃日期
		nums := strings.Split(params.Qry, ",")
		if len(nums) != 2 {
			middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
			return
		}
		start, _ := strconv.Atoi(nums[0])
		end, _ := strconv.Atoi(nums[1])
		chatIds, count = comDao.GetChatByActive(int64(start), int64(end), params.Limit, params.Offset)
	} else if params.Type == 12 {  // 12.创建日期
		nums := strings.Split(params.Qry, ",")
		if len(nums) != 2 {
			middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
			return
		}
		start, _ := strconv.Atoi(nums[0])
		end, _ := strconv.Atoi(nums[1])
		chatIds, count = chatDao.GetChatIdsByCreate(int64(start), int64(end), params.Limit, params.Offset)
	} else {
		chatIds, count = chatDao.GetChatIdsDefault(params.Limit, params.Offset)
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
	type Param struct {
		ChatId int32 `form:"chat_id"`
	}
	p := &Param{}
	if err := c.ShouldBind(p); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	//chatIdStr := c.GetHeader("chat_id")
	//fmt.Println("chatIdStr=", chatIdStr)
	//chatId,_ := strconv.Atoi(chatIdStr)
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
