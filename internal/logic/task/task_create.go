package task

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	v1 "wktline-server/api/task_operation_log/v1"
	"wktline-server/internal/dao"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
	"wktline-server/utility"
)

func (s *sTask) CreateTask(ctx context.Context, in model_task.CreateTaskInput) (taskId string, err error) {
	userId := service.BizCtx().Get(ctx).User.Uid
	err = dao.Task.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		in.Id = utility.GenUniIdByGuid()
		if _, err = dao.Task.Ctx(ctx).Data(in).OmitEmptyData().Insert(); err != nil {
			return err
		}
		// 添加操作日志
		if err != nil {
			return err
		}
		if err = s.CreateTaskOperationLog(model_task.CreateOperationLogInput{
			Id:      utility.GenUniIdByGuid(),
			TaskId:  in.Id,
			Name:    "创建了任务",
			Content: *in.Name,
			Type:    v1.DynamicTaskCreate,
			UserId:  userId,
		}); err != nil {
			return err
		}
		// 添加任务参与人
		userIds := []string{userId}
		if err = s.AddTaskActor(ctx, model_task.AddTaskActorInput{
			TaskId:  in.Id,
			UserIds: userIds,
			IsNew:   true,
		}); err != nil {
			return err
		}
		return err
	})
	return in.Id, err
}
