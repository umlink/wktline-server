// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskStatus is the golang structure of table task_status for DAO operations like Where/Data.
type TaskStatus struct {
	g.Meta     `orm:"table:task_status, do:true"`
	Id         interface{} // 状态 id
	ProjectId  interface{} // 项目 id
	Name       interface{} // 状态名
	Enum       interface{} // 枚举值，用作任务状态的筛选
	Color      interface{} // 颜色
	Sort       interface{} // 排序
	OperatorId interface{} // 操作这 id
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
	Default    interface{} //
}
