package task_status

import (
	"context"

	"wktline-server/api/task_status/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) GetTaskStatusList(ctx context.Context, req *v1.GetTaskStatusListReq) (res *v1.GetTaskStatusListRes, err error) {
	out, err := service.Task().GetTaskStatus(ctx, model_task.GetTaskStatusInput{
		ProjectId: req.ProjectId,
		PageNo:    req.PageNo,
		PageSize:  req.PageSize,
	})
	return &v1.GetTaskStatusListRes{
		List:     out.List,
		Total:    out.Total,
		PageSize: out.PageSize,
		PageNo:   out.PageNo,
	}, err
}
