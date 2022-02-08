package chat_manage

import "github.com/gin-gonic/gin"

type ChatController struct{}

func ChatManageRegister(group *gin.RouterGroup) {
	service := &ChatController{}
	group.POST("/update_note", service.AddChatNote)
	group.GET("/query_chat", service.GetChatInfo)
	group.GET("/query_chat_member", service.GetChatMembber)
	group.GET("/query_chat_msg", service.GetChatMessage)
	group.POST("/update_chat_status", service.AddChatStatus)
	group.GET("/qry_status_record", service.QryChatStatusRecord)

	//group.GET("/active_data", service.ActiveData)
	//group.GET("/max_member_chat", service.MaxMemberChat)
}
