package task_type

import (
	"context"

	"wktline-server/api/task_type/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) DelTaskType(ctx context.Context, req *v1.DelTaskTypeReq) (res *v1.DelTaskTypeRes, err error) {
	err = service.Task().DelTaskType(ctx, model_task.DelTaskTypeInput{
		ProjectId: req.ProjectId,
		Id:        req.Id,
	})
	return nil, err
}
