package project_user

import (
	"context"

	"wktline-server/api/project_user/v1"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"
)

func (c *ControllerV1) AddProjectUserByIds(ctx context.Context, req *v1.AddProjectUserByIdsReq) (res *v1.AddProjectUserByIdsRes, err error) {
	err = service.Project().AddProjectUserByIds(ctx, model_project.AddProjectUserByIdsInput{
		ProjectId: req.ProjectId,
		UserIds:   req.UserIds,
	})
	return nil, err
}
