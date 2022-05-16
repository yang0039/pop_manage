package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ip2location/ip2location-go"
	"go.uber.org/zap"
	"pop-api/baselib/logger"
	"pop-api/baselib/minio_client"
	"pop-api/baselib/mysql_client"
	"pop-api/baselib/redis_client"
	"pop-api/baselib/util"
	_ "pop-api/controller/_dummy/impl"
	"pop-api/controller/auth"
	"pop-api/controller/chat_manage"
	"pop-api/controller/dashboard"
	"pop-api/controller/data_cache"
	"pop-api/controller/member_manage"
	"pop-api/controller/message_record"
	"pop-api/controller/store_manage"
	"pop-api/controller/system_manage"
	"pop-api/dal/dao"
	"pop-api/middleware"
	"pop-api/public"
)

var (
	Log   *zap.Logger
	Sugar *zap.SugaredLogger
)

func init() {
	Log, Sugar = logger.InitLogger("debug", 9003)
}

func Default() *gin.Engine {
	//debugPrintWARNINGDefault()
	engine := gin.New()
	engine.Use(gin.Recovery())
	return engine
}

func main() {
	var err error
	if err = InitConfig(); err != nil {
		panic(err)
	}
	fmt.Println("db_index=", Conf.Dbindex)
	Sugar.Infof("config=%v", Conf)
	dao.InstallMysqlDAOManager(mysql_client.NewSqlxDB(Conf.Mysql), Conf.Dbindex)
	redis_client.InstallRedisClientManager(*Conf.Redis)

	minio_client.InitData(Conf.Minio)
	// 先写固定值，后续从配置文件获取
	//minio_client.InitClientConfig("127.0.0.1:9123", "127.0.0.1:9000", "Wink@YaMyB2GmOEetkib6O#+KRfuze6T", "DQMYMM5HIJ4EF2XROGRK", "UooDmD1HwHvv67fjuVHYFpQcMGmyUCjyJt+B+n24")

	//minio_client.PresignedGetObject("photo", "0/20220118/1483260720857632768.jpg")
	//fmt.Println("url=", url)

	// 定时任务
	timeTask()

	gin.SetMode("release")
	router := Default()
	router.Use(middleware.Cors(), middleware.RequestLog())
	//router.Group("/api")

	util.InitSnowFlakeId(1, 1)
	ip2location.Open("./IP-COUNTRY-REGION-CITY-LATITUDE-LONGITUDE-ISP.BIN")

	// initdata
	public.InitData()

	// dashboadr
	dashboardRouter := router.Group("/dashboard")
	dashboardRouter.Use(
		middleware.JwtVerify(), middleware.OperaRecord(),
	)

	// chat
	chatRouter := router.Group("/chat_manage")
	chatRouter.Use(
		middleware.JwtVerify(), middleware.OperaRecord(),
	)

	// user
	userRouter := router.Group("/user_manage")
	userRouter.Use(
		middleware.JwtVerify(), middleware.OperaRecord(),
	)

	// auth
	authRouter := router.Group("/auth")
	authRouter.Use(
		middleware.OperaRecord(),
	)

	// system
	systemRouter := router.Group("/system")
	systemRouter.Use(
		middleware.JwtVerify(), middleware.OperaRecord(),
	)

	// chat_record
	recordRouter := router.Group("/record")
	recordRouter.Use(
		middleware.JwtVerify(), middleware.OperaRecord(),
	)

	// chat_record
	StoreRouter := router.Group("/store_manage")
	StoreRouter.Use(
		middleware.JwtVerify(), middleware.OperaRecord(),
	)

	//Sugar.Infof("md5=%s", md5V("Kyh@51814"))

	{
		dashboard.DashboardRegister(dashboardRouter)
		chat_manage.ChatManageRegister(chatRouter)
		member_manage.ChatUserRegister(userRouter)
		auth.AuthRegister(authRouter)
		system_manage.SystemRegister(systemRouter)
		message_record.ChatRecordRegister(recordRouter)
		store_manage.StoreRegister(StoreRouter)
	}

	Sugar.Infof("start api server")
	router.Run(":8080")
	//router.RunTLS(":8080", "ssl.pem", "ssl.key")
}

//func md5V(str string) string {
//	h := md5.New()
//	h.Write([]byte(str))
//	return hex.EncodeToString(h.Sum(nil))
//}

func timeTask() {
	go data_cache.DailyActiveData()
}