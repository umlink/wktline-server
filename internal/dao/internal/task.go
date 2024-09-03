// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskDao is the data access object for table task.
type TaskDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns TaskColumns // columns contains all the column names of Table for convenient usage.
}

// TaskColumns defines and stores column names for table task.
type TaskColumns struct {
	Id           string // 任务 id
	Name         string // 任务名
	Description  string // 任务描述，没住
	ProjectId    string // 项目 id
	HandlerId    string // 任务处理者
	GroupId      string // 任务分组/迭代 id
	StatusId     string // 任务状态
	CreatorId    string // 任务创建者 id
	ParentId     string // 父任务 id
	TypeId       string // 任务类型 id
	Priority     string // 任务优先级:P0|P1|P2|P3
	Progress     string // 进度
	PlanHour     string // 计划工时（小时）
	LaborHour    string // 总工时（小时）
	StartTime    string // 预计开始时间
	EndTime      string // 预计结束时间
	FinishedTime string // 任务实际完成时间
	CreatedAt    string //
	UpdatedAt    string //
}

// taskColumns holds the columns for table task.
var taskColumns = TaskColumns{
	Id:           "id",
	Name:         "name",
	Description:  "description",
	ProjectId:    "project_id",
	HandlerId:    "handler_id",
	GroupId:      "group_id",
	StatusId:     "status_id",
	CreatorId:    "creator_id",
	ParentId:     "parent_id",
	TypeId:       "type_id",
	Priority:     "priority",
	Progress:     "progress",
	PlanHour:     "plan_hour",
	LaborHour:    "labor_hour",
	StartTime:    "start_time",
	EndTime:      "end_time",
	FinishedTime: "finished_time",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewTaskDao creates and returns a new DAO object for table data access.
func NewTaskDao() *TaskDao {
	return &TaskDao{
		group:   "default",
		table:   "task",
		columns: taskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskDao) Columns() TaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
