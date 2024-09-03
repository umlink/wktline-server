package task_status

import (
	"context"

	"wktline-server/api/task_status/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) UpdateTaskStatus(ctx context.Context, req *v1.UpdateTaskStatusReq) (res *v1.UpdateTaskStatusRes, err error) {
	err = service.Task().UpdateTaskStatus(ctx, model_task.UpdateTaskStatusInput{
		ProjectId: req.ProjectId,
		Id:        req.Id,
		Name:      req.Name,
		Color:     req.Color,
		Enum:      req.Enum,
	})
	return nil, err
}
