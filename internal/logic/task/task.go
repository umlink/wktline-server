package task

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"sync"
	"wktline-server/internal/consts"
	"wktline-server/internal/dao"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

type sTask struct{}

func init() {
	service.RegisterTask(New())
}

func New() *sTask {
	return &sTask{}
}

func (s *sTask) CheckUpdateForProjectUserAuth(ctx context.Context, projectId string) (role string, err error) {
	user := service.BizCtx().Get(ctx).User
	// 非本人：系统管理严、超管和项目管理严都可编辑
	if user.Role == consts.User {
		pUser, _ := service.Project().GetProjectUserDetail(ctx, model_project.GetProjectUserDetailInput{
			ProjectId: projectId,
			UserId:    user.Uid,
		})
		if g.IsEmpty(pUser) {
			return "S-USER", gerror.New("非本项目成员无权更新")
		}
		if pUser.Role == consts.User {
			return "P-USER", nil
		} else {
			return "P-ADMIN", nil
		}
	}
	return "S-ADMIN", nil
}

func (s *sTask) DelTask(ctx context.Context, in model_task.DelTaskInput) (err error) {
	_, err = dao.Task.Ctx(ctx).WherePri(in.Id).Delete()
	return err
}

func (s *sTask) GetTaskCountById(ctx context.Context, id string) (count int, err error) {
	return dao.Task.Ctx(ctx).WherePri(id).Count()
}

func (s *sTask) UpdateTask(ctx context.Context, in model_task.UpdateTaskInput) (err error) {
	var task *model_task.GetTaskDetailOutput
	var status *model_task.GetTaskStatusByIdOutput
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if !g.IsEmpty(in.StatusId) {
			status, err = service.Task().GetTaskStatusById(ctx, *in.StatusId)
		}
	}()
	// 查出旧的任务详情 - 用于更新日志
	go func() {
		defer wg.Done()
		task, err = service.Task().GetTaskDetail(model_task.GetTaskDetailInput{
			Id: in.Id,
		})
	}()
	wg.Wait()
	g.Dump(status)
	// 若任务处理人变更，且不再任务参与人列表中，则插入
	if err = dao.Task.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		taskId := in.Id
		if !g.IsEmpty(status) && status.Enum == consts.TaskDone {
			in.FinishedTime = gtime.Date()
			fmt.Println("完成时间", in.FinishedTime)
		}
		if _, err := dao.Task.Ctx(ctx).Data(in).OmitEmptyData().WherePri(taskId).Update(); err != nil {
			return err
		}
		// 若任务处理人变更，且不再任务参与人列表中，则插入
		if !g.IsEmpty(in.HandlerId) {
			if count, _ := service.Task().FindTaskActor(ctx, model_task.FindTaskActorInput{
				TaskId: taskId,
				UserId: *in.HandlerId,
			}); count < 1 {
				if err = service.Task().AddTaskActor(ctx, model_task.AddTaskActorInput{
					TaskId:  taskId,
					UserIds: []string{*in.HandlerId},
				}); err != nil {
					return err
				}
			}
		}
		return nil
	}); err != nil {
		return err
	}
	// 异步生成日志信息
	go func() {
		err = s.GenerateTaskUpdateLog(service.BizCtx().Get(ctx).User.Uid, in, *task)
		fmt.Println(err)
	}()
	return
}
