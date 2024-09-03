package task_actor

import (
	"context"
	"fmt"
	"wktline-server/internal/model"

	"wktline-server/api/task_actor/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) AddTaskActor(ctx context.Context, req *v1.AddTaskActorReq) (res *v1.AddTaskActorRes, err error) {
	if _, err = service.Task().CheckUpdateForProjectUserAuth(ctx, req.ProjectId); err != nil {
		return nil, err
	}
	err = service.Task().AddTaskActor(ctx, model_task.AddTaskActorInput{
		TaskId:  req.TaskId,
		UserIds: req.UserIds,
	})
	go func() {
		err = service.Message().CreateTaskActionUserMsg(ctx, model.CreateTaskActionUsersMsgInput{
			TaskId:    req.TaskId,
			ProjectId: req.ProjectId,
			UserIds:   req.UserIds,
			Title:     "你被添加到任务【%s】参与人",
			MsgType:   "TASK",
		})
		fmt.Println(err)
	}()
	return nil, err
}
