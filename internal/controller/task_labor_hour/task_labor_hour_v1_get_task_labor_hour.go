package task_labor_hour

import (
	"context"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"

	"wktline-server/api/task_labor_hour/v1"
)

func (c *ControllerV1) GetTaskLaborHour(ctx context.Context, req *v1.GetTaskLaborHourReq) (res *v1.GetTaskLaborHourRes, err error) {
	out, err := service.Task().GetTaskLaborHourList(ctx, model_task.GetTaskLaborHourListInput{
		TaskId:   req.TaskId,
		PageSize: req.PageSize,
		PageNo:   req.PageNo,
	})
	return &v1.GetTaskLaborHourRes{
		List:     out.List,
		Total:    out.Total,
		PageNo:   out.PageNo,
		PageSize: out.PageSize,
	}, err
}
