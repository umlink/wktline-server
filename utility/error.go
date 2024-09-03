package utility

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"strings"
)

type pErr struct {
	maps map[int]string
}

var Err = &pErr{
	maps: map[int]string{
		0:     "请求成功",
		10001: "用户名或密码错误",
		10002: "用户不存在",
		99999: "未知错误",
	},
}

// GetMsg 获取code码对应的msg
func (c *pErr) GetMsg(code int) string {
	return c.maps[code]
}

// Skip 抛出一个业务级别的错误，不会打印错误堆栈信息
func (c *pErr) Skip(code int, msg ...string) (err error) {
	var msgStr string
	if len(msg) == 0 {
		msgStr = c.GetMsg(code)
	} else {
		msg = append([]string{c.GetMsg(code)}, msg...)
		msgStr = strings.Join(msg, ", ")
	}
	// NewWithOption 在低版本的 gf 上不存在，请改用 NewOption
	return gerror.NewWithOption(gerror.Option{
		Stack: false,
		Text:  msgStr,
		Code:  gcode.New(code, "", nil),
	})
}

// Sys 抛出一个系统级别的错误，使用特殊的code码：99999
// !!! 使用该方法时，它会打印错误堆栈信息到日志，但是一定不要把任何错误信息抛出到客户端，防止泄露系统信息
// !!! 推荐做法是在后置中间件中捕获 code 99999 的错误，然后返回给客户端一个统一的错误提示
//func (c *pErr) Sys(err error) error {
//	return gerror.NewCode(gcode.New(CodeErrSys, "", nil), err.Error())
//}
