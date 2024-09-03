// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskGroup is the golang structure of table task_group for DAO operations like Where/Data.
type TaskGroup struct {
	g.Meta      `orm:"table:task_group, do:true"`
	Id          interface{} // 任务分组 id
	ProjectId   interface{} // 项目 id
	GroupName   interface{} // 任务分组名（迭代）
	Description interface{} // 任务分组描述，可用于一个迭代描述
	OperatorId  interface{} // 操作用户 id
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
