package task

import (
	"context"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"

	"wktline-server/api/task/v1"
)

func (c *ControllerV1) GetTaskListByInterval(ctx context.Context, req *v1.GetTaskListByIntervalReq) (res *v1.GetTaskListByIntervalRes, err error) {
	out, err := service.Task().GetTaskListByInterval(ctx, model_task.GetTaskListByIntervalInput{
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	return (*v1.GetTaskListByIntervalRes)(out), err
}
