package task

import (
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"wktline-server/internal/model/model_task"
)

func (s *sTask) GetChildTaskList(in model_task.GetChildTaskListInput) (res *model_task.GetTaskListOutput, err error) {
	out := &model_task.GetTaskListOutput{
		PageNo:   in.PageNo,
		PageSize: in.PageSize,
	}
	d := g.Model("task").OmitEmptyWhere().Where(in)
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
	return out, err
}
