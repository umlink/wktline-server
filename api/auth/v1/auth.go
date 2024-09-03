package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type AuthReq struct {
	g.Meta `path:"/auth" method:"post" summary:"授权登录" tags:"Auth"`
	Code   string `json:"code" v:"required" dc:"授权 code"`
}
type AuthRes struct {
	Id       string    `json:"id" v:"required" dc:"用户id"`
	Username string    `json:"username" v:"required" dc:"用户名"`
	Nickname string    `json:"nickname" v:"required" dc:"用户昵称、别名"`
	Avatar   string    `json:"avatar" v:"required" dc:"头像"`
	Role     string    `json:"role" v:"required" dc:"角色id"`
	Status   int       `json:"status" v:"required" v:"required" dc:"状态"`
	Token    string    `json:"token" v:"required" dc:"token"`
	Expire   time.Time `json:"expire" v:"required" dc:"token"`
}
