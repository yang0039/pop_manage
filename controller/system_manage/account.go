package system_manage

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"pop-api/baselib/redis_client"
	"pop-api/dal/dao"
	"pop-api/dto"
	"pop-api/middleware"
	"pop-api/public"
	"strconv"
	"time"
)

func (service *SystemController) AddAcount(c *gin.Context) {
	params := &dto.Acount{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.AccountName == "" || params.UserName == "" || params.Pwd == "" {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	accountDao := dao.GetAccountDAO()
	roleDAO := dao.GetRoleDAO()

	userName, _ := c.Get("user_name")
	name, _ := userName.(string)

	if accountDao.AccountIsExit(params.AccountName) {
		middleware.ResponseError(c, 400, "账户名已存在", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	var util int32
	if params.PwdUtil != 0 {
		util = int32(time.Now().AddDate(0, 0, int(params.PwdUtil)).Unix())

	}
	effectRoleIds := roleDAO.GetEffectRoleIds(params.RoleIds)
	pwd := md5V(params.Pwd)
	id := accountDao.AddAccount(params.AccountName, params.UserName, pwd, name, util, effectRoleIds)
	res := map[string]int32{
		"account_id": id,
	}

	if len(effectRoleIds) > 0 {
		public.AcountRole[id] = effectRoleIds
	}

	middleware.ResponseSuccess(c, res)
}

func (service *SystemController) EditAcount(c *gin.Context) {
	params := &dto.Acount{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Pwd == "" || params.Id == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	accId, _ := c.Get("account_id")
	id, _ := accId.(int32)

	accountDao := dao.GetAccountDAO()
	roleDAO := dao.GetRoleDAO()

	// 校验密码
	pwdRes := VerifyPwd(id, params.Pwd)
	if pwdRes == 1 {
		middleware.ResponseError(c, 400, "密码错误", errors.New("密码错误"))
		return
	} else if pwdRes == 2 {
		middleware.ResponseError(c, 400, "密码已过期", errors.New("密码已过期"))
		return
	}
	effectRoleIds := roleDAO.GetEffectRoleIds(params.RoleIds)
	if len(effectRoleIds) > 0 {
		accountDao.EditAccountRole(params.Id, effectRoleIds)
	}
	if len(params.NewPwd) > 0 {
		newPwd := md5V(params.NewPwd)
		accountDao.EditAccountPwd(params.Id, newPwd)
	}
	var util int32
	if params.PwdUtil > 0 {
		util = int32(time.Now().AddDate(0, 0, int(params.PwdUtil)).Unix())
	}
	accountDao.EditAccountPwdUtil(params.Id, util)

	if len(effectRoleIds) > 0 {
		public.AcountRole[params.Id] = effectRoleIds
	}

	res := map[string]int32{
		"account_id": params.Id,
	}
	middleware.ResponseSuccess(c, res)
}

// 更新账户状态
func (service *SystemController) EditAcountState(c *gin.Context) {
	params := &dto.Acount{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Id == 0 || (params.ForbiddenType != 0 && params.ForbiddenType != 1) {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	accountDao := dao.GetAccountDAO()
	if !accountDao.AccountIsExitById2(params.Id) {
		middleware.ResponseError(c, 400, "账户不存在", errors.New("账户不存在"))
		return
	}

	accountDao.EditAccountState(params.Id, params.ForbiddenType)

	if params.ForbiddenType == 0 {    // 禁用
		delete(public.AcountRole, params.Id)
	} else {      // 解除禁用
		roles := accountDao.GetAccountRole(params.Id)
		roleIds := make([]int32, 0, len(roles))
		for _, r := range roles {
			rId,_ := r["id"].(int32)
			if rId != 0 {
				roleIds = append(roleIds, rId)
			}
		}
		public.AcountRole[params.Id] = roleIds
	}

	res := map[string]int32{
		"account_id": params.Id,
	}
	middleware.ResponseSuccess(c, res)
}

// 删除账号
func (service *SystemController) DelAcount(c *gin.Context) {
	params := &dto.Acount{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Id == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	accountDao := dao.GetAccountDAO()

	accId, _ := c.Get("account_id")
	id, _ := accId.(int32)

	// 校验密码
	pwdRes := VerifyPwd(id, params.Pwd)
	if pwdRes == 1 {
		middleware.ResponseError(c, 400, "密码错误", errors.New("密码错误"))
		return
	} else if pwdRes == 2 {
		middleware.ResponseError(c, 400, "密码已过期", errors.New("密码已过期"))
		return
	}


	if !accountDao.AccountIsExitById(params.Id) {
		middleware.ResponseError(c, 400, "账户不存在", errors.New("账户不存在"))
		return
	}

	accountDao.DelAccount(params.Id)

	delete(public.AcountRole, params.Id)

	res := map[string]int32{
		"account_id": params.Id,
	}
	middleware.ResponseSuccess(c, res)
}


// 更新密码(用户自己修改)
func (service *SystemController) EditAcountPwd(c *gin.Context) {
	params := &dto.Acount{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Pwd == "" || params.NewPwd == "" {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}

	accId, _ := c.Get("account_id")
	id, _ := accId.(int32)

	accountDao := dao.GetAccountDAO()

	pwdRes := VerifyPwd(id, params.Pwd)
	if pwdRes == 1 {
		middleware.ResponseError(c, 400, "密码错误", errors.New("密码错误"))
		return
	} else if pwdRes == 2 {
		middleware.ResponseError(c, 400, "密码已过期", errors.New("密码已过期"))
		return
	}
	newPwd := md5V(params.NewPwd)
	//util := time.Now().AddDate(0, 0, int(params.PwdUtil)).Unix()
	accountDao.EditAccountPwd(id, newPwd)
	res := map[string]int32{
		"account_id": params.Id,
	}
	middleware.ResponseSuccess(c, res)

}

func (service *SystemController) GetAccount(c *gin.Context) {
	accId, _ := c.Get("account_id")
	id, _ := accId.(int32)

	accountDao := dao.GetAccountDAO()
	roleDAO := dao.GetRoleDAO()
	menuDao := dao.GetMenuDAO()
	permisDao := dao.GetPermissionsDAO()

	acc := accountDao.GetAccountById(id)
	if acc.Id == 0 {
		middleware.ResponseError(c, 400, "账户不存在", errors.New("账户不存在"))
		return
	}

	// 获取角色对应的权限id
	roleIds := accountDao.GetAccountRoleIds(id)
	permisIds := roleDAO.GetRolePermissionIds(roleIds)

	roles := accountDao.GetAccountRole(id)
	permissions := permisDao.GetPermissionsByIds(permisIds)
	menuMap := menuDao.GetMenuMap()

	funcMap := make(map[string]map[string]interface{})
	funcPermis := make(map[string][]map[string]interface{})

	// 菜单对应的功能
	menuFunc := make(map[int32][]string)
	menuIds := make([]int32, 0)

	for _, p := range permissions {
		_, ok := funcPermis[p.FuncName]
		if !ok {
			funcPermis[p.FuncName] = make([]map[string]interface{}, 0)
		}
		m := map[string]interface{}{
			"id": p.Id,
			"name": p.Name,
			"title": p.Title,
		}
		funcPermis[p.FuncName] = append(funcPermis[p.FuncName], m)

		_, ok2 := funcMap[p.FuncName]
		if !ok2 {
			m2 := map[string]interface{}{
				"menu_id": p.MenuId,
				"func_name": p.FuncName,
				"func_title": p.FuncTitle,
			}
			funcMap[p.FuncName] = m2
		}

		_, ok3 := menuFunc[p.MenuId]
		if !ok3 {
			menuFunc[p.MenuId] = make([]string, 0)
			menuIds = append(menuIds, p.MenuId)
		}
		var has bool
		for _, v := range menuFunc[p.MenuId] {
			if v == p.FuncName {
				has = true
				break
			}
		}
		if !has {
			menuFunc[p.MenuId] = append(menuFunc[p.MenuId], p.FuncName)
		}
		//menuFunc[p.MenuId]
	}

	for funcName,_ :=  range funcMap {
		funcMap[funcName]["permissions"] = funcPermis[funcName]
	}

	// 获取账户的菜单
	menuRes := make([]map[string]interface{}, 0)
	for _, id := range menuIds {
		m := make(map[string]interface{})
		m["id"] = id
		m["title"] = menuMap[id].Title
		m["name"] = menuMap[id].Name

		menufuns := make([]interface{}, 0)
		for _, funcName := range menuFunc[id] {
			menufuns = append(menufuns, funcMap[funcName])
		}
		m["menu_func"] = menufuns

		menuRes = append(menuRes, m)
	}


	//pageName := commonDao.GetPageName(permisIds)
	res := map[string]interface{}{
		"id":           acc.Id,
		"account_name": acc.AccountName,
		"user_name":    acc.UserName,
		"pwd_util":     acc.PwdUtil,
		"add_time":     acc.AddTime,
		"pages":        menuRes,
		"roles":        roles,
	}
	middleware.ResponseSuccess(c, res)
}

// 获取全部账户
func (service *SystemController) GetAllAccount(c *gin.Context) {
	accountDao := dao.GetAccountDAO()
	loginDao := dao.GetLoginLogDAO()
	accounts := accountDao.GetAllAccount()
	roleIds, roleNames := accountDao.GetAccountRoleName()
	loginLog := loginDao.GetLastLog()

	res := make([]map[string]interface{}, 0, len(accounts))
	for _, a := range accounts {
		lastTime, _ := redis_client.RedisCache.HGet("manage:last:opera", strconv.Itoa(int(a.Id))).Int64()
		m := make(map[string]interface{})
		m["id"] = a.Id
		m["account_name"] = a.AccountName
		m["user_name"] = a.UserName
		m["pwd_util"] = a.PwdUtil
		m["is_effect"] = a.IsEffect
		m["add_time"] = a.AddTime
		m["role_ids"] = roleIds[a.Id]
		m["role_name"] = roleNames[a.Id]
		m["last_login_time"] = loginLog[a.Id].LoginTime
		m["last_ip"] = loginLog[a.Id].LoginIp
		m["last_time"] = lastTime
		res = append(res, m)
	}

	middleware.ResponseSuccess(c, res)
}

// 获取登录日志
func (service *SystemController) GetAccountLogin(c *gin.Context) {
	params := &dto.Acount{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Id == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	if params.Limit == 0 {
		params.Limit = 20
	}

	loginDao := dao.GetLoginLogDAO()

	logs, count := loginDao.GetLoginLog(params.Id, params.Limit, params.Offset)

	data := map[string]interface{}{
		"login_log": logs,
		"count":     count,
	}
	middleware.ResponseSuccess(c, data)

}

// 0:通过, 1:密码错误, 2:密码已过期
func VerifyPwd(accId int32, verPwd string) int32 {
	now := time.Now().Unix()
	accountDao := dao.GetAccountDAO()
	pwd, pwdUtil := accountDao.QryAccountPwd(accId)
	if pwd != md5V(verPwd) {
		return 1
	}
	if pwdUtil != 0 {
		if now > int64(pwdUtil) {
			return 2
		}
	}
	return 0
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
