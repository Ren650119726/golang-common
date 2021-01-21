package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//前端传入参数错误（包括表单，查询参数等）的返回，状态码返回400
func ValidationErrorHttpReturn(c *gin.Context, object interface{}, err error, msg ...string) {
	if len(msg) == 0 {
		msg = make([]string, 1)
		msg[0] = "错误的参数"
	}
	wvf := WrapValidateFailed(object, err, msg[0])
	HttpReturn(c, wvf.Code, wvf.Data, wvf.Msg)
}

//http返回的内层结构体
type HttpInnerResult struct {
	Code int         `json:"code"` //真实的http状态码
	Data interface{} `json:"data"` //返回的数据
	Msg  string      `json:"msg"`  //错误信息，如无错误，该字段为空
}

//服务器内部错误的返回（如：数据库错误，内部panic等）
func InnerErrorHttpReturn(c *gin.Context) {
	HttpReturn(c, InnerErrorCode, "", InnerErrorMsg)
}

//没有任何问题的返回
func StatusOKHttpReturn(c *gin.Context, data interface{}, msg ...string) {
	if len(msg) == 0 {
		msg = make([]string, 1)
		msg[0] = "success"
	}
	HttpReturn(c, http.StatusOK, data, msg[0])
}

//统一的http返回格式，状态码都为200，真实状态码都在httpInnerResult.Code里
func HttpReturn(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(http.StatusOK, HttpInnerResult{
		Code: code,
		Data: data,
		Msg:  msg,
	})
	c.Set("dpi_response", code)
	c.Abort()
}

// token相关的错误
func TokenErrorHttpReturn(c *gin.Context, data interface{}, msg ...string) {
	if len(msg) == 0 {
		msg = make([]string, 1)
		msg[0] = "未授权或需要重新登陆"
	}
	HttpReturn(c, http.StatusUnauthorized, data, msg[0])
}
