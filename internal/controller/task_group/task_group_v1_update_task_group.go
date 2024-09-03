package task_group

import (
	"context"

	"wktline-server/api/task_group/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) UpdateTaskGroup(ctx context.Context, req *v1.UpdateTaskGroupReq) (res *v1.UpdateTaskGroupRes, err error) {
	err = service.Task().UpdateTaskGroup(ctx, model_task.UpdateTaskGroupInput{
		Id:          req.Id,
		GroupName:   req.GroupName,
		Description: req.Description,
	})
	return nil, err
}
