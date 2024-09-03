package model

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Context 请求上下文结构
type Context struct {
	User *ContextPayloadUser // 上下文用户信息
	Data g.Map               // 自定KV变量不固定
}

type ContextPayloadUser struct {
	Uid      string
	Username string
	Avatar   string
	Role     string
}
