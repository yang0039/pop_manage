package message_record

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"pop-api/baselib/minio_client"
	"pop-api/baselib/redis_client"
	"pop-api/baselib/util"
	"pop-api/controller/_dummy/chatapi"
	//"pop-api/controller/_dummy/memberapi"
	"pop-api/dal/dao"
	"pop-api/dal/dataobject"
	"pop-api/dto"
	"pop-api/middleware"
	"pop-api/mtproto"
	//"strconv"
)

/*

原始记录的消息类型
	MESSAGE                 = 0
	FILTER_DOCUMENT         = 1
	FILTER_PHOTO            = 2  // message.media.photo != nil
	FILTER_VIDEO            = 3  // message.media.document.attributes[x] == TL_documentAttributeVideo#ef02ce6 and round_message == false
	FILTER_URL              = 4  // message.entities[x] == TL_messageEntityUrl#6ed02538
	FILTER_GIF              = 5  // message.media.document.mime_type == "image/gif"
	FILTER_VOICE            = 6  // message.media.document.attributes[x] == TL_documentAttributeAudio#9852f9c6 and documentAttributeAudio.voice == true
	FILTER_MUSIC            = 7  // message.media.document.attributes[x] == TL_documentAttributeAudio#9852f9c6 and documentAttributeAudio.voice == false

	FILTER_CHATPHOTO        = 8  // message.action == TL_messageActionChatEditPhoto#7fcb13a8
	FILTER_PHONECALL        = 9  // message.action == TL_messageActionPhoneCall#80e11a7f
	FILTER_PHONECALL_MISSED = 10 // message.action == TL_messageActionPhoneCall#80e11a7f and reason == phoneCallDiscardReasonMissed#85e42301

	FILTER_ROUND_VIDEO      = 11 // message.media.document.attributes[x] == TL_documentAttributeVideo#ef02ce6 and round_message == false
	FILTER_GEO              = 12 // message.media == TL_messageMediaGeo#56e0d474
	FILTER_CONTACT          = 13 // message.media == TL_messageMediaContact#5e7d2f39
	FITER_MYMENTION         = 14

*/

const (
	ALL         = 0
	MESSAGE     = 1
	DOCUMENT    = 2
	PHOTO       = 3
	VIDEO       = 4
	URL         = 5
	GIF         = 6
	VOICE       = 7
	MUSIC       = 8
	ROUND_VIDEO = 9
	GEO         = 10
	CONTACT     = 11
	OTHER       = 12
)

func ToOriginMsgType(msgType int32) int32 {
	switch msgType {
	case ALL:
		return -1
	case MESSAGE:
		return 0
	case DOCUMENT:
		return 1
	case PHOTO:
		return 2
	case VIDEO:
		return 3
	case URL:
		return 4
	case GIF:
		return 5
	case VOICE:
		return 6
	case MUSIC:
		return 7
	case ROUND_VIDEO:
		return 11
	case GEO:
		return 12
	case CONTACT:
		return 13
	case OTHER:
		return -2
	default:
		return -1
	}
}

func ToApiMsgType(msgType int32) int32 {
	switch msgType {
	case 0:
		return MESSAGE
	case 1:
		return DOCUMENT
	case 2:
		return PHOTO
	case 3:
		return VIDEO
	case 4:
		return URL
	case 5:
		return GIF
	case 6:
		return VOICE
	case 7:
		return MUSIC
	case 11:
		return ROUND_VIDEO
	case 12:
		return GEO
	case 13:
		return CONTACT
	default:
		return OTHER
	}
}

/*
查询条件(组合形式)：
1.发送方，2.接收方(群或人id)，3.关键字(暂无)，4.消息类型(全部，文本，文件，图片，视频，链接，gif，语音，音乐，圆形视频，位置，联系人)
5.发送的开始和结束日期
*/

func (service *RecordController) GetMessageRecord(c *gin.Context) {
	params := &dto.MessageRecord{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Limit == 0 {
		params.Limit = 20
	}
	if params.PeerId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	chatDao := dao.GetChatDAO()
	uMsgDao := dao.GetUserMsgRowDAO()
	cMsgDao := dao.GetChannelMsgRowDAO()

	if params.End == 0 {
		params.End = 2147483647
	}
	var count int32
	var messages []*dto.Message
	//var peer dto.Peer

	msgType := ToOriginMsgType(params.MessageType)
	if params.PeerId > 100000 { // 用户id
		if params.FromId == 0 {
			middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
			return
		}

		msgs, msgCount := uMsgDao.GetUserMsgRows(params.FromId, params.PeerId,2, msgType, 0, params.End, params.Limit, params.Offset)
		count = msgCount
		messages = UserMsgToMessage(msgs)
		//peer = GetPeer(2, params.PeerId)

	} else { // 群id
		chat := chatDao.GetChat(params.PeerId)
		if chat.Type == 1 { // 普通群
			msgs, msgCount := uMsgDao.GetUserMsgRows(chat.CreatorId, params.PeerId, 3, msgType, params.Start, params.End, params.Limit, params.Offset)
			count = msgCount
			messages = UserMsgToMessage(msgs)

		} else if chat.Type == 2 || chat.Type == 3 { // 超级群或频道
			msgs, msgCount := cMsgDao.GetChannelMsgRowsByFrom(params.PeerId, msgType, params.Start, params.End, params.Limit, params.Offset)
			count = msgCount
			messages = ChannelMsgToMessage(msgs)
		}
		//peer = GetPeer(3, params.PeerId)
	}

	/*
	else {  // peerId为0
		if params.FromId == 0 {
			middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
			return
		}
		dialogs,dCount := memberapi.GetUserDialog(params.FromId, params.Limit, params.Offset)
		count = dCount
		for _, dialog := range dialogs {
			key := fmt.Sprintf("%d:%d:%d", params.FromId, dialog["peer_type"], dialog["peer_id"])
			top,_ := redis_client.RedisCache.HGet(key, "top").Result()
			topMsgId,_ := strconv.Atoi(top)
			if topMsgId == 0 {
				continue
			}
			if dialog["peer_type"] == 2 || dialog["peer_type"] == 3 { // 用户或普通群
				msgs := uMsgDao.GetUserMsgRowsById(params.FromId, []int32{int32(topMsgId)})
				messages = append(messages, UserMsgToMessage(msgs)...)
			} else {  // 超级群或频道
				msgs := cMsgDao.GetChannelMsgRowsById(dialog["peer_id"], []int32{int32(topMsgId)})
				messages = append(messages, ChannelMsgToMessage(msgs)...)
			}
		}
	}
	*/

	data := map[string]interface{} {
		"message": messages,
		//"peer": peer,
		"count":  count,
	}
	middleware.ResponseSuccess(c, data)
}

func UserMsgToMessage(msgs []*dataobject.UserMsgRow) []*dto.Message {
	rawDao := dao.GetRawMessageRowDAO()

	rowIds := make([]int64, 0, len(msgs))
	for _, m := range msgs {
		rowIds = append(rowIds, m.RawId)
	}
	rawMap := rawDao.GetRawMessageRows(rowIds)
	messages := make([]*dto.Message, 0, len(msgs))
	for _, m := range msgs {
		msg := chatapi.User_assemble(m, rawMap[m.RawId])
		msgType := ToApiMsgType(int32(m.Type))
		message := &dto.Message{
			MsgId:   msg.Data2.Id,
			From:    GetUser(msg.Data2.FromId),
			Peer:    GetPeer(rawMap[m.RawId].PeerType, rawMap[m.RawId].PeerId),
			Date:    msg.Data2.Date,
			Message: msg.Data2.Message,
			MsgType: msgType,
			Url:     getFileUrl(msgType, msg),
		}
		messages = append(messages, message)
	}

	return messages
}

func ChannelMsgToMessage(msgs []*dataobject.ChannelMsgRow) []*dto.Message {
	rawDao := dao.GetRawMessageRowDAO()

	rowIds := make([]int64, 0, len(msgs))
	for _, m := range msgs {
		rowIds = append(rowIds, m.RawId)
	}
	rawMap := rawDao.GetRawMessageRows(rowIds)
	messages := make([]*dto.Message, 0, len(msgs))
	for _, m := range msgs {
		msg := chatapi.Channel_assemble(0, m, rawMap[m.RawId])
		msgType := ToApiMsgType(int32(m.Type))
		message := &dto.Message{
			MsgId:   msg.Data2.Id,
			From:    GetUser(msg.Data2.FromId),
			Peer:    GetPeer(rawMap[m.RawId].PeerType, rawMap[m.RawId].PeerId),
			Date:    msg.Data2.Date,
			Message: msg.Data2.Message,
			MsgType: msgType,
			Url:     getFileUrl(msgType, msg),
		}
		messages = append(messages, message)
	}
	return messages
}

func GetUser(userId int32) dto.User {
	fName, _ := redis_client.RedisCache.HGet(fmt.Sprintf("auth:%d:first_name", userId), "0").Result()
	lName, _ := redis_client.RedisCache.HGet(fmt.Sprintf("auth:%d:last_name", userId), "0").Result()
	return dto.User{
		UserId:    userId,
		FirstName: fName,
		LastName:  lName,
	}
}

func GetPeer(peerType, peerId int32) dto.Peer {
	p := dto.Peer{
		//PeerType: peerType,
		PeerId: peerId,
	}
	if peerType == 2 {
		fName, _ := redis_client.RedisCache.HGet(fmt.Sprintf("auth:%d:first_name", peerId), "0").Result()
		lName, _ := redis_client.RedisCache.HGet(fmt.Sprintf("auth:%d:last_name", peerId), "0").Result()
		p.FirstName = fName
		p.LastName = lName
		p.PeerType = util.PeerUser
	} else {
		chatDao := dao.GetChatDAO()
		c := chatDao.GetChat(peerId)
		p.Title = c.Title
		if c.Type == 1 || c.Type == 2 { // 普通群和超级群
			p.PeerType = util.PeerChat
		} else if c.Type == 3 { // 频道
			p.PeerType = util.PeerChannel
		} else if c.Type == 4 {
			p.PeerType = util.PeerMass
		}
	}
	return p
}

func getFileUrl(msgType int32, message *mtproto.Message) string {
	media := message.Data2.Media
	if media == nil {
		return ""
	}

	photoDao := dao.GetPhotoDAO()
	documentDao := dao.GetDocumentDAO()
	switch msgType {
	case MESSAGE: // 普通消息
		return ""
	case DOCUMENT,VIDEO,VOICE,GIF,MUSIC,ROUND_VIDEO: // 文件
		document := media.Data2.Document
		if document == nil {
			return ""
		}
		documentId := document.Data2.Id
		d := documentDao.SelectById(documentId)
		return fmt.Sprintf("http://%s%s", minio_client.MinioIp, d.FilePath)

	case PHOTO: // 图片
		photo := media.Data2.Photo_1
		if photo == nil {
			return ""
		}
		photoId := photo.Data2.Id
		p := photoDao.SelectByPhotoId(photoId, 0)
		return fmt.Sprintf("http://%s%s", minio_client.MinioIp, p.FilePath)
	default:
		return ""
	}
}
