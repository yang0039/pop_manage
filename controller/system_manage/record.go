package system_manage

import (
	"github.com/gin-gonic/gin"
	"pop-api/dal/dao"
	"pop-api/dal/dataobject"
	"pop-api/dto"
	"pop-api/middleware"
)

// 获取操作日志
func (service *SystemController) GetRequestRecord(c *gin.Context) {
	params := &dto.Acount{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	//if params.Id == 0 {
	//	middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
	//	return
	//}
	if params.Limit == 0 {
		params.Limit = 20
	}
	account := dao.GetAccountDAO().GetAllAccount()
	accountMap := make(map[int32]dataobject.Account, len(account))
	for _, a := range account {
		accountMap[a.Id] = a
	}
	urls := dao.GetPermissionsUrlDAO().GetAllPermissionsUrl()
	urlMap := make(map[string]dataobject.PermissionsUrl, len(account))
	for _, u := range urls {
		urlMap[u.Url] = u
	}

	record, count := dao.GetRequestRecoreDAO().GetRequestRecords(params.Id, params.Limit, params.Offset)
	res := make([]map[string]interface{}, 0, len(record))
	for _, r := range record {
		re := map[string]interface{}{
			"account_id":   r.AccountId,
			"account_name": accountMap[r.AccountId].AccountName,
			"user_name":    accountMap[r.AccountId].UserName,
			"url":          r.Url,
			"method":       r.Method,
			"request_name": urlMap[r.Url].MethodName,
			"client_ip":    r.ClientIp,
			"request_data": r.ReqData,
			"is_success":   r.IsSuccess,
			"fail_reason":  r.Reason,
			"request_time": r.AddTime,
		}
		res = append(res, re)
	}

	data := map[string]interface{}{
		"record": res,
		"count":  count,
	}
	middleware.ResponseSuccess(c, data)

}
