// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskActor is the golang structure of table task_actor for DAO operations like Where/Data.
type TaskActor struct {
	g.Meta     `orm:"table:task_actor, do:true"`
	Id         interface{} //
	TaskId     interface{} // 任务 id
	UserId     interface{} // 用户 id
	OperatorId interface{} // 操作者 id
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
