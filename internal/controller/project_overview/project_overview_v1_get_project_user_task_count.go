package project_overview

import (
	"context"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"

	"wktline-server/api/project_overview/v1"
)

func (c *ControllerV1) GetProjectUserTaskCount(ctx context.Context, req *v1.GetProjectUserTaskCountReq) (res *v1.GetProjectUserTaskCountRes, err error) {
	out, err := service.Project().GetProjectUserTaskCountStat(ctx, model_project.GetProjectUserTaskCountStatInput{
		ProjectId: req.ProjectId,
		GroupId:   req.GroupId,
		UserId:    req.UserId,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	return (*v1.GetProjectUserTaskCountRes)(out), err
}
