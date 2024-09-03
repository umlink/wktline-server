// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskStatus is the golang structure for table task_status.
type TaskStatus struct {
	Id         string      `json:"id"         orm:"id"          description:"状态 id"`
	ProjectId  string      `json:"projectId"  orm:"project_id"  description:"项目 id"`
	Name       string      `json:"name"       orm:"name"        description:"状态名"`
	Enum       string      `json:"enum"       orm:"enum"        description:"枚举值，用作任务状态的筛选"`
	Color      string      `json:"color"      orm:"color"       description:"颜色"`
	Sort       int         `json:"sort"       orm:"sort"        description:"排序"`
	OperatorId string      `json:"operatorId" orm:"operator_id" description:"操作这 id"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:""`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:""`
	Default    int         `json:"default"    orm:"default"     description:""`
}
