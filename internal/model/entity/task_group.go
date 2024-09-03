// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskGroup is the golang structure for table task_group.
type TaskGroup struct {
	Id          string      `json:"id"          orm:"id"          description:"任务分组 id"`
	ProjectId   string      `json:"projectId"   orm:"project_id"  description:"项目 id"`
	GroupName   string      `json:"groupName"   orm:"group_name"  description:"任务分组名（迭代）"`
	Description string      `json:"description" orm:"description" description:"任务分组描述，可用于一个迭代描述"`
	OperatorId  string      `json:"operatorId"  orm:"operator_id" description:"操作用户 id"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:""`
}
