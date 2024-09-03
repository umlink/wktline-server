package user

import (
	"context"

	"wktline-server/api/user/v1"
	"wktline-server/internal/model"
	"wktline-server/internal/service"
)

func (c *ControllerV1) GetUserInfoById(ctx context.Context, req *v1.GetUserInfoByIdReq) (res *v1.GetUserInfoByIdRes, err error) {
	out, err := service.User().GetUserInfoById(ctx, model.GetUserInfoInput{
		UserId: req.Id,
	})
	return &v1.GetUserInfoByIdRes{
		Id:       out.Id,
		Username: out.Username,
		Nickname: out.Nickname,
		Avatar:   out.Avatar,
		Role:     out.Role,
		Status:   out.Status,
	}, err
}
