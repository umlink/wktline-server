package project_invite

import (
	"context"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"

	"wktline-server/api/project_invite/v1"
)

func (c *ControllerV1) GenProjectInviteCode(ctx context.Context, req *v1.GenProjectInviteCodeReq) (res *v1.GenProjectInviteCodeRes, err error) {
	code, err := service.Project().GenProjectInviteCode(ctx, model_project.GenProjectInviteCodeInput{
		ProjectId:      req.ProjectId,
		Deadline:       req.Deadline,
		MaxInviteCount: req.MaxInviteCount,
	})
	return &v1.GenProjectInviteCodeRes{
		Id: code,
	}, err
}
