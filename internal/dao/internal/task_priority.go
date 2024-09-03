// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskPriorityDao is the data access object for table task_priority.
type TaskPriorityDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns TaskPriorityColumns // columns contains all the column names of Table for convenient usage.
}

// TaskPriorityColumns defines and stores column names for table task_priority.
type TaskPriorityColumns struct {
	Id        string // 优先级
	Name      string // 优先级名
	Color     string // 颜色
	CreatedAt string //
	UpdatedAt string //
}

// taskPriorityColumns holds the columns for table task_priority.
var taskPriorityColumns = TaskPriorityColumns{
	Id:        "id",
	Name:      "name",
	Color:     "color",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewTaskPriorityDao creates and returns a new DAO object for table data access.
func NewTaskPriorityDao() *TaskPriorityDao {
	return &TaskPriorityDao{
		group:   "default",
		table:   "task_priority",
		columns: taskPriorityColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskPriorityDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskPriorityDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskPriorityDao) Columns() TaskPriorityColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskPriorityDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskPriorityDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskPriorityDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
