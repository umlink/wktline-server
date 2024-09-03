package task_operation_log

import (
	"context"
	"wktline-server/utility"

	"wktline-server/api/task_operation_log/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) CreateTaskOperationLog(ctx context.Context, req *v1.CreateTaskOperationLogReq) (res *v1.CreateTaskOperationLogRes, err error) {
	logId := utility.GenUniIdByGuid()
	if err != nil {
		return nil, err
	}
	err = service.Task().CreateTaskOperationLog(model_task.CreateOperationLogInput{
		Id:      logId,
		TaskId:  req.TaskId,
		UserId:  service.BizCtx().Get(ctx).User.Uid,
		Type:    req.Type,
		Name:    req.Name,
		Content: req.Content,
		Desc:    req.Desc,
	})
	return nil, err
}
