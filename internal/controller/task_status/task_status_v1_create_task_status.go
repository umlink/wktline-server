package task_status

import (
	"context"
	"wktline-server/utility"

	"wktline-server/api/task_status/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) CreateTaskStatus(ctx context.Context, req *v1.CreateTaskStatusReq) (res *v1.CreateTaskStatusRes, err error) {
	statusId := utility.GenUniIdByGuid()
	if err != nil {
		return nil, err
	}
	err = service.Task().CreateTaskStatus(ctx, model_task.CreateTaskStatusInput{
		Id:         statusId,
		ProjectId:  req.ProjectId,
		Name:       req.Name,
		Color:      req.Color,
		Enum:       req.Enum,
		OperatorId: service.BizCtx().Get(ctx).User.Uid,
	})
	return &v1.CreateTaskStatusRes{
		Id: statusId,
	}, err
}
