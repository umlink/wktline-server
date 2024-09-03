package task_actor

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"wktline-server/api/task_actor/v1"
	"wktline-server/internal/model"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) DelTaskActor(ctx context.Context, req *v1.DelTaskActorReq) (res *v1.DelTaskActorRes, err error) {
	if _, err = service.Task().CheckUpdateForProjectUserAuth(ctx, req.ProjectId); err != nil {
		return nil, err
	}
	task, err := service.Task().GetTaskDetail(model_task.GetTaskDetailInput{Id: req.TaskId})
	if err != nil {
		return nil, err
	}
	if g.IsEmpty(task) {
		return nil, gerror.New("任务不存在")
	}
	userIdsArray := garray.NewStrArrayFrom(gconv.Strings(req.UserIds), true)
	// 自动过滤任务创建者和任务执行者
	if userIdsArray.Contains(task.HandlerId) {
		userIdsArray.RemoveValue(task.HandlerId)
	}
	if userIdsArray.Contains(task.CreatorId) {
		userIdsArray.RemoveValue(task.CreatorId)
	}
	err = service.Task().DelTaskActor(ctx, model_task.DelTaskActorInput{
		TaskId:  req.TaskId,
		UserIds: req.UserIds,
	})
	go func() {
		err = service.Message().CreateTaskActionUserMsg(ctx, model.CreateTaskActionUsersMsgInput{
			ProjectId: req.ProjectId,
			TaskId:    req.TaskId,
			UserIds:   gconv.Strings(userIdsArray),
			Title:     "你被移出任务【%s】",
			MsgType:   "TASK",
		})
		fmt.Println(err)
	}()
	return nil, err
}
