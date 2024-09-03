package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"wktline-server/internal/model"
)

type CommonPaginationReq struct {
	PageNo   int `json:"pageNo" d:"1" dc:"分页号码，默认1"`
	PageSize int `json:"pageSize" d:"10" dc:"分页数量，最大100"`
}

type GetUserInfoReq struct {
	g.Meta `path:"/user/getUseInfo" method:"get" summary:"获取用户信息-登录信息" tags:"User"`
}
type GetUserInfoRes struct {
	Id       string `json:"id" v:"required" dc:"用户id"`
	Username string `json:"username" v:"required" dc:"用户名"`
	Nickname string `json:"nickname" v:"required" dc:"用户昵称、别名"`
	Avatar   string `json:"avatar" v:"required" dc:"头像"`
	Role     string `json:"role" v:"required" dc:"角色"`
	Status   int    `json:"status" v:"required" v:"required" dc:"状态"`
}

type GetUserInfoByIdReq struct {
	g.Meta `path:"/user/getUserInfoById" method:"get" summary:"根据用户id获取用户信息" tags:"User"`
	Id     string `path:"/id" v:"required#用户 id 不能为空" dc:"用户 id"`
}
type GetUserInfoByIdRes struct {
	Id       string `json:"id" v:"required" dc:"用户id"`
	Username string `json:"username" v:"required" dc:"用户名"`
	Nickname string `json:"nickname" v:"required" dc:"用户昵称、别名"`
	Avatar   string `json:"avatar" v:"required" dc:"头像"`
	Role     string `json:"role" v:"required" dc:"角色"`
	Status   int    `json:"status" v:"required" v:"required" dc:"状态"`
}

type GetUserListReq struct {
	g.Meta   `path:"/user/getUserList" method:"get" summary:"获取用户列表" tags:"User"`
	Keywords string `json:"keywords" dc:"用户名/用户昵称"`
	Role     string `json:"role" dc:"角色"`
	CommonPaginationReq
}
type GetUserListRes struct {
	List     []*model.UserBaseInfo `json:"list" v:"required" dc:"用户列表"`
	PageNo   int                   `json:"pageNo" v:"required" dc:"页码"`
	PageSize int                   `json:"pageSize" v:"required" dc:"页码大小"`
	Total    int                   `json:"total" v:"required" dc:"总数"`
}

type DelUserByIdReq struct {
	g.Meta `path:"/user/delUserById" method:"post" summary:"根据用户id删除用户" tags:"User"`
	Id     string `json:"id" v:"required#用户 id 不能为空" dc:"用户 id"`
}
type DelUserByIdRes struct{}

type UpdateUserInfoReq struct {
	g.Meta   `path:"/user/updateUserInfo" method:"post" summary:"根据用户id更新用户信息" tags:"User"`
	Id       string  `json:"id" v:"required#用户 id 不能为空" dc:"用户id"`
	Username *string `json:"username" dc:"用户名"`
	Nickname *string `json:"nickname" dc:"用户昵称、别名"`
	Phone    *int64  `json:"phone" dc:"手机号"`
	Email    *string `json:"email" dc:"邮箱"`
	Avatar   *string `json:"avatar" dc:"头像"`
	Role     *string `json:"role" dc:"角色"`
	Status   *int    `json:"status" dc:"状态"`
}
type UpdateUserInfoRes struct{}

type CreateUserReq struct {
	g.Meta   `path:"/user/createUser" method:"post" summary:"创建新用户" tags:"User"`
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Nickname string `json:"nickname" v:"required#用户昵称不能为空" dc:"用户昵称、别名"`
	Phone    int64  `json:"phone" v:"required#用户手机号不能为空" dc:"手机号"`
	Email    string `json:"email" v:"required#用户邮箱不能为空" dc:"邮箱"`
	Avatar   string `json:"avatar" v:"required#用户头像不能为空" dc:"头像"`
	Role     string `json:"role" v:"required" dc:"角色"`
	Status   int    `json:"status" dc:"状态"`
}
type CreateUserRes struct{}
