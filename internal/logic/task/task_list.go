package task

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sync"
	v1 "wktline-server/api/task/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (s *sTask) GetTaskListByInterval(ctx context.Context, in model_task.GetTaskListByIntervalInput) (res *model_task.GetTaskListByIntervalOutput, err error) {
	var out model_task.GetTaskListByIntervalOutput
	d := g.Model("task t")
	d.LeftJoin("task_status ts", "ts.id=t.status_id").
		Fields("t.id,t.name,t.project_id,t.start_time,t.end_time,ts.name as statusName,ts.enum as statusEnum,ts.color as statusColor").
		Where("t.handler_id", service.BizCtx().Get(ctx).User.Uid).
		Where(d.Builder().
			WhereBetween("t.start_time", in.StartTime, in.EndTime).
			WhereOrBetween("t.end_time", in.StartTime, in.EndTime)).
		Group("t.id").
		Order("t.id")

	if err := d.Scan(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *sTask) GetTaskList(in model_task.GetTaskListInput) (res *model_task.GetTaskListOutput, err error) {
	out := &model_task.GetTaskListOutput{
		PageNo:   in.PageNo,
		PageSize: in.PageSize,
	}
	d := g.Model("task").OmitEmptyWhere().Where(in)
	if !g.IsEmpty(in.Keywords) {
		likePattern := `%` + *in.Keywords + `%`
		d = d.WhereLike("task.name", likePattern)
	}
	out.List = make([]*v1.TaskDetailItem, 0)
	if !g.IsEmpty(in.SortMode) {
		var column = "task.updated_at"
		if !g.IsEmpty(in.SortMode.SortKey) {
			column = "task." + in.SortMode.SortKey
		}
		if in.SortMode.Mode == "Desc" {
			d = d.OrderDesc(column)
		} else {
			d = d.OrderAsc(column)
		}
	}
	cd := d.Clone()

	d.LeftJoin("user hu", "hu.id=task.handler_id").
		LeftJoin("task_type tt", "tt.id=task.type_id").
		LeftJoin("task_status ts", "ts.id=task.status_id").
		LeftJoin("task_group tg", "tg.id=task.group_id").
		LeftJoin("user cu", "cu.id=task.creator_id").
		Fields("task.*").
		Fields("hu.username as handlerName,hu.avatar as handlerAvatar").
		Fields("cu.username as creatorName,cu.avatar as creatorAvatar").
		Fields("tt.name as typeName,tt.color as typeColor").
		Fields("ts.name as statusName,ts.color as statusColor,ts.enum as statusEnum").
		Fields("tg.group_name as groupName").
		OrderDesc("task.id")

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		out.Total, err = cd.Count()
	}()
	go func() {
		defer wg.Done()
		err = d.Page(in.PageNo, in.PageSize).Scan(&out.List)
	}()
	wg.Wait()
	return out, nil
}
