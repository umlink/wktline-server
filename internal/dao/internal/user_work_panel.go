// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserWorkPanelDao is the data access object for table user_work_panel.
type UserWorkPanelDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns UserWorkPanelColumns // columns contains all the column names of Table for convenient usage.
}

// UserWorkPanelColumns defines and stores column names for table user_work_panel.
type UserWorkPanelColumns struct {
	Id         string // id
	WorkmateId string // 同事 id
	UserId     string // 用户 id(当前用户)
	CreatedAt  string //
	UpdatedAt  string //
}

// userWorkPanelColumns holds the columns for table user_work_panel.
var userWorkPanelColumns = UserWorkPanelColumns{
	Id:         "id",
	WorkmateId: "workmate_id",
	UserId:     "user_id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewUserWorkPanelDao creates and returns a new DAO object for table data access.
func NewUserWorkPanelDao() *UserWorkPanelDao {
	return &UserWorkPanelDao{
		group:   "default",
		table:   "user_work_panel",
		columns: userWorkPanelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserWorkPanelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserWorkPanelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserWorkPanelDao) Columns() UserWorkPanelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserWorkPanelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserWorkPanelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserWorkPanelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
