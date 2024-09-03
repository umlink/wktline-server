package project

import (
	"context"

	"wktline-server/api/project/v1"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"
)

func (c *ControllerV1) DelProject(ctx context.Context, req *v1.DelProjectReq) (res *v1.DelProjectRes, err error) {
	err = service.Project().DelProjectById(ctx, model_project.DelProjectInput{
		Id: req.ProjectId,
	})
	return nil, err
}
