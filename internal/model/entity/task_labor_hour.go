// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskLaborHour is the golang structure for table task_labor_hour.
type TaskLaborHour struct {
	Id          string      `json:"id"          orm:"id"          description:""`
	TaskId      string      `json:"taskId"      orm:"task_id"     description:"任务 id"`
	ProjectId   string      `json:"projectId"   orm:"project_id"  description:"项目 id"`
	UserId      string      `json:"userId"      orm:"user_id"     description:"添加工时用户 id"`
	Date        *gtime.Time `json:"date"        orm:"date"        description:"工作内容所在日期"`
	Hour        float64     `json:"hour"        orm:"hour"        description:"工时（单位：小时）"`
	Description string      `json:"description" orm:"description" description:"工作描述"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:""`
}
