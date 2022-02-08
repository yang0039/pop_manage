package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"pop-api/baselib/logger"
	"pop-api/baselib/redis_client"
	"pop-api/public"
	"strings"
	"time"
)

// 所有用户都能通过的接口
var AllowUrl = []string{
	"/dashboard/total_data", "/dashboard/active_data",
	"/dashboard/max_member_chat", "/system/query_account",
	"/system/edit_account_pwd", "/system/query_all_permission",
	"/system/get_label",
}

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.LogSugar.Infof("method:%s, uri:%s", c.Request.Method, c.Request.RequestURI)
		c.Next()
	}
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization,X-Requested-With,Accept,Origin,token")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}


func JwtVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("token")
		if tokenStr == "" {
			ResponseError(c, 400, "缺少token", errors.New("缺少token"))
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("popchat_jwt"), nil
		})

		if err != nil {
			ResponseError(c, 400, "无效的token", err)
			c.Abort()
			return
		}

		if !token.Valid {
			ResponseError(c, 400, "无效的token", errors.New("无效的token"))
			c.Abort()
			return
		}
		claims, ok := token.Claims.(*MyCustomClaims)

		if !ok {
			ResponseError(c, 400, "无效的token", errors.New("无效的token"))
			return
		}
		//将uid写入请求参数
		id := claims.Id
		name := claims.Name
		c.Set("account_id", id)
		c.Set("user_name", name)
		// 权限验证
		m := strings.ToLower(c.Request.Method)
		if !verifyPermission(id, c.FullPath(), m) {
			ResponseError(c, 400, "您没有该权限", errors.New("您没有该权限"))
			return
		}

		// 记录时间
		now := time.Now().Unix()
		redis_client.RedisCache.HSet("manage:last:opera", id, now)

		c.Next()
	}
}

func verifyPermission(id int32, url, method string) bool {
	logger.LogSugar.Infof("verifyPermission  account_id:%d, method:%s, uri:%s", id, method, url)

	// 超级管理员不需要鉴定权限
	if id == public.SuperAdminId {
		return true
	}

	for _, u := range AllowUrl {
		if u == url {
			return true
		}
	}

	roleIds := public.AcountRole[id]
	if len(roleIds) == 0 {
		return false
	}
	allPerIds := make([]int32, 0)
	for _, rId := range roleIds {
		allPerIds = append(allPerIds, public.RolePermission[rId]...)
	}

	for _, pId := range allPerIds {
		pUrls := public.PermisUrl[pId]
		for _, pu := range pUrls {
			if url == pu.Url && method == pu.Method {
				return true
			}
		}
	}
	return false
}

