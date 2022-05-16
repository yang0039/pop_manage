package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"pop-api/baselib/logger"
	"pop-api/dal/dao"
)

type ResponseCode int

//1000以下为通用码，1000以上为用户自定义码
const (
	SuccessCode ResponseCode = iota
	UndefErrorCode
	ValidErrorCode
	InternalErrorCode

	InvalidRequestErrorCode ResponseCode = 401
	CustomizeCode           ResponseCode = 1000

	GROUPALL_SAVE_FLOWERROR ResponseCode = 2001
)

type Response struct {
	ErrorCode ResponseCode `json:"err_no"`
	ErrorMsg  string       `json:"err_msg"`
	DisMsg    string       `json:"dis_msg"`
	Data      interface{}  `json:"data"`
	//TraceId   interface{}  `json:"trace_id"`
	//Stack     interface{}  `json:"stack"`
}

func ResponseError(c *gin.Context, code ResponseCode, disMsg string, err error) {
	//trace, _ := c.Get("trace")
	//traceContext, _ := trace.(*lib.TraceContext)
	//traceId := ""
	//if traceContext != nil {
	//	traceId = traceContext.TraceId
	//}
	//stack := ""
	//if c.Query("is_debug") == "1" || lib.GetConfEnv() == "dev" {
	//	stack = strings.Replace(fmt.Sprintf("%+v", err), err.Error()+"\n", "", -1)
	//}
	//resp := &Response{ErrorCode: code, ErrorMsg: err.Error(), Data: "", TraceId: traceId, Stack: stack}

	reId, _ := c.Get("record_id")
	id, _ := reId.(int32)
	if id != 0 {
		dao.GetRequestRecoreDAO().AddRequestResult(id, 0, disMsg)
	}
	logger.LogSugar.Errorf("uri:%s, err:%s", c.Request.RequestURI, err.Error())
	resp := &Response{ErrorCode: code, DisMsg: disMsg, ErrorMsg: err.Error(), Data: ""}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(200, err)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	//trace, _ := c.Get("trace")
	//traceContext, _ := trace.(*lib.TraceContext)
	//traceId := ""
	//if traceContext != nil {
	//	traceId = traceContext.TraceId
	//}
	//resp := &Response{ErrorCode: SuccessCode, ErrorMsg: "", Data: data, TraceId: traceId}

	reId, _ := c.Get("record_id")
	id, _ := reId.(int32)
	if id != 0 {
		dao.GetRequestRecoreDAO().AddRequestResult(id, 1, "")
	}
	resp := &Response{ErrorCode: SuccessCode, DisMsg: "成功", ErrorMsg: "", Data: data}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
