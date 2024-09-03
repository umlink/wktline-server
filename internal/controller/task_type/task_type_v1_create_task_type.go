package task_type

import (
	"context"
	"wktline-server/utility"

	"wktline-server/api/task_type/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) CreateTaskType(ctx context.Context, req *v1.CreateTaskTypeReq) (res *v1.CreateTaskTypeRes, err error) {
	typeId := utility.GenUniIdByGuid()
	if err != nil {
		return nil, err
	}
	err = service.Task().CreateTaskType(ctx, model_task.CreateTaskTypeInput{
		Id:         typeId,
		ProjectId:  req.ProjectId,
		Name:       req.Name,
		Color:      req.Color,
		OperatorId: service.BizCtx().Get(ctx).User.Uid,
	})
	return &v1.CreateTaskTypeRes{
		Id: typeId,
	}, err
}
