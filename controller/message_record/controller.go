package message_record

import "github.com/gin-gonic/gin"

type RecordController struct{}

func ChatRecordRegister(group *gin.RouterGroup) {
	service := &RecordController{}
	//group.POST("/update_note", service.AddUserNote)
	group.GET("/query_message", service.GetMessageRecord)
	group.GET("/query_dialogs", service.GetDialogs)

	group.POST("/del_file", service.DelFile)
	group.POST("/del_peer_file", service.DelFileMessages)
}
