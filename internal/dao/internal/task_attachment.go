// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskAttachmentDao is the data access object for table task_attachment.
type TaskAttachmentDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns TaskAttachmentColumns // columns contains all the column names of Table for convenient usage.
}

// TaskAttachmentColumns defines and stores column names for table task_attachment.
type TaskAttachmentColumns struct {
	Id         string // id
	ProjectId  string // 项目 id
	TaskId     string // 任务 id
	ResourceId string // 资源 id
	CreatorId  string // 创建者 id
	CreatedAt  string //
	UpdatedAt  string //
}

// taskAttachmentColumns holds the columns for table task_attachment.
var taskAttachmentColumns = TaskAttachmentColumns{
	Id:         "id",
	ProjectId:  "project_id",
	TaskId:     "task_id",
	ResourceId: "resource_id",
	CreatorId:  "creator_id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewTaskAttachmentDao creates and returns a new DAO object for table data access.
func NewTaskAttachmentDao() *TaskAttachmentDao {
	return &TaskAttachmentDao{
		group:   "default",
		table:   "task_attachment",
		columns: taskAttachmentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskAttachmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskAttachmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskAttachmentDao) Columns() TaskAttachmentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskAttachmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskAttachmentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskAttachmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
