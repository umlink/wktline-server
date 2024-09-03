package task

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "wktline-server/api/task_operation_log/v1"
	"wktline-server/internal/dao"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
	"wktline-server/utility"
)

func (s *sTask) AddTaskLaborHour(ctx context.Context, in model_task.AddTaskLaborHourInput) (err error) {
	return dao.Task.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		in.Id = utility.GenUniIdByGuid()
		var wg = sync.WaitGroup{}
		wg.Add(3)
		go func() {
			defer wg.Done()
			// 更新实际工时
			if _, e := dao.Task.Ctx(ctx).WherePri(in.TaskId).Increment("labor_hour", in.Hour); e != nil {
				err = e
			}
		}()
		go func() {
			defer wg.Done()
			// 插入工时记录
			if _, e := dao.TaskLaborHour.Ctx(ctx).Data(in).Insert(); e != nil {
				err = e
			}
		}()
		go func() {
			defer wg.Done()
			// 添加操作日志
			if e := s.CreateTaskOperationLog(model_task.CreateOperationLogInput{
				Id:      utility.GenUniIdByGuid(),
				TaskId:  in.TaskId,
				Name:    "添加实际工时",
				Type:    v1.DynamicTaskLaborTime,
				Content: gconv.String(in.Hour) + "小时",
				UserId:  service.BizCtx().Get(ctx).User.Uid,
			}); e != nil {
				err = e
			}
		}()
		wg.Wait()
		return err
	})
}

func (s *sTask) DelLaborHourById(ctx context.Context, in model_task.DelTaskLaborHourInput) (err error) {
	return dao.Task.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 更新实际工时
		if _, err = dao.Task.Ctx(ctx).WherePri(in.TaskId).Decrement("labor_hour", -in.Hour); err != nil {
			return err
		}
		// 添加操作日志
		if err != nil {
			return err
		}
		if err = s.CreateTaskOperationLog(model_task.CreateOperationLogInput{
			Id:      utility.GenUniIdByGuid(),
			TaskId:  in.TaskId,
			Name:    "删除实际工时",
			Type:    v1.DynamicTaskLaborTime,
			Content: gconv.String(in.Hour) + "小时",
			UserId:  service.BizCtx().Get(ctx).User.Uid,
		}); err != nil {
			return err
		}
		// 删除工时记录
		_, err = dao.TaskLaborHour.Ctx(ctx).WherePri(in.Id).Delete()
		return err
	})
}

func (s *sTask) UpdateTaskLaborHour(ctx context.Context, in model_task.UpdateTaskLaborHourInput) (err error) {
	// 获取旧工时
	hour, err := dao.TaskLaborHour.Ctx(ctx).Fields("hour").WherePri(in.Id).Value()
	if err != nil {
		return err
	}
	// 更新实际工时
	if _, err = dao.Task.Ctx(ctx).WherePri(in.TaskId).Increment("labor_hour", in.Hour-gconv.Float64(hour)); err != nil {
		return err
	}
	_, err = dao.TaskLaborHour.Ctx(ctx).Data(in).WherePri(in.Id).Update()
	return err
}

func (s *sTask) GetTaskLaborHourList(ctx context.Context, in model_task.GetTaskLaborHourListInput) (res *model_task.GetTaskLaborHourListOutput, err error) {
	out := &model_task.GetTaskLaborHourListOutput{
		PageNo:   in.PageNo,
		PageSize: in.PageSize,
	}
	d := g.Model("task_labor_hour lh, user u")
	d = d.Where("lh.user_id=u.id")
	d = d.Where("lh.task_id", in.TaskId).OrderDesc("lh.created_at")
	d = d.Page(in.PageNo, in.PageSize)
	d = d.Fields("lh.id,lh.date,lh.hour,lh.description,u.id as user_id,u.avatar,u.username")
	err = d.ScanAndCount(&out.List, &out.Total, false)
	return out, err
}
