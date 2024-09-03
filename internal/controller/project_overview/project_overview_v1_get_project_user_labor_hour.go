package project_overview

import (
	"context"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"

	"wktline-server/api/project_overview/v1"
)

func (c *ControllerV1) GetProjectUserLaborHour(ctx context.Context, req *v1.GetProjectUserLaborHourReq) (res *v1.GetProjectUserLaborHourRes, err error) {
	out, err := service.Project().GetProjectUserLaborHourStat(ctx, model_project.GetProjectUserLaborHourStatInput{
		ProjectId: req.ProjectId,
		GroupId:   req.GroupId,
		UserId:    req.UserId,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	return (*v1.GetProjectUserLaborHourRes)(out), err
}
