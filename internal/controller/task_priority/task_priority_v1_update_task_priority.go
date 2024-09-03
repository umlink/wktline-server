package task_priority

import (
	"context"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"

	"wktline-server/api/task_priority/v1"
)

func (c *ControllerV1) UpdateTaskPriority(ctx context.Context, req *v1.UpdateTaskPriorityReq) (res *v1.UpdateTaskPriorityRes, err error) {
	return nil, service.Task().UpdateTaskPriority(ctx, model_task.UpdateTaskPriorityInput{
		Id:    req.Id,
		Name:  req.Name,
		Color: req.Color,
	})
}
