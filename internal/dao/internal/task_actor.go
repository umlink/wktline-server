// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskActorDao is the data access object for table task_actor.
type TaskActorDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns TaskActorColumns // columns contains all the column names of Table for convenient usage.
}

// TaskActorColumns defines and stores column names for table task_actor.
type TaskActorColumns struct {
	Id         string //
	TaskId     string // 任务 id
	UserId     string // 用户 id
	OperatorId string // 操作者 id
	CreatedAt  string //
	UpdatedAt  string //
}

// taskActorColumns holds the columns for table task_actor.
var taskActorColumns = TaskActorColumns{
	Id:         "id",
	TaskId:     "task_id",
	UserId:     "user_id",
	OperatorId: "operator_id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewTaskActorDao creates and returns a new DAO object for table data access.
func NewTaskActorDao() *TaskActorDao {
	return &TaskActorDao{
		group:   "default",
		table:   "task_actor",
		columns: taskActorColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskActorDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskActorDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskActorDao) Columns() TaskActorColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskActorDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskActorDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskActorDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
