package public

import (
	"fmt"
	"pop-api/dal/dao"
)

type PermissionUrl struct {
	PermissionId int32
	Url string
	Method string
}

// 账户对应的角色
var AcountRole map[int32][]int32
// 角色对应的权限
var RolePermission map[int32][]int32
// 权限对应的接口
var PermisUrl map[int32][]PermissionUrl

// 超级管理员id
var SuperAdminId int32


func InitData(){
	accountDao := dao.GetAccountDAO()
	roleDao := dao.GetRoleDAO()
	urlDao := dao.GetPermissionsUrlDAO()
	AcountRole, _ = accountDao.GetAccountRoleName()
	RolePermission = roleDao.GetAllRolePermission()
	perUrls := urlDao.GetAllPermissionsUrl()
	PermisUrl = make(map[int32][]PermissionUrl, 0)
	for _, u := range perUrls {
		if !u.IsEffect {
			continue
		}
		_, ok := PermisUrl[u.PermissionsId]
		if !ok {
			PermisUrl[u.PermissionsId] = make([]PermissionUrl, 0)
		}
		pu := PermissionUrl{
			PermissionId: u.PermissionsId,
			Url: u.Url,
			Method: u.Method,
		}
		PermisUrl[u.PermissionsId] = append(PermisUrl[u.PermissionsId], pu)
	}

	SuperAdminId = accountDao.QryAdminId()
	fmt.Println("---------------------------")
	fmt.Println("  SuperAdminId：", SuperAdminId)
	fmt.Println("---------------------------")
	for k,v := range AcountRole {
		fmt.Print("账户：", k)
		fmt.Println("  角色：", v)
	}
	fmt.Println("---------------------------")
	for k,v := range RolePermission {
		fmt.Print("角色：", k)
		fmt.Println("  权限：", v)
	}
	fmt.Println("---------------------------")
	for k,v := range PermisUrl {
		fmt.Print("权限：", k)
		fmt.Println("  url：", v)
	}
	fmt.Println("---------------------------")
}