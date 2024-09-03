// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskTypeDao is the data access object for table task_type.
type TaskTypeDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns TaskTypeColumns // columns contains all the column names of Table for convenient usage.
}

// TaskTypeColumns defines and stores column names for table task_type.
type TaskTypeColumns struct {
	Id         string //
	ProjectId  string // 项目 id
	Name       string // 类型名称
	Color      string // 类型颜色
	OperatorId string // 操作者 id
	CreatedAt  string //
	UpdatedAt  string //
}

// taskTypeColumns holds the columns for table task_type.
var taskTypeColumns = TaskTypeColumns{
	Id:         "id",
	ProjectId:  "project_id",
	Name:       "name",
	Color:      "color",
	OperatorId: "operator_id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewTaskTypeDao creates and returns a new DAO object for table data access.
func NewTaskTypeDao() *TaskTypeDao {
	return &TaskTypeDao{
		group:   "default",
		table:   "task_type",
		columns: taskTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskTypeDao) Columns() TaskTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
