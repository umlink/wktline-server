// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskOperationLog is the golang structure of table task_operation_log for DAO operations like Where/Data.
type TaskOperationLog struct {
	g.Meta    `orm:"table:task_operation_log, do:true"`
	Id        interface{} // 日志 ld
	TaskId    interface{} // 任务 id
	UserId    interface{} // 用户 id
	Type      interface{} // 日志类型
	Name      interface{} // 日志名
	Content   interface{} // 日志内容
	Desc      interface{} // 其它信息
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
