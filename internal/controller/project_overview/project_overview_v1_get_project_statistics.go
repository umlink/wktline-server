package project_overview

import (
	"context"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"

	"wktline-server/api/project_overview/v1"
)

func (c *ControllerV1) GetProjectStatistics(ctx context.Context, req *v1.GetProjectStatisticsReq) (res *v1.GetProjectStatisticsRes, err error) {
	out, _ := service.Project().GetProjectStatistics(ctx, model_project.GetProjectStatisticsInput{
		ProjectId: req.ProjectId,
		GroupId:   req.GroupId,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	return (*v1.GetProjectStatisticsRes)(out), nil
}
