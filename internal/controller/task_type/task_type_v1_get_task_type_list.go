package task_type

import (
	"context"

	"wktline-server/api/task_type/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) GetTaskTypeList(ctx context.Context, req *v1.GetTaskTypeListReq) (res *v1.GetTaskTypeListRes, err error) {
	out, err := service.Task().GetTaskType(ctx, model_task.GetTaskTypeInput{
		Keywords:  req.Keywords,
		ProjectId: req.ProjectId,
		PageNo:    req.PageNo,
		PageSize:  req.PageSize,
	})
	return &v1.GetTaskTypeListRes{
		List:     out.List,
		Total:    out.Total,
		PageSize: out.PageSize,
		PageNo:   out.PageNo,
	}, err
}
