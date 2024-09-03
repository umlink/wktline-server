package user

import (
	"context"

	"wktline-server/api/user/v1"
	"wktline-server/internal/model"
	"wktline-server/internal/service"
)

func (c *ControllerV1) GetUserList(ctx context.Context, req *v1.GetUserListReq) (res *v1.GetUserListRes, err error) {
	out, err := service.User().GetUserList(ctx, model.GetUserListInput{
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
		Keywords: req.Keywords,
		Role:     req.Role,
	})
	return &v1.GetUserListRes{
		List:     out.List,
		PageNo:   out.PageNo,
		PageSize: out.PageSize,
		Total:    out.Total,
	}, err
}
