package task_type

import (
	"context"

	"wktline-server/api/task_type/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) UpdateTaskType(ctx context.Context, req *v1.UpdateTaskTypeReq) (res *v1.UpdateTaskTypeRes, err error) {
	err = service.Task().UpdateTaskType(ctx, model_task.UpdateTaskTypeInput{
		ProjectId: req.ProjectId,
		Id:        req.Id,
		Name:      req.Name,
		Color:     req.Color,
	})
	return nil, err
}
