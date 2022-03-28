package member_manage


import "github.com/gin-gonic/gin"

type UserController struct{}

func ChatUserRegister(group *gin.RouterGroup) {
	service := &UserController{}
	group.POST("/update_note", service.AddUserNote)
	group.GET("/query_user", service.GetUserInfo)
	group.GET("/query_user_chat", service.GetChatByUser)
	group.POST("/delete_user", service.DelAccount)
	group.POST("/banned_user", service.BannedUser)
	group.GET("/query_phone_transaction", service.QryUserPhone)
	group.GET("/query_user_banned", service.GetUserBanned)
	group.GET("/query_user_report", service.GetReport)
	group.GET("/query_file", service.GetFile)
	group.GET("/query_user_contact", service.UserContact)
	group.GET("/query_user_dialogs", service.UserDialogs)
	group.GET("/query_user_relation", service.UserRelation)
	group.POST("/set_official_user", service.SetOfficialUser)
}

