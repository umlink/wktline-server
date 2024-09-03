package task

import (
	"github.com/gogf/gf/v2/frame/g"

	"wktline-server/internal/model/model_task"
)

func (s *sTask) GetTaskDetail(in model_task.GetTaskDetailInput) (res *model_task.GetTaskDetailOutput, err error) {
	if err = g.Model("task").LeftJoin("user hu", "hu.id=task.handler_id").
		LeftJoin("task_type tt", "tt.id=task.type_id").
		LeftJoin("task_status ts", "ts.id=task.status_id").
		LeftJoin("task_group tg", "tg.id=task.group_id").
		LeftJoin("user cu", "cu.id=task.creator_id").
		Fields("task.*").
		Fields("hu.username as handlerName,hu.avatar as handlerAvatar").
		Fields("cu.username as creatorName,cu.avatar as creatorAvatar").
		Fields("tt.name as typeName,tt.color as typeColor").
		Fields("ts.name as statusName,ts.color as statusColor,ts.enum as statusEnum").
		Fields("tg.group_name as groupName").WherePri(in.Id).Scan(&res); err != nil {
		return nil, err
	}
	return res, err
}
