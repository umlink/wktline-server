package task

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"wktline-server/api/task/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) GetChildTaskList(ctx context.Context, req *v1.GetChildTaskListReq) (res *v1.GetChildTaskListRes, err error) {
	task, err := c.GetTaskDetail(ctx, &v1.GetTaskDetailReq{
		Id: req.ParentId,
	})
	if err != nil {
		return nil, err
	}
	if g.IsEmpty(task) {
		return nil, gerror.New("父任务不存在")
	}
	out, err := service.Task().GetChildTaskList(model_task.GetChildTaskListInput{
		ParentId: req.ParentId,
		PageSize: req.PageSize,
		PageNo:   req.PageNo,
	})
	return (*v1.GetChildTaskListRes)(out), err
}
