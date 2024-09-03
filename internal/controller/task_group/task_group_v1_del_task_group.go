package task_group

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"wktline-server/internal/dao"

	"wktline-server/api/task_group/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) DelTaskGroup(ctx context.Context, req *v1.DelTaskGroupReq) (res *v1.DelTaskGroupRes, err error) {
	taskCount, err := dao.Task.Ctx(ctx).Where("group_id", req.Id).Count()
	if err != nil {
		return nil, err
	}
	if taskCount > 0 {
		return nil, gerror.New("该分组下包含任务，无法删除")
	}
	err = service.Task().DelTaskGroup(ctx, model_task.DelTaskGroupInput{
		Id: req.Id,
	})
	return nil, err
}
