package system_manage

import "github.com/gin-gonic/gin"

type SystemController struct{}

func SystemRegister(group *gin.RouterGroup) {
	service := &SystemController{}
	group.GET("/query_all_permission", service.GetAllPermissions)
	group.POST("/add_role", service.AddRole)
	group.POST("/edit_role_permission", service.EditRolePermission)
	group.POST("/delete_role", service.DelRole)
	group.GET("/query_all_role", service.GetAllRolePermission)

	// 添加账户
	group.POST("/add_account", service.AddAcount)
	// 分配角色
	group.POST("/edit_account", service.EditAcount)
	// 禁用账户
	group.POST("/edit_account_state", service.EditAcountState)
	// 修改密码
	group.POST("/edit_account_pwd", service.EditAcountPwd)
	// 获取角色对应的页面
	group.GET("/query_account", service.GetAccount)
	group.GET("/query_all_account", service.GetAllAccount)
	group.GET("/query_login_log", service.GetAccountLogin)
	// 删除账户
	group.POST("/del_account", service.DelAcount)

	group.POST("/add_white_list", service.AddAllowIp)
	group.POST("/del_white_list", service.DelAllowIp)
	group.GET("/query_white_list", service.GetAllowIp)

	group.POST("/edit_config", service.EditConfig)
	group.GET("/get_config", service.GetConfig)

	// 标签相关
	group.POST("/add_label", service.AddLabel)
	group.POST("/del_label", service.DelLabel)
	group.POST("/update_label", service.UpdateLabel)
	group.GET("/get_label", service.GetLabel)
	group.GET("/get_label_note", service.GetLabbelNote)

	// 操作记录
	group.GET("/get_request_record", service.GetRequestRecord)
}

