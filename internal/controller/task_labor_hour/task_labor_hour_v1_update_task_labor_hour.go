package task_labor_hour

import (
	"context"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"

	"wktline-server/api/task_labor_hour/v1"
)

func (c *ControllerV1) UpdateTaskLaborHour(ctx context.Context, req *v1.UpdateTaskLaborHourReq) (res *v1.UpdateTaskLaborHourRes, err error) {
	if _, err = service.Task().CheckUpdateForProjectUserAuth(ctx, req.ProjectId); err != nil {
		return nil, err
	}
	err = service.Task().UpdateTaskLaborHour(ctx, model_task.UpdateTaskLaborHourInput{
		Id:          req.Id,
		TaskId:      req.TaskId,
		Hour:        req.Hour,
		Description: req.Description,
		Date:        req.Date,
	})
	return nil, err
}
