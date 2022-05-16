package dashboard

import (
	"fmt"
	"pop-api/baselib/logger"
	"github.com/gin-gonic/gin"
	"pop-api/baselib/util"
	"pop-api/dal/dao"
	"pop-api/middleware"
	"syscall"
	"time"
)

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

// disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func (service *DashboardController) TotalData(c *gin.Context) {

	start := int64(0)
	end := time.Now().Unix()

	y, m, d := time.Now().Date()
	time1 := time.Date(y, m, d, 0, 0, 0, 0, time.FixedZone("CST", 3600*8))
	//Before5Unix := time1.Add(-5 * 24 * time.Hour).Unix()
	//Before7Unix := time1.Add(-7 * 24 * time.Hour).Unix()
	//Before30Unix := time1.Add(-30 * 24 * time.Hour).Unix()
	date5 := time1.Add(-24 * time.Hour).Format("2006-01-02")

	userDao := dao.GetUserDAO()
	chatDao := dao.GetChatDAO()
	//callDao := dao.GetCallDAO()
	commomDao := dao.GetCommonDAO()

	logger.LogSugar.Info("TotalData 1")
	allNum := userDao.GetUserNum(0)
	logger.LogSugar.Info("TotalData 2")
	todayNum := userDao.GetUserNum(util.GetTodayUnix())
	logger.LogSugar.Info("TotalData 3")
	// 获取指定时间创建的3人以上的群
	memChatNum := commomDao.GetMemberChatNum(start, end, 3)
	logger.LogSugar.Info("TotalData 4")
	chatNum := chatDao.GetChatNum(start, end)
	logger.LogSugar.Info("TotalData 5")
	msgCount,callNum := commomDao.GetMsgPhoneCount()
	// 活跃账号数
	// todo 先给0，这个非常耗时
	activeCountALl := 0
	//activeCountALl := commomDao.GetActiveUserCount(start, end)

	// 发送消息的数量
	//msgCount := commomDao.GetSendMsgCount()
	logger.LogSugar.Info("TotalData 6")
	// 语音通话发起数
	//callNum := callDao.GetCallNum(start, end)
	logger.LogSugar.Info("TotalData 7")
	// 5日内活跃账号数
	//activeCount5 := commomDao.GetActiveUserCount(Before5Unix, end)
	// 5日内活跃群组数
	//chatIds5 := commomDao.GetActiveChatIds(Before5Unix, end)
	uActiveCount5, cActiveCount5 := commomDao.GetActiveData5(date5)
	logger.LogSugar.Info("TotalData 8")
	disk := DiskUsage("/data")
	fmt.Printf("All: %.2f GB", float64(disk.All)/float64(GB))
	fmt.Printf("Used: %.2f GB", float64(disk.Used)/float64(GB))
	fmt.Printf("Free: %.2f GB", float64(disk.Free)/float64(GB))

	data := map[string]interface{}{
		"total_user_num":        allNum,
		"new_register_num":      todayNum,
		"three_member_chat_num": memChatNum,
		"total_chat_num":        chatNum,
		"active_num":            activeCountALl,
		"msg_count":             msgCount,
		"call_num":              callNum,

		"five_active_num":      uActiveCount5,
		"five_active_chat_num": cActiveCount5,

		"all_store":  util.Folat2(float64(disk.All) / float64(GB)),
		"used_store": util.Folat2(float64(disk.Used) / float64(GB)),
		"free_store": util.Folat2(float64(disk.Free) / float64(GB)),
	}
	middleware.ResponseSuccess(c, data)
}
