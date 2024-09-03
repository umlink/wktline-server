package task_actor

import (
	"context"

	"wktline-server/api/task_actor/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) GetTaskActor(ctx context.Context, req *v1.GetTaskActorReq) (res *v1.GetTaskActorRes, err error) {
	out, err := service.Task().GetTaskActorList(ctx, model_task.GetTaskActorListInput{
		TaskId:   req.TaskId,
		PageSize: req.PageSize,
		PageNo:   req.PageNo,
	})
	return &v1.GetTaskActorRes{
		List:     out.List,
		Total:    out.Total,
		PageSize: out.PageSize,
		PageNo:   out.PageNo,
	}, err
}
