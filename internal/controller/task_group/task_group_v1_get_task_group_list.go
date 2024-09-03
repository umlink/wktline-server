package task_group

import (
	"context"

	"wktline-server/api/task_group/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) GetTaskGroupList(ctx context.Context, req *v1.GetTaskGroupListReq) (res *v1.GetTaskGroupListRes, err error) {
	out, err := service.Task().GetTaskGroup(ctx, model_task.GetTaskGroupInput{
		Keywords:  req.Keywords,
		ProjectId: req.ProjectId,
		PageNo:    req.PageNo,
		PageSize:  req.PageSize,
	})
	return &v1.GetTaskGroupListRes{
		List:     out.List,
		Total:    out.Total,
		PageSize: out.PageSize,
		PageNo:   out.PageNo,
	}, err
}
