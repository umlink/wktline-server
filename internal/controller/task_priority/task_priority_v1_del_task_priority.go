package task_priority

import (
	"context"
	"wktline-server/internal/dao"
	"wktline-server/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"wktline-server/api/task_priority/v1"
)

func (c *ControllerV1) DelTaskPriority(ctx context.Context, req *v1.DelTaskPriorityReq) (res *v1.DelTaskPriorityRes, err error) {
	count, err := dao.Task.Ctx(ctx).Where("priority", req.Id).Count()
	if count > 0 {
		return nil, gerror.NewCode(gcode.CodeNotSupported, "该优先级已被使用，无法删除")
	}
	return nil, service.Task().DelTaskPriority(ctx, req.Id)
}
