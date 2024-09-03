package task_labor_hour

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"wktline-server/internal/dao"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"

	"wktline-server/api/task_labor_hour/v1"
)

func (c *ControllerV1) DelTaskLaborHour(ctx context.Context, req *v1.DelTaskLaborHourReq) (res *v1.DelTaskLaborHourRes, err error) {
	if _, err = service.Task().CheckUpdateForProjectUserAuth(ctx, req.ProjectId); err != nil {
		return nil, err
	}
	data, err := dao.TaskLaborHour.Ctx(ctx).Fields("hour,user_id").WherePri(req.Id).One()
	if err != nil {
		return nil, err
	}
	if gconv.String(data["user_id"]) != service.BizCtx().Get(ctx).User.Uid {
		return nil, gerror.New("只能删除自己的工时")
	}
	err = service.Task().DelLaborHourById(ctx, model_task.DelTaskLaborHourInput{
		Id:     req.Id,
		TaskId: req.TaskId,
		Hour:   gconv.Float64(data["hour"]),
	})
	return nil, err
}
