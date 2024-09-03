package user

import (
	"context"
	"wktline-server/api/user/v1"
	"wktline-server/internal/model"
	"wktline-server/internal/service"
	"wktline-server/utility"
)

// GetUserInfo UserInfo 只有合法的 token 才能进来
func (c *ControllerV1) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (res *v1.GetUserInfoRes, err error) {
	userId := service.BizCtx().Get(ctx).User.Uid
	out, err := service.User().GetUserInfoById(ctx, model.GetUserInfoInput{
		UserId: userId,
	})
	// 用户信息更新过，则刷新 token,以更新权限类信息
	if out.IsUpdate == 1 {
		token, _ := service.Jwt().Auth().LoginHandler(ctx)
		utility.SetCookie(ctx, token)
		var defaultVal int = 0
		var IsUpdate = &defaultVal
		if err = service.User().UpdateUserInfo(ctx, model.UpdateUserInfoInput{
			Id:       userId,
			IsUpdate: IsUpdate,
		}); err != nil {
			return nil, err
		}
	}
	return &v1.GetUserInfoRes{
		Id:       out.Id,
		Username: out.Username,
		Nickname: out.Nickname,
		Avatar:   out.Avatar,
		Role:     out.Role,
		Status:   out.Status,
	}, nil
}
