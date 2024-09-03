package task_operation_log

import (
	"context"

	"wktline-server/api/task_operation_log/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) DelTaskOperationLog(ctx context.Context, req *v1.DelTaskOperationLogReq) (res *v1.DelTaskOperationLogRes, err error) {
	err = service.Task().DelTaskOperationLog(ctx, model_task.DelOperationLogInput{
		Id: req.Id,
	})
	return nil, err
}
