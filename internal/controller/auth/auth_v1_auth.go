package auth

import (
	"context"
	"wktline-server/internal/model"
	"wktline-server/internal/service"
	"wktline-server/utility"

	"wktline-server/api/auth/v1"
)

func (c *ControllerV1) Auth(ctx context.Context, req *v1.AuthReq) (res *v1.AuthRes, err error) {
	// 此处自行实现 code 的验证，获取用户到用户id,然后获取用户信息，生成 Token
	id := "用户 id"
	userInfo, err := service.User().GetUserInfoById(ctx, model.GetUserInfoInput{
		UserId: id,
	})
	out := &v1.AuthRes{
		Id:       userInfo.Id,
		Avatar:   userInfo.Avatar,
		Username: userInfo.Username,
		Nickname: userInfo.Nickname,
		Role:     userInfo.Role,
		Status:   userInfo.Status,
	}
	out.Token, out.Expire = service.Jwt().Auth().LoginHandler(ctx)
	utility.SetCookie(ctx, out.Token)
	return out, nil
}
