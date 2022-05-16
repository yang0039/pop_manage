package store_manage

import "github.com/gin-gonic/gin"

type StoreController struct{}

func StoreRegister(group *gin.RouterGroup) {
	service := &StoreController{}

	group.GET("/all_store", service.AllStore)
	group.GET("/user_store", service.UserStore)
	group.GET("/last_upload", service.LastUpload)
}
