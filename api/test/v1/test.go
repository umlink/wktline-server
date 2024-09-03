package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TestConnectReq struct {
	g.Meta   `path:"/test" method:"get" summary:"测试" tags:"Test"`
	Username string `json:"username" dc:"用户名"`
}
type TestConnectRes struct{}
