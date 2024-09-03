// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskType is the golang structure for table task_type.
type TaskType struct {
	Id         string      `json:"id"         orm:"id"          description:""`
	ProjectId  string      `json:"projectId"  orm:"project_id"  description:"项目 id"`
	Name       string      `json:"name"       orm:"name"        description:"类型名称"`
	Color      string      `json:"color"      orm:"color"       description:"类型颜色"`
	OperatorId string      `json:"operatorId" orm:"operator_id" description:"操作者 id"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:""`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:""`
}
