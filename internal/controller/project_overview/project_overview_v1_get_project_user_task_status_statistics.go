package project_overview

import (
	"context"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"

	"wktline-server/api/project_overview/v1"
)

func (c *ControllerV1) GetProjectUserTaskStatusStatistics(ctx context.Context, req *v1.GetProjectUserTaskStatusStatisticsReq) (res *v1.GetProjectUserTaskStatusStatisticsRes, err error) {
	out, err := service.Project().GetProjectUserTaskStatusStatistics(ctx, model_project.GetProjectUserTaskStatusStatisticsInput{
		ProjectId: req.ProjectId,
		GroupId:   req.GroupId,
		UserId:    req.UserId,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	return (*v1.GetProjectUserTaskStatusStatisticsRes)(out), err
}
