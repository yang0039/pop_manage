package member_manage

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"pop-api/dal/dao"
	"pop-api/dal/dataobject"
	"pop-api/dto"
	"pop-api/middleware"
	"strconv"
)

func (service *UserController) QryUserPhone(c *gin.Context) {
	params := &dto.PhoneTransaction{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Qry == "" {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}
	phoneDao := dao.GetPhoneTransactionsDAO()
	if params.Limit == 0 {
		params.Limit = 20
	}
	var phoneTransactions []*dataobject.AuthPhoneTransactions
	var count int32
	if params.Type == 0 {
		phoneTransactions, count = phoneDao.GetPhoneTransactionsByPhone(params.Qry, params.Limit, params.Offset)
	} else {
		t,_ := strconv.Atoi(params.Qry)
		start := int64(t)
		end := start + 24 * 3600
		phoneTransactions, count = phoneDao.GetPhoneTransactionsByTime(start, end, params.Limit, params.Offset)
	}

	data := map[string]interface{}{
		"auth_phone_transactions": phoneTransactions,
		"count":                   count,
	}
	middleware.ResponseSuccess(c, data)
}
