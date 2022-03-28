package data_cache

import (
	"pop-api/baselib/logger"
	"pop-api/dal/dao"
	"runtime/debug"
	"time"
)

// 缓存每日活跃数据，每天00:00:01执行
func DailyActiveData() {
	defer func() {
		if err := recover(); err != nil {
			logger.LogSugar.Errorf("recover error DailyActiveData panic: %v\n%s", err, string(debug.Stack()))
		}
	}()

	now := time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 10, 0, now.Location()) //获取下一个凌晨的日期
	loopTimer := time.NewTimer(next.Sub(now))                                          //计算当前时间到凌晨的时间间隔，设置一个定时器
	//fmt.Println("==== next.Sub(now)=", next.Sub(now))
	defer loopTimer.Stop()

	for {
		select {
		case <-loopTimer.C:
			AddActiveData()
		}
		//loopTimer.Reset(time.Second * 5)
		loopTimer.Reset(time.Hour * 24)
	}
}

func AddActiveData() {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	end := t.Unix() - 8 *3600
	start := end - 3600*24
	logger.LogSugar.Infof("AddActiveData start:%d, end:%d", start, end)

	commomDao := dao.GetCommonDAO()
	uCount := commomDao.GetDayActiveUser(start, end)
	cCount := commomDao.GetDayActiveChat(start, end)
	commomDao.AddActieCache( uCount, cCount, t.Add(-16 * time.Hour).Format("2006-01-02"))
}
