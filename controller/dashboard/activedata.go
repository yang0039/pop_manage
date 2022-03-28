package dashboard

import (
	"github.com/gin-gonic/gin"
	"pop-api/controller/_dummy/chatapi"
	"pop-api/dal/dao"
	"pop-api/middleware"
	"time"
)

// 前30日活跃数据
func (service *DashboardController) ActiveData(c *gin.Context) {

	//end := time.Now().Unix()
	//y, m, d := time.Now().Date()
	//time1 := time.Date(y, m, d, 0, 0, 0, 0, time.FixedZone("CST", 3600*8))
	//Before30Unix := time1.Add(-30 * 24 * time.Hour).Unix()

	now := time.Now()
	end := time.Date(now.Year(),  now.Month(), now.Day(), 0, 0, 0, 0, time.Now().Location())
	start := end.Add(-29 * 24 * time.Hour)
	s := start.Format("2006-01-02")
	e := end.Format("2006-01-02")
	commomDao := dao.GetCommonDAO()

	// 前30日每日活跃账号图表
	//daysActiveUser := commomDao.Get30DaysActiveUser(Before30Unix)
	// 前30日每日活跃群组图表
	//daysActiveChat := commomDao.Get30DaysActiveChat(Before30Unix)

	// 前30日每日活跃图表
	daysActiveUser, daysActiveChat := commomDao.Get30DaysActiveData(s,e)

	data := map[string]interface{}{
		"day_active_user": daysActiveUser,
		"day_active_chat": daysActiveChat,
	}
	middleware.ResponseSuccess(c, data)
}

// 前100人数群消息
func (service *DashboardController) MaxMemberChat(c *gin.Context) {
	end := time.Now().Unix()

	y, m, d := time.Now().Date()
	time1 := time.Date(y, m, d, 0, 0, 0, 0, time.FixedZone("CST", 3600*8))
	Before7Unix := time1.Add(-7 * 24 * time.Hour).Unix()

	commomDao := dao.GetCommonDAO()

	// 获取前100人数群id
	chatIds := commomDao.Get100ChatIds()

	// 前7天活跃的群
	chatIds7 := commomDao.GetActiveChatIds(Before7Unix, end)
	activeMap := make(map[int32]bool, len(chatIds7))
	for _,id := range chatIds7 {
		activeMap[id] = true
	}

	chats := chatapi.ChatInfo(chatIds)

	for _, c := range chats {
		chatId, _ := c["chat_id"].(int32)
		c["is_active"] = activeMap[chatId]
	}

	data := map[string]interface{}{
		"100_chat": chats,
	}
	middleware.ResponseSuccess(c, data)
}


