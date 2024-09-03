// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskGroupDao is the data access object for table task_group.
type TaskGroupDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns TaskGroupColumns // columns contains all the column names of Table for convenient usage.
}

// TaskGroupColumns defines and stores column names for table task_group.
type TaskGroupColumns struct {
	Id          string // 任务分组 id
	ProjectId   string // 项目 id
	GroupName   string // 任务分组名（迭代）
	Description string // 任务分组描述，可用于一个迭代描述
	OperatorId  string // 操作用户 id
	CreatedAt   string //
	UpdatedAt   string //
}

// taskGroupColumns holds the columns for table task_group.
var taskGroupColumns = TaskGroupColumns{
	Id:          "id",
	ProjectId:   "project_id",
	GroupName:   "group_name",
	Description: "description",
	OperatorId:  "operator_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewTaskGroupDao creates and returns a new DAO object for table data access.
func NewTaskGroupDao() *TaskGroupDao {
	return &TaskGroupDao{
		group:   "default",
		table:   "task_group",
		columns: taskGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskGroupDao) Columns() TaskGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
