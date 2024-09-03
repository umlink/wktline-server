package model

import v1 "wktline-server/api/login/v1"

type UserBaseInfo struct {
	Id       string `json:"id" v:"required" dc:"用户id"`
	Username string `json:"username" v:"required" dc:"用户名"`
	Nickname string `json:"nickname" v:"required" dc:"用户昵称、别名"`
	Avatar   string `json:"avatar" v:"required" dc:"头像"`
	Role     string `json:"role" v:"required" dc:"角色"`
	Status   int    `json:"status" v:"required" v:"required" dc:"状态"`
}

type UserLoginInput struct {
	Username string
	Password string
}
type UserLoginOut = v1.LoginRes

type CreateUserItem struct {
	Id       string
	Username string
	Nickname string
	Avatar   string
	Role     string
	Status   int
	Phone    int64
	Email    string
}

type CreateUserInput = CreateUserItem
type CreateUserOutput struct {
	Id string `json:"id" dc:"用户id"`
}

type CreateBatchUserInput = []CreateUserItem

type DeleteUserInput struct {
	Id string `json:"id" v:"required" dc:"用户id"`
}

type UpdateUserInfoInput struct {
	Id       string
	Username *string
	Nickname *string
	Avatar   *string
	Role     *string
	Status   *int
	Phone    *int64
	Email    *string
	IsUpdate *int
}

type GetUserListInput struct {
	Keywords string
	Role     string
	Status   int
	PageNo   int
	PageSize int
}
type GetUserListOutput struct {
	List     []*UserBaseInfo
	PageNo   int
	PageSize int
	Total    int
}

type GetUserInfoInput struct {
	UserId string
}

type GetUserInfoOutput struct {
	Id       string
	Username string
	Nickname string
	Avatar   string
	Role     string
	Status   int
	IsUpdate int
}
