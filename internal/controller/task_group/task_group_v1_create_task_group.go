package task_group

import (
	"context"
	"wktline-server/utility"

	"wktline-server/api/task_group/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) CreateTaskGroup(ctx context.Context, req *v1.CreateTaskGroupReq) (res *v1.CreateTaskGroupRes, err error) {
	groupId := utility.GenUniIdByGuid()
	if err != nil {
		return nil, err
	}
	err = service.Task().CreateTaskGroup(ctx, model_task.CreateTaskGroupInput{
		Id:          groupId,
		ProjectId:   req.ProjectId,
		GroupName:   req.GroupName,
		Description: req.Description,
		OperatorId:  service.BizCtx().Get(ctx).User.Uid,
	})
	return &v1.CreateTaskGroupRes{
		Id: groupId,
	}, err
}
