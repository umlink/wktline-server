// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskActor is the golang structure for table task_actor.
type TaskActor struct {
	Id         string      `json:"id"         orm:"id"          description:""`
	TaskId     string      `json:"taskId"     orm:"task_id"     description:"任务 id"`
	UserId     string      `json:"userId"     orm:"user_id"     description:"用户 id"`
	OperatorId string      `json:"operatorId" orm:"operator_id" description:"操作者 id"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:""`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:""`
}
