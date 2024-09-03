// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProjectUserDao is the data access object for table project_user.
type ProjectUserDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ProjectUserColumns // columns contains all the column names of Table for convenient usage.
}

// ProjectUserColumns defines and stores column names for table project_user.
type ProjectUserColumns struct {
	Id        string //
	ProjectId string // 项目 id
	UserId    string // 用户 id
	Role      string // 用户在项目中的角色
	CreatedAt string //
	UpdatedAt string //
}

// projectUserColumns holds the columns for table project_user.
var projectUserColumns = ProjectUserColumns{
	Id:        "id",
	ProjectId: "project_id",
	UserId:    "user_id",
	Role:      "role",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewProjectUserDao creates and returns a new DAO object for table data access.
func NewProjectUserDao() *ProjectUserDao {
	return &ProjectUserDao{
		group:   "default",
		table:   "project_user",
		columns: projectUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProjectUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProjectUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProjectUserDao) Columns() ProjectUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProjectUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProjectUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProjectUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
