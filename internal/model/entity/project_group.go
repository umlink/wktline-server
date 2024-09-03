// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProjectGroup is the golang structure for table project_group.
type ProjectGroup struct {
	Id          string      `json:"id"          orm:"id"          description:"项目分组 id"`
	Name        string      `json:"name"        orm:"name"        description:"项目分组 id"`
	Description string      `json:"description" orm:"description" description:"项目分组描述/简介"`
	OperatorId  string      `json:"operatorId"  orm:"operator_id" description:"操作者 id"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:""`
}
