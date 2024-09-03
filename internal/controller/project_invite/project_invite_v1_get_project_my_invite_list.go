package project_invite

import (
	"context"
	"wktline-server/internal/service"

	"wktline-server/api/project_invite/v1"
)

func (c *ControllerV1) GetProjectMyInviteList(ctx context.Context, req *v1.GetProjectMyInviteListReq) (res *v1.GetProjectMyInviteListRes, err error) {
	out, err := service.Project().GetProjectMyInviteList(ctx, req.ProjectId)
	return (*v1.GetProjectMyInviteListRes)(out), err
}
