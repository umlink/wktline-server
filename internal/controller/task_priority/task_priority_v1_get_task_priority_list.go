package task_priority

import (
	"context"

	"wktline-server/api/task_priority/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) GetTaskPriorityList(ctx context.Context, req *v1.GetTaskPriorityListReq) (res *v1.GetTaskPriorityListRes, err error) {
	out, err := service.Task().GetTaskPriority(ctx, model_task.GetTaskPriorityInput{
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
	})
	return &v1.GetTaskPriorityListRes{
		List:     out.List,
		Total:    out.Total,
		PageSize: out.PageSize,
		PageNo:   out.PageNo,
	}, err
}
