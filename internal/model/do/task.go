// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Task is the golang structure of table task for DAO operations like Where/Data.
type Task struct {
	g.Meta       `orm:"table:task, do:true"`
	Id           interface{} // 任务 id
	Name         interface{} // 任务名
	Description  interface{} // 任务描述，没住
	ProjectId    interface{} // 项目 id
	HandlerId    interface{} // 任务处理者
	GroupId      interface{} // 任务分组/迭代 id
	StatusId     interface{} // 任务状态
	CreatorId    interface{} // 任务创建者 id
	ParentId     interface{} // 父任务 id
	TypeId       interface{} // 任务类型 id
	Priority     interface{} // 任务优先级:P0|P1|P2|P3
	Progress     interface{} // 进度
	PlanHour     interface{} // 计划工时（小时）
	LaborHour    interface{} // 总工时（小时）
	StartTime    *gtime.Time // 预计开始时间
	EndTime      *gtime.Time // 预计结束时间
	FinishedTime *gtime.Time // 任务实际完成时间
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
