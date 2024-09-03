// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskOperationLog is the golang structure for table task_operation_log.
type TaskOperationLog struct {
	Id        string      `json:"id"        orm:"id"         description:"日志 ld"`
	TaskId    string      `json:"taskId"    orm:"task_id"    description:"任务 id"`
	UserId    string      `json:"userId"    orm:"user_id"    description:"用户 id"`
	Type      string      `json:"type"      orm:"type"       description:"日志类型"`
	Name      string      `json:"name"      orm:"name"       description:"日志名"`
	Content   string      `json:"content"   orm:"content"    description:"日志内容"`
	Desc      string      `json:"desc"      orm:"desc"       description:"其它信息"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
