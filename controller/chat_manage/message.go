package chat_manage

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"pop-api/baselib/logger"
	"pop-api/baselib/redis_client"
	"pop-api/dal/dao"
	"pop-api/dal/dataobject"
	"pop-api/dto"
	"pop-api/middleware"
	"pop-api/mtproto"
)


//FILTER_DOCUMENT         = 1
//FILTER_PHOTO            = 2
//FILTER_VIDEO            = 3
//FILTER_URL              = 4
//FILTER_GIF              = 5
//FILTER_VOICE            = 6
//FILTER_MUSIC            = 7
//FILTER_CHATPHOTO        = 8    // messageService
//FILTER_PHONECALL        = 9
//FILTER_PHONECALL_MISSED = 10
//FILTER_ROUND_VIDEO      = 11
//FILTER_GEO              = 12
//FILTER_CONTACT          = 13
//FITER_MYMENTION         = 14

type ChatMessage struct {
	MsgId    int32    `json:"msg_id"`
	ChatType int32    `json:"chat_type"` // 群类型： 1：普通群 2：超级群 3：频道
	ChatId   int32    `json:"chat_id"`
	From     User     `json:"from"`     // 消息发送者id
	Date     int32    `json:"date"`     // 消息发送时间
	MsgType  int32    `json:"msg_type"` // 消息类型
	Message  string   `json:"message"`  // 消息内容
	//Document Document `json:"document"` // 文档
	Url      string `json:"url"`
}

type User struct {
	Id        int32  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Document struct {
	DocumentId int64  `json:"document_id"`
	MimeType   string `json:"mime_type"`
	Size       int32  `json:"size"`
}

func (service *ChatController) GetChatMessage(c *gin.Context) {
	//params := &dto.ChatMsg{}
	//if err := c.ShouldBind(params); err != nil {
	//	middleware.ResponseError(c, 500, "系统错误", err)
	//	return
	//}
	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	params, _ := bindData.(*dto.ChatMsg)
	if params.ChatId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	if params.Limit == 0 {
		params.Limit = 20
	}

	msgDao := dao.GetChannelMsgRowDAO()
	rawDao := dao.GetRawMessageRowDAO()
	chatDao := dao.GetChatDAO()
	uMsgDao := dao.GetUserMsgRowDAO()

	chat := chatDao.GetChat(params.ChatId)
	if chat == nil || chat.Id == 0 {
		middleware.ResponseError(c, 400, "该群不存在", errors.New("该群不存在"))
		return
	}
	var messages []*mtproto.Message
	msgType := make(map[int32]int32, 0)
	var count int32
	if chat.Type == 1 { // 普通群
		// 普通群以群主身份查找
		msgs, msgCount := uMsgDao.GetUserMsgRows(chat.CreatorId, params.ChatId, 3, -1, 0, int64(params.MaxTime), params.Limit, params.Offset)
		rowIds := make([]int64, 0, len(msgs))
		count = msgCount
		for _, m := range msgs {
			rowIds = append(rowIds, m.RawId)
		}
		rawMap := rawDao.GetRawMessageRows(rowIds)
		for _, m := range msgs {
			msg := User_assemble(m, rawMap[m.RawId])
			messages = append(messages, msg)
			if msg.GetConstructor() ==  mtproto.TLConstructor_CRC32_messageService {
				msgType[msg.Data2.Id] = 8
			} else {
				msgType[msg.Data2.Id] = int32(m.Type)
			}
		}

	} else { // 超级群和频道
		msgs, msgCount := msgDao.GetChannelMsgRows(int64(params.ChatId), 0, int64(params.MaxTime), params.Limit, params.Offset)
		count = msgCount
		rowIds := make([]int64, 0, len(msgs))
		for _, m := range msgs {
			rowIds = append(rowIds, m.RawId)
		}
		rawMap := rawDao.GetRawMessageRows(rowIds)
		for _, m := range msgs {
			msg := Channel_assemble(0, m, rawMap[m.RawId])
			messages = append(messages, msg)
			if msg.GetConstructor() ==  mtproto.TLConstructor_CRC32_messageService {
				msgType[msg.Data2.Id] = 8
			} else {
				msgType[msg.Data2.Id] = int32(m.Type)
			}
		}
	}

	chagMsgs := make([]ChatMessage, 0, len(messages))
	for _, m := range messages {
		fName,_ := redis_client.RedisCache.HGet(fmt.Sprintf("auth:%d:first_name", m.Data2.FromId), "0").Result()
		lName,_ := redis_client.RedisCache.HGet(fmt.Sprintf("auth:%d:last_name", m.Data2.FromId), "0").Result()
		from := User{
			Id: m.Data2.FromId,
			FirstName: fName,
			LastName: lName,
		}

		// CRC32_messageService
		chatMsg := ChatMessage{
			MsgId:    m.Data2.Id,
			ChatType: int32(chat.Type),
			ChatId:   chat.Id,
			From:     from,
			Date:     m.Data2.Date,
			MsgType:  msgType[m.Data2.Id],
			Message:  m.Data2.Message,
			Url: "",
		}
		//if msgType[m.Data2.Id] != 0 {
		//	media := m.Data2.Media
		//	if media != nil {
		//		doc := media.Data2.Document
		//		d := Document{
		//			DocumentId: doc.Data2.Id,
		//			MimeType:   doc.Data2.MimeType,
		//			Size:       doc.Data2.Size_,
		//		}
		//		chatMsg.Document = d
		//	}
		//
		//}
		chagMsgs = append(chagMsgs, chatMsg)
	}

	data := map[string]interface{}{
		"chat_message": chagMsgs,
		"count":        count,
	}
	middleware.ResponseSuccess(c, data)

}

func Channel_assemble(self_id int32, msg *dataobject.ChannelMsgRow, raw *dataobject.RawMessageRow) *mtproto.Message {
	/* 由用户信箱和原始消息拼装成完整的message */
	if raw == nil || raw.MessageBlob == nil {
		return nil
	}

	dbuf := mtproto.NewDecodeBuf(raw.MessageBlob)
	obj := dbuf.Object()
	if obj == nil {
		logger.Logger.Error("decode message_blob fail!")
	}
	if dbuf.GetError() != nil {
		logger.Logger.Fatal(dbuf.GetError().Error())
	}

	var message *mtproto.Message
	switch v := obj.(type) {
	case *mtproto.TLMessage114:
		message = obj.(*mtproto.TLMessage114).To_Message()
	case *mtproto.TLMessageService:
		message = obj.(*mtproto.TLMessageService).To_Message()
	case *mtproto.TLMessageEmpty:
		message = obj.(*mtproto.TLMessageEmpty).To_Message()
	default:
		logger.LogSugar.Errorf("assemble fail: %v", v)
	}
	message.Data2.Id = msg.MsgId
	message.Data2.Out = self_id == msg.FromId
	for _, e := range message.Data2.Entities {
		if e.Constructor == mtproto.TLConstructor_CRC32_messageEntityMentionName && self_id == e.Data2.UserId_5 {
			message.Data2.Mentioned = true
			break
		} else if e.Constructor == mtproto.TLConstructor_CRC32_inputMessageEntityMentionName && self_id == e.Data2.UserId_6.Data2.UserId {
			message.Data2.Mentioned = true
			break
		}
	}
	message.Data2.FromId = msg.FromId
	message.Data2.MediaUnread = msg.MediaUnread
	message.Data2.ReplyToMsgId = msg.ReplyToMsgId
	message.Data2.Views = raw.Views
	message.Data2.EditDate = raw.EditDate
	return message
}

func User_assemble(msg *dataobject.UserMsgRow, raw *dataobject.RawMessageRow) *mtproto.Message {
	/* 由用户信箱和原始消息拼装成完整的message */
	if raw == nil || raw.MessageBlob == nil {
		return nil
	}
	dbuf := mtproto.NewDecodeBuf(raw.MessageBlob)
	obj := dbuf.Object()
	if obj == nil {
		logger.Logger.Error("decode message_blob fail!")
		return nil
	}
	if dbuf.GetError() != nil {
		logger.LogSugar.Fatal(dbuf.GetError().Error())
	}

	var message *mtproto.Message
	switch v := obj.(type) {
	case *mtproto.TLMessage114:
		message = obj.(*mtproto.TLMessage114).To_Message()
	case *mtproto.TLMessageService:
		message = obj.(*mtproto.TLMessageService).To_Message()
	case *mtproto.TLMessageEmpty:
		message = obj.(*mtproto.TLMessageEmpty).To_Message()
	default:
		logger.LogSugar.Errorf("assemble fail: %v", v)
	}
	message.Data2.Id = msg.MsgId
	//message.Data2.Out = msg.Out()
	message.Data2.Mentioned = msg.Mentioned
	message.Data2.MediaUnread = msg.MediaUnread
	message.Data2.ReplyToMsgId = msg.ReplyToMsgId
	message.Data2.Views = raw.Views
	message.Data2.EditDate = raw.EditDate
	logger.LogSugar.Infof("message:%v", message)
	return message
}
