// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskPriority is the golang structure of table task_priority for DAO operations like Where/Data.
type TaskPriority struct {
	g.Meta    `orm:"table:task_priority, do:true"`
	Id        interface{} // 优先级
	Name      interface{} // 优先级名
	Color     interface{} // 颜色
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
