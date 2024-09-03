package task_labor_hour

import (
	"context"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"

	"wktline-server/api/task_labor_hour/v1"
)

func (c *ControllerV1) AddTaskLaborHour(ctx context.Context, req *v1.AddTaskLaborHourReq) (res *v1.AddTaskLaborHourRes, err error) {
	if _, err = service.Task().CheckUpdateForProjectUserAuth(ctx, req.ProjectId); err != nil {
		return nil, err
	}
	err = service.Task().AddTaskLaborHour(ctx, model_task.AddTaskLaborHourInput{
		TaskId:      req.TaskId,
		ProjectId:   req.ProjectId,
		Hour:        req.Hour,
		Date:        req.Date,
		Description: req.Description,
		UserId:      service.BizCtx().Get(ctx).User.Uid,
	})
	return nil, err
}
