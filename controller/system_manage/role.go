package system_manage

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"pop-api/dal/dao"
	"pop-api/dto"
	"pop-api/middleware"
	"pop-api/public"
)

// 增加角色
func (service *SystemController) AddRole(c *gin.Context) {
	params := &dto.Role{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Name == "" {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	permissionDao := dao.GetPermissionsDAO()
	roleDao := dao.GetRoleDAO()

	userName, _ := c.Get("user_name")
	name, _ := userName.(string)

	// 过滤权限id
	effectIds := permissionDao.GetEffectPermissionIds(params.PermissionIds)

	// 添加角色
	roleId := roleDao.AddRole(params.Name, name, effectIds)

	if len(effectIds) > 0 {
		public.RolePermission[int32(roleId)] = effectIds
	}

	res := map[string]int32{
		"role_id": int32(roleId),
	}
	middleware.ResponseSuccess(c, res)
}

// 编辑角色权限
func (service *SystemController) EditRolePermission(c *gin.Context) {
	params := &dto.Role{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Id == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}

	permissionDao := dao.GetPermissionsDAO()
	roleDao := dao.GetRoleDAO()

	// 查找角色是否有效
	isEffect := roleDao.RoleIsEffect(params.Id)
	if !isEffect {
		middleware.ResponseError(c, 400, "角色不存在", errors.New(fmt.Sprintf("role_id:%d", params.Id)))
		return
	}

	// 过滤权限id
	effectIds := permissionDao.GetEffectPermissionIds(params.PermissionIds)

	// 编辑角色
	roleDao.EditRole(params.Id, effectIds)

	if len(effectIds) > 0 {
		public.RolePermission[params.Id] = effectIds
	}

	res := map[string]int32{
		"role_id": params.Id,
	}
	middleware.ResponseSuccess(c, res)
}

// 删除角色
func (service *SystemController) DelRole(c *gin.Context) {
	params := &dto.Role{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Id == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}

	roleDao := dao.GetRoleDAO()
	roleDao.DeleteRole(params.Id)
	delete(public.RolePermission, params.Id)
	res := map[string]int32{
		"role_id": params.Id,
	}
	middleware.ResponseSuccess(c, res)
}

// 获取全部角色权限
func (service *SystemController) GetAllRolePermission(c *gin.Context) {
	roleDao := dao.GetRoleDAO()

	roles := roleDao.GetAllRole()
	perMap := roleDao.GetAllRolePermission()

	res := make([]map[string]interface{}, 0)
	for _, r := range roles {
		rId,_ := r["role_id"].(int32)
		m := map[string]interface{}{
			"role_id": r["role_id"],
			"name": r["name"],
			"permission_ids": perMap[rId],
		}
		res = append(res, m)
	}
	middleware.ResponseSuccess(c, res)
}