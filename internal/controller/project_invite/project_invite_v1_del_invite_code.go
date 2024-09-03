package project_invite

import (
	"context"
	"wktline-server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"

	"wktline-server/api/project_invite/v1"
)

func (c *ControllerV1) DelInviteCode(ctx context.Context, req *v1.DelInviteCodeReq) (res *v1.DelInviteCodeRes, err error) {
	count, err := service.Project().CheckProjectMyInvite(ctx, req.Code)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, gerror.New("未找到该邀请码")
	}
	err = service.Project().DelProjectMyInvite(ctx, req.Code)
	return nil, err
}
