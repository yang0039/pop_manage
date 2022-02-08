package dashboard

import "github.com/gin-gonic/gin"

type DashboardController struct{}

func DashboardRegister(group *gin.RouterGroup) {
	service := &DashboardController{}
	group.GET("/total_data", service.TotalData)
	group.GET("/active_data", service.ActiveData)
	group.GET("/max_member_chat", service.MaxMemberChat)
}
