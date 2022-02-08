package system_manage

import (
	"github.com/gin-gonic/gin"
	"pop-api/dal/dao"
	"pop-api/middleware"
)

type Permission struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

type MenuFunc struct {
	MenuId      int32        `json:"menu_id"`
	FuncName    string       `json:"func_name"`
	FuncTitle   string       `json:"func_title"`
	Permissions []Permission `json:"permissions"`
}

func (service *SystemController) GetAllPermissions(c *gin.Context) {
	menuDao := dao.GetMenuDAO()
	perDao := dao.GetPermissionsDAO()
	//perUrlDao := dao.GetPermissionsUrlDAO()

	menus := menuDao.GetAllMenu()
	menuFunc := perDao.GetAllMenuFunc()
	permissions := perDao.GetAllPermissions()


	perMap := make(map[string][]Permission, 0)
	funcMap := make(map[int32][]*MenuFunc, 0)

	funcMenuMap := make(map[string]*MenuFunc, 0)

	for _, mf := range menuFunc {
		f := &MenuFunc{
			MenuId: mf.MenuId,
			FuncName: mf.FuncName,
			FuncTitle: mf.FuncTitle,
		}
		funcMenuMap[mf.FuncName] = f
		_, ok := funcMap[mf.MenuId]
		if !ok {
			funcMap[mf.MenuId] = make([]*MenuFunc, 0)
		}
		funcMap[mf.MenuId] = append(funcMap[mf.MenuId], f)
	}

	for _, p := range permissions {
		_, ok :=perMap[p.FuncName]
		if !ok {
			perMap[p.FuncName] = make([]Permission, 0)
		}
		per := Permission{
			Id: p.Id,
			Name: p.Name,
			Title: p.Title,
		}
		perMap[p.FuncName] = append(perMap[p.FuncName], per)

	}

	for name, f := range funcMenuMap {
		funcMenuMap[name].Permissions = perMap[f.FuncName]
	}

	res := make([]map[string]interface{}, 0)
	for _, m := range menus {
		menu := map[string]interface{}{
			"id": m.Id,
			"name": m.Name,
			"title": m.Title,
			"menu_func": funcMap[m.Id],
		}
		res = append(res, menu)
	}
	middleware.ResponseSuccess(c, res)
}
