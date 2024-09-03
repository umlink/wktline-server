// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskStatusDao is the data access object for table task_status.
type TaskStatusDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns TaskStatusColumns // columns contains all the column names of Table for convenient usage.
}

// TaskStatusColumns defines and stores column names for table task_status.
type TaskStatusColumns struct {
	Id         string // 状态 id
	ProjectId  string // 项目 id
	Name       string // 状态名
	Enum       string // 枚举值，用作任务状态的筛选
	Color      string // 颜色
	Sort       string // 排序
	OperatorId string // 操作这 id
	CreatedAt  string //
	UpdatedAt  string //
	Default    string //
}

// taskStatusColumns holds the columns for table task_status.
var taskStatusColumns = TaskStatusColumns{
	Id:         "id",
	ProjectId:  "project_id",
	Name:       "name",
	Enum:       "enum",
	Color:      "color",
	Sort:       "sort",
	OperatorId: "operator_id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	Default:    "default",
}

// NewTaskStatusDao creates and returns a new DAO object for table data access.
func NewTaskStatusDao() *TaskStatusDao {
	return &TaskStatusDao{
		group:   "default",
		table:   "task_status",
		columns: taskStatusColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskStatusDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskStatusDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskStatusDao) Columns() TaskStatusColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskStatusDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskStatusDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskStatusDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
