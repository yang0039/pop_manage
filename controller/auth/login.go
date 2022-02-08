package auth

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"pop-api/dal/dao"
	"pop-api/dto"
	"pop-api/middleware"
	"strings"
	"time"
)

func (service *AccountController) Login(c *gin.Context) {
	account := &dto.Account{}
	if err := c.ShouldBind(account); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if account.Name == "" || account.Pwd == "" {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", account)))
		return
	}

	accountDao := dao.GetAccountDAO()
	loginDao := dao.GetLoginLogDAO()
	allowIpDao := dao.GetAllowIpDAO()
	commomDao := dao.GetCommonDAO()

	var ip string
	ips := strings.Split(c.Request.RemoteAddr, ":")
	if len(ips) == 2 {
		ip = ips[0]
	}
	value := commomDao.GetConfig("white")
	allowIps,_ := allowIpDao.GetAllowIp()
	if value == "1" {
		var allowLogin bool
		for _, alIp := range allowIps {
			if ip == alIp.Ip {
				allowLogin = true
				break
			}
		}
		if !allowLogin {
			middleware.ResponseError(c, 400, "ip不允许登录", errors.New("ip不允许登录"))
			return
		}
	}

	acc := accountDao.GetAccountByName(account.Name)
	if acc.AccountName != account.Name {
		middleware.ResponseError(c, 400, "用户名错误", errors.New("用户名错误"))
		return
	}

	if acc.Pwd != account.Pwd && acc.Pwd != md5V(account.Pwd) {
		middleware.ResponseError(c, 400, "密码错误", errors.New("密码错误"))
		return
	}
	now := time.Now().Unix()
	if acc.PwdUtil != 0 && now > int64(acc.PwdUtil) {
		middleware.ResponseError(c, 400, "密码已过期", errors.New("密码已过期"))
		return
	}

	t := int32(7 * 24 * 3600)
	if int64(acc.PwdUtil) - now < int64(7 * 24 * 3600) {
		t = acc.PwdUtil - int32(now)
	}
	token,err := SignToken(acc.Id, acc.UserName, account.Pwd, t)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}

	// 记录登录
	loginDao.RecordLogin(acc.Id, ip)
	data := map[string]string{
		"token": token,
	}
	middleware.ResponseSuccess(c, data)

}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func SignToken(id int32, name, pwd string, remainTime int32) (tokenStr string, err error) {
	// 带权限创建令牌
	claims := make(jwt.MapClaims)
	claims["id"] = id
	claims["name"] = name
	claims["pwd"] = pwd

	sec := time.Duration(remainTime)
	//sec := time.Duration(20)
	claims["exp"] = time.Now().Add(time.Second * sec).Unix() //自定义有效期，过期需要重新登录获取token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用自定义字符串加密 and get the complete encoded token as a string
	tokenStr, err = token.SignedString([]byte("popchat_jwt"))

	return tokenStr, err
}
