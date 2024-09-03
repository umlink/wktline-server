package task_status

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"wktline-server/api/task_status/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) DelTaskStatus(ctx context.Context, req *v1.DelTaskStatusReq) (res *v1.DelTaskStatusRes, err error) {
	currentStatusTaskNum, err := service.Task().GetTaskCountByStatus(ctx, req.Id)
	if currentStatusTaskNum > 0 {
		return nil, gerror.New("该状态下含有任务，无法删除")
	}
	err = service.Task().DelTaskStatus(ctx, model_task.DelTaskStatusInput{
		Id:        req.Id,
		ProjectId: req.ProjectId,
	})
	return nil, err
}
