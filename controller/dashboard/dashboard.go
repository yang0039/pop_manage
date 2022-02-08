package dashboard

import (
	"github.com/gin-gonic/gin"
	"pop-api/baselib/util"
	"pop-api/dal/dao"
	"pop-api/middleware"
	"time"
)

func (service *DashboardController) TotalData(c *gin.Context) {

	start := int64(0)
	end := time.Now().Unix()

	y, m, d := time.Now().Date()
	time1 := time.Date(y, m, d, 0, 0, 0, 0, time.FixedZone("CST", 3600*8))
	Before5Unix := time1.Add(-5 * 24 * time.Hour).Unix()
	//Before7Unix := time1.Add(-7 * 24 * time.Hour).Unix()
	//Before30Unix := time1.Add(-30 * 24 * time.Hour).Unix()

	userDao := dao.GetUserDAO()
	chatDao := dao.GetChatDAO()
	callDao := dao.GetCallDAO()
	commomDao := dao.GetCommonDAO()

	allNum := userDao.GetUserNum(0)

	todayNum := userDao.GetUserNum(util.GetTodayUnix())

	// 获取指定时间创建的3人以上的群
	memChatNum := commomDao.GetMemberChatNum(start, end, 3)

	chatNum := chatDao.GetChatNum(start, end)

	// 活跃账号数
	activeCountALl := commomDao.GetActiveUserCount(start, end)

	// 发送消息的数量
	msgCount := commomDao.GetSendMsgCount(start, end)

	// 语音通话发起数
	callNum := callDao.GetCallNum(start, end)

	// 5日内活跃账号数
	activeCount5 := commomDao.GetActiveUserCount(Before5Unix, end)

	// 5日内活跃群组数
	chatIds5 := commomDao.GetActiveChatIds(Before5Unix, end)

	data := map[string]interface{}{
		"total_user_num":        allNum,
		"new_register_num":      todayNum,
		"three_member_chat_num": memChatNum,
		"total_chat_num":        chatNum,
		"active_num":            activeCountALl,
		"msg_count":             msgCount,
		"call_num":              callNum,

		"five_active_num":      activeCount5,
		"five_active_chat_num": len(chatIds5),
		//"100_chat":             chats,
		//"day_active_user":      daysActiveUser,
		//"day_active_chat":      daysActiveChat,
	}
	middleware.ResponseSuccess(c, data)
}
