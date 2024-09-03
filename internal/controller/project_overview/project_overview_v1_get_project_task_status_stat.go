package project_overview

import (
	"context"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"

	"wktline-server/api/project_overview/v1"
)

func (c *ControllerV1) GetProjectTaskStatusStat(ctx context.Context, req *v1.GetProjectTaskStatusStatReq) (res *v1.GetProjectTaskStatusStatRes, err error) {
	out, err := service.Project().GetProjectTaskStatusStatistics(ctx, model_project.GetProjectTaskStatusStatInput{
		ProjectId: req.ProjectId,
		GroupId:   req.GroupId,
		UserId:    req.UserId,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	return (*v1.GetProjectTaskStatusStatRes)(out), err
}
