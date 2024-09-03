// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"wktline-server/internal/model"
)

type (
	IUser interface {
		GetUserByUsernameAndPassword(ctx context.Context, in model.UserLoginInput) (user *model.UserLoginOut, err error)
		CreateUser(ctx context.Context, in model.CreateUserInput) (res *model.CreateUserOutput, err error)
		CreateBatchUser(ctx context.Context, in model.CreateBatchUserInput) (err error)
		DeleteUserById(ctx context.Context, in model.DeleteUserInput) (err error)
		UpdateUserInfo(ctx context.Context, in model.UpdateUserInfoInput) (err error)
		GetUserInfoById(ctx context.Context, in model.GetUserInfoInput) (user *model.GetUserInfoOutput, err error)
		GetUserList(ctx context.Context, in model.GetUserListInput) (userList *model.GetUserListOutput, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
