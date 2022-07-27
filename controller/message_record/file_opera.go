package message_record

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"pop-api/baselib/util"
	"pop-api/dal/dao"
	"pop-api/dal/dataobject"
	"pop-api/dto"
	"pop-api/middleware"
)

func (service *RecordController) DelFile(c *gin.Context) {
	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	params, _ := bindData.(*dto.RemoveFileMessage)
	if params.PeerId == 0 || len(params.MsgIds) == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	channelMsgDao := dao.GetChannelMsgRowDAO()
	userMsgDao := dao.GetUserMsgRowDAO()
	partDao := dao.GetChatParticipantDAO()
	var uMsgs []*dataobject.UserMsgRow
	var cMsgs []*dataobject.ChannelMsgRow
	peerType := params.PeerType

	if params.PeerType == util.PeerChat {
		chatDao := dao.GetChatDAO()
		chat := chatDao.GetChat(params.PeerId)
		if chat.Type == 1 {
			peerType = util.PeerChat
		} else if chat.Type == 2 || chat.Type == 3 {
			peerType = util.PeerChannel
		}
	}

	if peerType == util.PeerUser {
		uMsgs = userMsgDao.GetUserMsgRowsById(params.UserId, params.MsgIds)
		// 查找文件消息对应的对方的消息,对方的消息也要删除
		for _, m := range uMsgs {
			ms := userMsgDao.GetUserMsgRowsByRawId(m.PeerId, []int64{m.RawId})
			uMsgs = append(uMsgs, ms...)
		}
	} else if peerType == util.PeerChat {
		uMsgs = userMsgDao.GetUserMsgRowsById(params.UserId, params.MsgIds)
		userIds := partDao.GetChatPartId(params.PeerId)
		rawIds := make([]int64, 0, len(uMsgs))
		for _, m := range uMsgs {
			rawIds = append(rawIds, m.RawId)
		}
		for _, u := range userIds {
			uMsgs2 := userMsgDao.GetUserMsgRowsByRawId(u, rawIds)
			uMsgs = append(uMsgs, uMsgs2...)
		}

	} else if peerType == util.PeerChannel {
		cMsgs = channelMsgDao.GetChannelMsgRowsById(params.PeerId, params.MsgIds)
	} else {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}

	// 删除文件
	go delMiniioFile(uMsgs, cMsgs, true, true)
	middleware.ResponseSuccess(c, nil)
}