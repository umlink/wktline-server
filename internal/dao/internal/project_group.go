// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProjectGroupDao is the data access object for table project_group.
type ProjectGroupDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ProjectGroupColumns // columns contains all the column names of Table for convenient usage.
}

// ProjectGroupColumns defines and stores column names for table project_group.
type ProjectGroupColumns struct {
	Id          string // 项目分组 id
	Name        string // 项目分组 id
	Description string // 项目分组描述/简介
	OperatorId  string // 操作者 id
	CreatedAt   string //
	UpdatedAt   string //
}

// projectGroupColumns holds the columns for table project_group.
var projectGroupColumns = ProjectGroupColumns{
	Id:          "id",
	Name:        "name",
	Description: "description",
	OperatorId:  "operator_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewProjectGroupDao creates and returns a new DAO object for table data access.
func NewProjectGroupDao() *ProjectGroupDao {
	return &ProjectGroupDao{
		group:   "default",
		table:   "project_group",
		columns: projectGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProjectGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProjectGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProjectGroupDao) Columns() ProjectGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProjectGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProjectGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProjectGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
