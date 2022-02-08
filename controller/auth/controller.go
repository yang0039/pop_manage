package auth

import "github.com/gin-gonic/gin"

type AccountController struct{}

func AuthRegister(group *gin.RouterGroup) {
	service := &AccountController{}
	group.POST("/login", service.Login)
}

