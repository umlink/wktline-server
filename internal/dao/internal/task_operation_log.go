// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskOperationLogDao is the data access object for table task_operation_log.
type TaskOperationLogDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns TaskOperationLogColumns // columns contains all the column names of Table for convenient usage.
}

// TaskOperationLogColumns defines and stores column names for table task_operation_log.
type TaskOperationLogColumns struct {
	Id        string // 日志 ld
	TaskId    string // 任务 id
	UserId    string // 用户 id
	Type      string // 日志类型
	Name      string // 日志名
	Content   string // 日志内容
	Desc      string // 其它信息
	CreatedAt string //
	UpdatedAt string //
}

// taskOperationLogColumns holds the columns for table task_operation_log.
var taskOperationLogColumns = TaskOperationLogColumns{
	Id:        "id",
	TaskId:    "task_id",
	UserId:    "user_id",
	Type:      "type",
	Name:      "name",
	Content:   "content",
	Desc:      "desc",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewTaskOperationLogDao creates and returns a new DAO object for table data access.
func NewTaskOperationLogDao() *TaskOperationLogDao {
	return &TaskOperationLogDao{
		group:   "default",
		table:   "task_operation_log",
		columns: taskOperationLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskOperationLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskOperationLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskOperationLogDao) Columns() TaskOperationLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskOperationLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskOperationLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskOperationLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
