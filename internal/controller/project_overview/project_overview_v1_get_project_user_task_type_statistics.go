package project_overview

import (
	"context"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"

	"wktline-server/api/project_overview/v1"
)

func (c *ControllerV1) GetProjectUserTaskTypeStatistics(ctx context.Context, req *v1.GetProjectUserTaskTypeStatisticsReq) (res *v1.GetProjectUserTaskTypeStatisticsRes, err error) {
	out, err := service.Project().GetProjectUserTaskTypeStatistics(ctx, model_project.GetProjectUserTaskTypeStatisticsInput{
		ProjectId: req.ProjectId,
		GroupId:   req.GroupId,
		UserId:    req.UserId,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	return (*v1.GetProjectUserTaskTypeStatisticsRes)(out), err
}
