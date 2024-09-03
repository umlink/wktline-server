package project_group

import (
	"context"

	"wktline-server/api/project_group/v1"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"
)

func (c *ControllerV1) UpdateProjectGroup(ctx context.Context, req *v1.UpdateProjectGroupReq) (res *v1.UpdateProjectGroupRes, err error) {
	err = service.Project().UpdateGroup(ctx, model_project.UpdateGroupInput{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		OperatorId:  service.BizCtx().Get(ctx).User.Uid,
	})
	return nil, err
}
