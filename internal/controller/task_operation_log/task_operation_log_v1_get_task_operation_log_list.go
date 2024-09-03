package task_operation_log

import (
	"context"

	"wktline-server/api/task_operation_log/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) GetTaskOperationLogList(ctx context.Context, req *v1.GetTaskOperationLogListReq) (res *v1.GetTaskOperationLogListRes, err error) {
	out, err := service.Task().GetTaskOperationLogList(ctx, model_task.GetOperationLogListInput{
		TaskId:   req.TaskId,
		Type:     req.Type,
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
	})
	return &v1.GetTaskOperationLogListRes{
		List:     out.List,
		Total:    out.Total,
		PageSize: out.PageSize,
		PageNo:   out.PageNo,
	}, err
}
