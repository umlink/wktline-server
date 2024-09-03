package response

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

// JsonRes 数据返回通用JSON数据结构
type JsonRes struct {
	Code    int         `json:"code" v:"required" dc:"状态码"`     // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message" v:"required" dc:"提示信息"` // 提示信息
	Data    interface{} `json:"data" v:"required" dc:"返回数据"`    // 返回数据(业务接口定义具体数据结构)
	Success bool        `json:"success" v:"required"`
}

// Json 返回标准JSON数据。
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = r.GetHandlerResponse()
	} else {
		responseData = nil
	}
	r.Response.WriteJson(JsonRes{
		Code:    code,
		Message: message,
		Data:    responseData,
		Success: code == 0,
	})
}

// JsonExit 返回标准JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, code int, message string, data ...interface{}) {
	Json(r, code, message, data...)
	r.Exit()
}
