package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" summary:"登录" tags:"Login"`
	Username string `json:"username" v:"required" dc:"用户名"`
	Password string `json:"password" v:"required" dc:"密码"`
}
type LoginRes struct {
	Id       int64     `json:"id" v:"required" dc:"用户id"`
	Username string    `json:"username" v:"required" dc:"用户名"`
	Nickname string    `json:"nickname" v:"required" dc:"用户昵称、别名"`
	Avatar   string    `json:"avatar" v:"required" dc:"头像"`
	Role     string    `json:"role" v:"required" dc:"角色id"`
	Status   int       `json:"status" v:"required" v:"required" dc:"状态"`
	Token    string    `json:"token" v:"required" dc:"token"`
	Expire   time.Time `json:"expire" v:"required" dc:"token"`
}
