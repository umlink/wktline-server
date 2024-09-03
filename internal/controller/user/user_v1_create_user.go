package user

import (
	"context"

	"wktline-server/internal/model"
	"wktline-server/internal/service"
	"wktline-server/utility"

	"wktline-server/api/user/v1"
)

func (c *ControllerV1) CreateUser(ctx context.Context, req *v1.CreateUserReq) (res *v1.CreateUserRes, err error) {
	if err != nil {
		return nil, err
	}
	_, err = service.User().CreateUser(ctx, model.CreateUserInput{
		Id:       utility.GenUniIdByGuid(),
		Nickname: req.Nickname,
		Username: req.Username,
		Role:     req.Role,
		Status:   req.Status,
		Phone:    req.Phone,
		Email:    req.Email,
		Avatar:   req.Avatar,
	})
	return nil, err
}
