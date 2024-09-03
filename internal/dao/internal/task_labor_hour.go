// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskLaborHourDao is the data access object for table task_labor_hour.
type TaskLaborHourDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns TaskLaborHourColumns // columns contains all the column names of Table for convenient usage.
}

// TaskLaborHourColumns defines and stores column names for table task_labor_hour.
type TaskLaborHourColumns struct {
	Id          string //
	TaskId      string // 任务 id
	ProjectId   string // 项目 id
	UserId      string // 添加工时用户 id
	Date        string // 工作内容所在日期
	Hour        string // 工时（单位：小时）
	Description string // 工作描述
	CreatedAt   string //
	UpdatedAt   string //
}

// taskLaborHourColumns holds the columns for table task_labor_hour.
var taskLaborHourColumns = TaskLaborHourColumns{
	Id:          "id",
	TaskId:      "task_id",
	ProjectId:   "project_id",
	UserId:      "user_id",
	Date:        "date",
	Hour:        "hour",
	Description: "description",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewTaskLaborHourDao creates and returns a new DAO object for table data access.
func NewTaskLaborHourDao() *TaskLaborHourDao {
	return &TaskLaborHourDao{
		group:   "default",
		table:   "task_labor_hour",
		columns: taskLaborHourColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskLaborHourDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskLaborHourDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskLaborHourDao) Columns() TaskLaborHourColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskLaborHourDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskLaborHourDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskLaborHourDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
