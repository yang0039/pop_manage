package chat_manage

import (
	"pop-api/dal/dao"
	"strconv"
	"strings"
	"time"
)

func ChatInfo(chatIds []int32) []map[string]interface{}{
	if len(chatIds) == 0 {
		return make([]map[string]interface{}, 0, 0)
	}
	commomDao := dao.GetCommonDAO()
	noteDao := dao.GetNoteDAO()
	peerDao := dao.GetPeerStatusDAO()

	//chatMap := make(map[int32]bool, len(chatIds7))
	//for _, chatId := range chatIds7 {
	//	chatMap[chatId] = true
	//}

	chats := commomDao.GetChatBaseInfo(chatIds)
	//chatIds := make([]int32, 0, len(chats))
	//for _, m := range chats {
	//	chatId, _ := m["chat_id"].(int32)
	//	m["is_active"] = chatMap[chatId]
	//	chatIds = append(chatIds, chatId)
	//}

	// 获取群最后活跃日期
	chatDate := commomDao.GetChatLastActiveDate(chatIds)

	// 群状态
	chatStatus := peerDao.GetChatStatus(chatIds)

	// 获取群的备注
	note := noteDao.GetNote(chatIds)

	now := time.Now().Unix()

	for _, m := range chats {
		chatId, _ := m["chat_id"].(int32)
		m["last_date"] = chatDate[chatId]
		noteM := make(map[string]interface{})
		noteMap,exit := note[chatId]
		if exit {
			labels := strings.Split(noteMap["labels"], ",")
			n := make([]map[string]interface{}, 0, len(labels))
			for _, label := range labels {
				lmap := make(map[string]interface{})
				l := strings.Split(label, "_@_")
				if len(l) == 2 {
					id,_ := strconv.Atoi(l[0])
					name := l[1]
					lmap["id"] = id
					lmap["name"] = name
					n = append(n, lmap)
				}
			}
			noteM["note"] = n
			noteM["content"] = noteMap["note"]
			m["note"] = noteM
		} else {
			m["note"] = nil
		}
		status := chatStatus[chatId]
		statusMap := map[string]interface{}{
			"status": 1,
			"util": 0,
			"content": "",
		}
		if status != nil {
			if status.Status == 3 && status.Util < now {
				statusMap["status"] = 1
				go peerDao.DelStatus(chatId)
			} else {
				statusMap["status"] = status.Status
				statusMap["util"] = status.Util
				statusMap["content"] = status.Note
			}
		}
		m["status"] = statusMap
	}
	return chats
}
