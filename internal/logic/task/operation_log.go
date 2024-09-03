package task

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"sort"
	"wktline-server/internal/model"

	"wktline-server/utility"

	"github.com/gogf/gf/v2/frame/g"
	v1 "wktline-server/api/task_operation_log/v1"
	"wktline-server/internal/dao"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func formatMsgContent(title string, preText string, lastText string) (ret string) {
	if g.IsEmpty(preText) {
		return fmt.Sprintf("%s变更为 <span class='last-text'>%s</span>", title, lastText)
	}
	return fmt.Sprintf("%s由 [<span class='pre-text'>%s</span>] 变更为 <span class='last-text'>%s</span>", title, preText, lastText)
}

func (s *sTask) GenerateTaskUpdateLog(userId string, in model_task.UpdateTaskInput, task model_task.GetTaskDetailOutput) (err error) {
	params := model_task.CreateOperationLogInput{}
	params.TaskId = in.Id
	params.UserId = userId
	// 组装消息
	msgParams := model.AddMessageInput{
		SenderId: userId,
		TaskId:   in.Id,
	}
	if !g.IsEmpty(in.Name) {
		content := formatMsgContent("任务标题", task.Name, *in.Name)
		params.Id = utility.GenUniIdByGuid()
		params.Name = "更新任务标题"
		params.Content = content
		params.Type = v1.DynamicTaskName
		if err = s.CreateTaskOperationLog(params); err != nil {
			return err
		}
		msgParams.Title = "任务标题变更通知"
		msgParams.Content = content
		if err = service.Message().CreateTaskMessage(msgParams); err != nil {
			return err
		}
	}
	if !g.IsEmpty(in.HandlerId) {
		userName, err := g.Model("user").Fields("username").WherePri(*in.HandlerId).Value()
		if err != nil {
			return err
		}
		tName := ""
		if !g.IsEmpty(task.HandlerId) {
			tName = task.HandlerName
		}
		content := formatMsgContent("任务负责人", tName, gconv.String(userName))
		if err == nil {
			params.Id = utility.GenUniIdByGuid()
			params.Name = "更新任务负责人"
			params.Content = content
			params.Type = v1.DynamicTaskHandler
			if err = s.CreateTaskOperationLog(params); err != nil {
				return err
			}
		}
		msgParams.Title = "任务负责人变更通知"
		msgParams.Content = content
		if err = service.Message().CreateTaskMessage(msgParams); err != nil {
			return err
		}
	}
	if !g.IsEmpty(in.StatusId) {
		statusName, err := g.Model("task_status").Fields("name").WherePri(*in.StatusId).Value()
		tStatus := ""
		if !g.IsEmpty(task.StatusId) {
			tStatus = task.StatusName
		}
		content := formatMsgContent("任务状态", tStatus, gconv.String(statusName))
		if err == nil {
			params.Id = utility.GenUniIdByGuid()
			params.Name = "更新任务状态"
			params.Content = content
			params.Type = v1.DynamicTaskStatus
			if err = s.CreateTaskOperationLog(params); err != nil {
				return err
			}
		}
		msgParams.Title = "任务状态变更通知"
		msgParams.Content = content
		if err = service.Message().CreateTaskMessage(msgParams); err != nil {
			return err
		}
	}
	if !g.IsEmpty(in.Priority) {
		tPriority := ""
		if !g.IsEmpty(task.Priority) {
			tPriority = task.Priority
		}
		content := formatMsgContent("任务优先级", tPriority, *in.Priority)
		if err == nil {
			params.Id = utility.GenUniIdByGuid()
			params.Name = "更新任务优先级"
			params.Content = content
			params.Type = v1.DynamicTaskPriority
			if err = s.CreateTaskOperationLog(params); err != nil {
				return err
			}
		}
		msgParams.Title = "任务优先级变更通知"
		msgParams.Content = content
		if err = service.Message().CreateTaskMessage(msgParams); err != nil {
			return err
		}
	}
	if !g.IsEmpty(in.TypeId) {
		typeName, err := g.Model("task_type").Fields("name").WherePri(*in.TypeId).Value()
		tType := ""
		if !g.IsEmpty(task.TypeId) {
			tType = task.TypeName
		}
		content := formatMsgContent("任务类型", tType, gconv.String(typeName))
		if err == nil {
			params.Id = utility.GenUniIdByGuid()
			params.Name = "更新任务类型"
			params.Content = content
			params.Type = v1.DynamicTaskType
			if err = s.CreateTaskOperationLog(params); err != nil {
				return err
			}
		}
		msgParams.Title = "任务类型变更通知"
		msgParams.Content = content
		if err = service.Message().CreateTaskMessage(msgParams); err != nil {
			return err
		}
	}
	if !g.IsEmpty(in.StartTime) {
		content := formatMsgContent("任务开始时间", task.StartTime, *in.StartTime)
		params.Id = utility.GenUniIdByGuid()
		params.Name = "更新任务开始时间"
		params.Content = content
		params.Type = v1.DynamicTaskTime
		if err = s.CreateTaskOperationLog(params); err != nil {
			return err
		}
		msgParams.Title = "任务开始时间变更通知"
		msgParams.Content = content
		if err = service.Message().CreateTaskMessage(msgParams); err != nil {
			return err
		}
	}
	if !g.IsEmpty(in.EndTime) {
		content := formatMsgContent("任务结束时间", task.EndTime, *in.EndTime)
		params.Id = utility.GenUniIdByGuid()
		params.Name = "更新任务结束时间"
		params.Content = content
		params.Type = v1.DynamicTaskTime
		if err = s.CreateTaskOperationLog(params); err != nil {
			return err
		}
		msgParams.Title = "任务结束时间变更通知"
		msgParams.Content = content
		if err = service.Message().CreateTaskMessage(msgParams); err != nil {
			return err
		}
	}
	if !g.IsEmpty(in.Description) {
		content := formatMsgContent("任务内容", task.Description, *in.Description)
		params.Id = utility.GenUniIdByGuid()
		params.Name = "更新任务内容"
		params.Content = content
		params.Type = v1.DynamicTaskDescription
		if err = s.CreateTaskOperationLog(params); err != nil {
			return err
		}
		msgParams.Title = "任务结束时间变更通知"
		msgParams.Content = content
		if err = service.Message().CreateTaskMessage(msgParams); err != nil {
			return err
		}
	}
	if !g.IsEmpty(in.GroupId) {
		groupName, err := g.Model("task_group").Fields("group_name").WherePri(*in.GroupId).Value()
		tGroup := ""
		if !g.IsEmpty(task.GroupName) {
			tGroup = task.GroupName
		}
		content := formatMsgContent("任务分组", tGroup, gconv.String(groupName))
		if err == nil {
			params.Id = utility.GenUniIdByGuid()
			params.Name = "更新了任务分组"
			params.Content = content
			params.Type = v1.DynamicTaskGroup
			if err = s.CreateTaskOperationLog(params); err != nil {
				return err
			}
		}
		msgParams.Title = "任务分组变更通知"
		msgParams.Content = content
		if err = service.Message().CreateTaskMessage(msgParams); err != nil {
			return err
		}
	}
	return err
}

func (s *sTask) CreateTaskOperationLog(in model_task.CreateOperationLogInput) (err error) {
	_, err = g.Model("task_operation_log").Data(in).Insert()
	return err
}

func (s *sTask) DelTaskOperationLog(ctx context.Context, in model_task.DelOperationLogInput) (err error) {
	_, err = dao.TaskOperationLog.Ctx(ctx).WherePri(in.Id).Delete()
	return err
}

func (s *sTask) GetTaskOperationLogList(ctx context.Context, in model_task.GetOperationLogListInput) (res *model_task.GetOperationLogListOutput, err error) {
	out := &model_task.GetOperationLogListOutput{
		PageSize: in.PageSize,
		PageNo:   in.PageNo,
	}
	d := g.Model("task_operation_log ol").Where("ol.task_id", in.TaskId)
	if !g.IsEmpty(in.Type) {
		d = d.Where("ol.type", in.Type)
	}
	d = d.InnerJoin("user u", "u.id=ol.user_id")
	d = d.Fields("ol.id,u.id as user_id,u.username,u.avatar,ol.type,ol.name,ol.content,ol.desc,ol.created_at").OrderDesc("ol.created_at")

	err = d.Page(in.PageNo, in.PageSize).ScanAndCount(&out.List, &out.Total, false)

	// 倒序返回
	sort.Slice(out.List, func(i, j int) bool {
		return out.List[i].CreatedAt < out.List[j].CreatedAt
	})
	return out, err
}
