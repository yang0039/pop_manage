package message_record

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"pop-api/baselib/logger"
	"pop-api/baselib/minio_client"
	"pop-api/baselib/util"
	"pop-api/controller/_dummy/chatapi"
	"pop-api/dal/dao"
	"pop-api/dal/dataobject"
	"pop-api/dto"
	"pop-api/middleware"
	"pop-api/mtproto"
	"strings"
)

//PeerUser = 1       // 个人
//PeerChat = 2       // 群
//PeerChannel = 3    // 频道
//PeerMass = 4       // 群发

func (service *RecordController) DelFileMessages(c *gin.Context) {
	//params := &dto.MessageRecord{}
	//if err := c.ShouldBind(params); err != nil {
	//	middleware.ResponseError(c, 500, "系统错误", err)
	//	return
	//}

	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	params, _ := bindData.(*dto.RemovePeerFile)
	if params.PeerId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	if params.PeerType != 1 && params.PeerType != 2 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}

	peerT := util.PeerUser
	if params.PeerType == 2 {
		chatDao := dao.GetChatDAO()
		c := chatDao.GetChat(params.PeerId)
		if c.Type == 1 {
			peerT = util.PeerChat
		} else if c.Type == 2 || c.Type == 3 {
			peerT = util.PeerChannel
		}
	}

	// 删除文件
	DelFile(int32(peerT), params.PeerId, params.Start, params.End, true)

	middleware.ResponseSuccess(c, nil)

}

func DelFileByMsgId(PeerType, PeerId, msgId int) {
	userMsgDao := dao.GetUserMsgRowDAO()
	channelMsgDao := dao.GetChannelMsgRowDAO()
	var uMsgs *dataobject.UserMsgRow
	var cMsgs *dataobject.ChannelMsgRow
	if PeerType == util.PeerUser || PeerType == util.PeerChat {
		uMsgs = userMsgDao.GetUserFileMsgById(int32(PeerId), int32(msgId))
	} else if PeerType == util.PeerChannel {
		cMsgs = channelMsgDao.GetChannelFileMsgById(int32(PeerId), int32(msgId))
	}
	delMiniioFile([]*dataobject.UserMsgRow{uMsgs}, []*dataobject.ChannelMsgRow{cMsgs}, true, true)
}

func DelFile(PeerType, PeerId int32, start, end int64, push bool) {
	userMsgDao := dao.GetUserMsgRowDAO()
	partDao := dao.GetChatParticipantDAO()
	channelMsgDao := dao.GetChannelMsgRowDAO()

	var uMsgs []*dataobject.UserMsgRow
	var cMsgs []*dataobject.ChannelMsgRow

	if PeerType == util.PeerUser {
		// 1. 查找单人聊天的
		// 查找自己发的单人文件消息
		uMsgs = userMsgDao.GetUserFileMsgRows(PeerId, 2, 0, start, end)
		// 查找文件消息对应的对方的消息,对方的消息也要删除
		for _, m := range uMsgs {
			ms := userMsgDao.GetUserMsgRowsByRawId(m.PeerId, []int64{m.RawId})
			uMsgs = append(uMsgs, ms...)
		}

		// 2.查找普通群的
		// 查找用户在普通群里的消息，普通群里面每个人的消息都要删除
		cUMsgs := userMsgDao.GetUserFileMsgRows(PeerId, 3, 0, start, end)
		uMsgs = append(uMsgs, cUMsgs...)
		for _, m := range cUMsgs {
			userIds := partDao.GetChatPartId(m.PeerId)
			for _, uId := range userIds {
				if uId == PeerId {
					continue
				}
				ms := userMsgDao.GetUserMsgRowsByRawId(uId, []int64{m.RawId})
				uMsgs = append(uMsgs, ms...)
			}
		}
		//msgs := userMsgDao.GetUserFileMsgRows(id, 3, PeerId, start, end)

		// 3.查找超级群的
		// 查找该用户在超级群和频道发的文件消息
		chatIds := partDao.GetChannelByUser(PeerId)
		for _, cId := range chatIds {
			cUMsgs := channelMsgDao.GetFileChannelMsgRows(PeerId, cId, start, end)
			cMsgs = append(cMsgs, cUMsgs...)
		}

	} else if PeerType == util.PeerChat {
		userIds := partDao.GetChatPartId(PeerId)
		var msgs []*dataobject.UserMsgRow
		for _, id := range userIds {
			msgs = userMsgDao.GetUserFileMsgRows(id, 3, PeerId, start, end)
			uMsgs = append(uMsgs, msgs...)
		}
	} else if PeerType == util.PeerChannel {
		cMsgs = channelMsgDao.GetFileChannelMsgRows(0, PeerId, start, end)
	} else if PeerType == util.PeerMass {
		logger.LogSugar.Infof("mass peer_type:%d, peer_id:%d", PeerType, PeerId)
		// todo
	} else {
		logger.LogSugar.Errorf("unknow peer_type:%d, peer_id:%d", PeerType, PeerId)
	}
	// 删除文件
	go delMiniioFile(uMsgs, cMsgs, true, push)
}

//type PeerMsg struct {
//	PeerType int32   `json:"peer_type"`
//	PeerId   int32   `json:"peer_id"`
//	MsgIds   []int32 `json:"msg_ids"`
//}

func delMiniioFile(uMsgs []*dataobject.UserMsgRow, cMsgs []*dataobject.ChannelMsgRow, msgDel, push bool) {
	// 里面有不同用户的重复消息，过滤掉重复的消息，缩短需要解析的数据量
	rwaMap := make(map[int64]bool)
	var NewUMsgs []*dataobject.UserMsgRow
	var NewCMsgs []*dataobject.ChannelMsgRow
	for _, m := range uMsgs {
		_, ok := rwaMap[m.RawId]
		if !ok {
			NewUMsgs = append(NewUMsgs, m)
			rwaMap[m.RawId] = true
		}
	}
	for _, m := range cMsgs {
		_, ok := rwaMap[m.RawId]
		if !ok {
			NewCMsgs = append(NewCMsgs, m)
			rwaMap[m.RawId] = true
		}
	}
	m1 := GetUserOriginMessage(NewUMsgs)
	m2 := GetChannelOriginMessage(NewCMsgs)
	m1 = append(m1, m2...)

	// message里面paths会有重复，用map返回，过滤掉重复的
	paths, paths2 := ParseFilePathByMessage(m1)
	filesDao := dao.GetFilesDAO()
	for p, v := range paths {
		logger.LogSugar.Infof("remove file path:%s", p)
		err := minio_client.RemoveObject(p)
		if err != nil {
			logger.LogSugar.Errorf("delMiniioFile err path:%s, err:%v", p, err)
			continue
		}
		if v != 0 {
			// 删除文件, 将files表的status改成1
			filesDao.DelByPath(v, p)
			partId := paths2[p]
			if partId != 0 {
				// 将文件的缩略图也删除
				f := filesDao.SelectByPartId(v, partId)
				if f != nil {
					logger.LogSugar.Infof("remove file path:%s", f.FilePath)
					err := minio_client.RemoveObject(f.FilePath)
					if err != nil {
						logger.LogSugar.Errorf("delMiniioFile err path:%s, err:%v", f.FilePath, err)
						continue
					}
					filesDao.DelById(f.FileId)
				}
			}
		}
	}
	logger.LogSugar.Infof("remove file path len:%d", len(paths))

	if !msgDel { // 不需要删除消息
		return
	}

	// 删除消息
	msgMap := make(map[string][]int32)
	// cMsgMap := make(map[int32][]int32)
	// 查找所有人的消息id
	for _, m := range uMsgs {
		peerType := fmt.Sprintf("%d:%d:%d", m.UserId, m.PeerType, m.PeerId)
		_, ok := msgMap[peerType]
		if ok {
			msgMap[peerType] = append(msgMap[peerType], m.MsgId)
		} else {
			msgMap[peerType] = []int32{m.MsgId}
		}
	}
	// 查找超级群的消息id
	for _, m := range cMsgs {
		// 针对超级群和频道，只需要知道群id和消息id即可
		peerType := fmt.Sprintf("0:4:%d", m.ChatId)
		_, ok := msgMap[peerType]
		if ok {
			msgMap[peerType] = append(msgMap[peerType], m.MsgId)
		} else {
			msgMap[peerType] = []int32{m.MsgId}
		}
	}

	for k, mIds := range msgMap {
		logger.LogSugar.Infof(fmt.Sprintf("peerKey:%s, msgIds:%v", k, mIds))
	}


	data := map[string]interface{}{
		"peer_msgs":    msgMap,
		"push": push,
	}
	//databody, _ := json.Marshal(data)
	// 调用pop服务接口，删除
	m := map[string]interface{}{
		"cmd":  util.DelMsg,
		"data": data,
	}
	body, _ := json.Marshal(m)
	err := util.Fetchdata("POST", util.Url, nil, body, nil)
	if err != nil {
		logger.LogSugar.Errorf("delMiniioFile err:%v", err)
	}
}

func GetUserOriginMessage(msgs []*dataobject.UserMsgRow) []*mtproto.Message {
	res := make([]*mtproto.Message, 0, len(msgs))
	if msgs == nil {
		return res
	}
	rawDao := dao.GetRawMessageRowDAO()
	rowIds := make([]int64, 0, len(msgs))
	for _, m := range msgs {
		rowIds = append(rowIds, m.RawId)
	}
	rawMap := rawDao.GetRawMessageRows(rowIds)
	for _, m := range msgs {
		msg := chatapi.User_assemble(m, rawMap[m.RawId])
		if msg == nil {
			continue
		}
		res = append(res, msg)
	}
	return res
}

func GetChannelOriginMessage(msgs []*dataobject.ChannelMsgRow) []*mtproto.Message {
	res := make([]*mtproto.Message, 0, len(msgs))
	if msgs == nil {
		return res
	}
	rawDao := dao.GetRawMessageRowDAO()

	rowIds := make([]int64, 0, len(msgs))
	for _, m := range msgs {
		rowIds = append(rowIds, m.RawId)
	}
	rawMap := rawDao.GetRawMessageRows(rowIds)
	for _, m := range msgs {
		msg := chatapi.Channel_assemble(0, m, rawMap[m.RawId])
		if msg == nil {
			continue
		}
		res = append(res, msg)
	}
	return res
}

func ParseFilePathByMessage(messages []*mtproto.Message) (map[string]int32, map[string]int64) {
	photoDao := dao.GetPhotoDAO()
	documentDao := dao.GetDocumentDAO()

	res := make(map[string]int32, len(messages))
	res2 := make(map[string]int64, len(messages))
	//fMap := make(map[string]dto.PeerMsg)
	for _, m := range messages {
		//logger.LogSugar.Infof("m:%v", m)
		fromId := m.Data2.FromId

		media := m.Data2.Media
		if media == nil || m.Data2.FwdFrom != nil {   // 被转发过来的文件不删除
			continue
		}
		//var fPath string
		switch media.Constructor {
		case mtproto.TLConstructor_CRC32_messageMediaPhoto114:
			photoId := media.Data2.Photo_1.Data2.Id
			//logger.LogSugar.Infof("photoId:%d", photoId)
			photos := photoDao.SelectPhotosById(photoId)
			for _, p := range photos {
				if p.LocalId == 0 {
					res[p.FilePath] = fromId
				} else {
					res[p.FilePath] = 0
				}
			}
		case mtproto.TLConstructor_CRC32_messageMediaDocument114:
			documentId := media.Data2.Document.Data2.Id
			document := documentDao.SelectById(documentId)
			res[document.FilePath] = fromId

			// file里面缩略图也要标记删除，通过document表里面thumb_id查找到对应的photo，
			// 再通过photo里面的file_part_id找到对应的用户
			//logger.LogSugar.Infof("ParseFilePathByMessage, ThumbId:%d", document.ThumbId)
			if document.ThumbId != 0 {
				p := photoDao.SelectByPhotoId(document.ThumbId, 0)
				//logger.LogSugar.Infof("ParseFilePathByMessage, p:%v", p)
				res2[p.FilePath] = p.FilePartId
				//logger.LogSugar.Infof("ParseFilePathByMessage, FilePath:%s, FilePartId:%d", p.FilePath, p.FilePartId)
			}
		}
	}
	return res, res2
}

func RemoveFileByUrl(url string) error {
	urls := strings.SplitAfterN(url, "/", 4)
	if len(urls) == 0 {
		return errors.New("invalid, url:" + url)
	}
	path := urls[len(urls)-1]
	if path == "" {
		return errors.New("invalid, url:" + url)
	}
	err := minio_client.RemoveObject("/" + path)
	return err
}
