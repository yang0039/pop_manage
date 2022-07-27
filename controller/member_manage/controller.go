package member_manage


import "github.com/gin-gonic/gin"

type UserController struct{}

func ChatUserRegister(group *gin.RouterGroup) {
	service := &UserController{}
	group.POST("/update_note", service.AddUserNote)
	group.GET("/query_user", service.GetUserInfo)
	group.GET("/query_user_chat", service.GetChatByUser)
	group.POST("/delete_user", service.DelUser)
	group.POST("/banned_user", service.BannedUser)
	group.GET("/query_phone_transaction", service.QryUserPhone)
	group.GET("/query_user_banned", service.GetUserBanned)
	group.GET("/query_user_report", service.GetReport)
	//group.GET("/query_file", service.GetFile)
	group.GET("/query_user_contact", service.UserContact)
	group.GET("/query_user_dialogs", service.UserDialogs)
	group.GET("/query_user_relation", service.UserRelation)
	group.POST("/set_official_user", service.SetOfficialUser)

	group.GET("/query_user_store", service.UserStore)
	group.GET("/query_user_file", service.UserFile)
	group.GET("/query_user_login", service.QueryUserLogin)

	group.POST("/update_user_username", service.UpdateUserName)
	group.POST("/update_user_phone", service.UpdateUserPhone)
	group.GET("/query_user_opera", service.GetUserOpera)
	group.POST("/delete_user_pwd", service.DelUserPassword)
}

