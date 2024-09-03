// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskType is the golang structure of table task_type for DAO operations like Where/Data.
type TaskType struct {
	g.Meta     `orm:"table:task_type, do:true"`
	Id         interface{} //
	ProjectId  interface{} // 项目 id
	Name       interface{} // 类型名称
	Color      interface{} // 类型颜色
	OperatorId interface{} // 操作者 id
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
