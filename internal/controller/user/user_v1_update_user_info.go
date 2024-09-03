package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"wktline-server/api/user/v1"
	"wktline-server/internal/model"
	"wktline-server/internal/service"
)

func (c *ControllerV1) UpdateUserInfo(ctx context.Context, req *v1.UpdateUserInfoReq) (res *v1.UpdateUserInfoRes, err error) {
	err = service.User().UpdateUserInfo(ctx, model.UpdateUserInfoInput{
		Id:       req.Id,
		Username: req.Username,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Status:   req.Status,
		Role:     req.Role,
		Phone:    req.Phone,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}
	return nil, gerror.NewCode(gcode.CodeOK, "更新成功")
}
