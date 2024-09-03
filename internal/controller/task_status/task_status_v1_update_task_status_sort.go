package task_status

import (
	"context"

	"wktline-server/api/task_status/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) UpdateTaskStatusSort(ctx context.Context, req *v1.UpdateTaskStatusSortReq) (res *v1.UpdateTaskStatusSortRes, err error) {
	err = service.Task().UpdateTaskStatusSort(ctx, model_task.UpdateTaskStatusSortInput{
		ProjectId:   req.ProjectId,
		SortMapList: req.SortMapList,
	})
	return nil, err
}
