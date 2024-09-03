// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskLaborHour is the golang structure of table task_labor_hour for DAO operations like Where/Data.
type TaskLaborHour struct {
	g.Meta      `orm:"table:task_labor_hour, do:true"`
	Id          interface{} //
	TaskId      interface{} // 任务 id
	ProjectId   interface{} // 项目 id
	UserId      interface{} // 添加工时用户 id
	Date        *gtime.Time // 工作内容所在日期
	Hour        interface{} // 工时（单位：小时）
	Description interface{} // 工作描述
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
