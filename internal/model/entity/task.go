// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Task is the golang structure for table task.
type Task struct {
	Id           string      `json:"id"           orm:"id"            description:"任务 id"`
	Name         string      `json:"name"         orm:"name"          description:"任务名"`
	Description  string      `json:"description"  orm:"description"   description:"任务描述，没住"`
	ProjectId    string      `json:"projectId"    orm:"project_id"    description:"项目 id"`
	HandlerId    string      `json:"handlerId"    orm:"handler_id"    description:"任务处理者"`
	GroupId      string      `json:"groupId"      orm:"group_id"      description:"任务分组/迭代 id"`
	StatusId     string      `json:"statusId"     orm:"status_id"     description:"任务状态"`
	CreatorId    string      `json:"creatorId"    orm:"creator_id"    description:"任务创建者 id"`
	ParentId     string      `json:"parentId"     orm:"parent_id"     description:"父任务 id"`
	TypeId       string      `json:"typeId"       orm:"type_id"       description:"任务类型 id"`
	Priority     string      `json:"priority"     orm:"priority"      description:"任务优先级:P0|P1|P2|P3"`
	Progress     int         `json:"progress"     orm:"progress"      description:"进度"`
	PlanHour     float64     `json:"planHour"     orm:"plan_hour"     description:"计划工时（小时）"`
	LaborHour    float64     `json:"laborHour"    orm:"labor_hour"    description:"总工时（小时）"`
	StartTime    *gtime.Time `json:"startTime"    orm:"start_time"    description:"预计开始时间"`
	EndTime      *gtime.Time `json:"endTime"      orm:"end_time"      description:"预计结束时间"`
	FinishedTime *gtime.Time `json:"finishedTime" orm:"finished_time" description:"任务实际完成时间"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:""`
}
