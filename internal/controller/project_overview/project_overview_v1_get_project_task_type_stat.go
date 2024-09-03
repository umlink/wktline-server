package project_overview

import (
	"context"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"

	"wktline-server/api/project_overview/v1"
)

func (c *ControllerV1) GetProjectTaskTypeStat(ctx context.Context, req *v1.GetProjectTaskTypeStatReq) (res *v1.GetProjectTaskTypeStatRes, err error) {
	out, err := service.Project().GetProjectTaskTypeStatistics(ctx, model_project.GetProjectTaskTypeStatInput{
		ProjectId: req.ProjectId,
		GroupId:   req.GroupId,
		UserId:    req.UserId,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	return (*v1.GetProjectTaskTypeStatRes)(out), err
}
