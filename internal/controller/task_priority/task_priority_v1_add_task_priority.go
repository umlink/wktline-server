package task_priority

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"wktline-server/api/task_priority/v1"
	"wktline-server/internal/dao"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) AddTaskPriority(ctx context.Context, req *v1.AddTaskPriorityReq) (res *v1.AddTaskPriorityRes, err error) {
	count, err := dao.Task.Ctx(ctx).Where("priority", req.Id).Count()
	if count > 0 {
		return nil, gerror.NewCode(gcode.CodeNotSupported, "该优先级已存在")
	}
	err = service.Task().AddTaskPriority(ctx, model_task.AddTaskPriorityInput{
		Id:    req.Id,
		Name:  req.Name,
		Color: req.Color,
	})
	return nil, err
}
