// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package user

import (
	"context"
	
	"wktline-server/api/user/v1"
)

type IUserV1 interface {
	GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (res *v1.GetUserInfoRes, err error)
	GetUserInfoById(ctx context.Context, req *v1.GetUserInfoByIdReq) (res *v1.GetUserInfoByIdRes, err error)
	GetUserList(ctx context.Context, req *v1.GetUserListReq) (res *v1.GetUserListRes, err error)
	DelUserById(ctx context.Context, req *v1.DelUserByIdReq) (res *v1.DelUserByIdRes, err error)
	UpdateUserInfo(ctx context.Context, req *v1.UpdateUserInfoReq) (res *v1.UpdateUserInfoRes, err error)
	CreateUser(ctx context.Context, req *v1.CreateUserReq) (res *v1.CreateUserRes, err error)
}


