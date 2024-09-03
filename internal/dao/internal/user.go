// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for table user.
type UserDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns UserColumns // columns contains all the column names of Table for convenient usage.
}

// UserColumns defines and stores column names for table user.
type UserColumns struct {
	Id           string // id
	Avatar       string // 用户头像
	Username     string // 用户名
	Nickname     string // 用户昵称
	Password     string // 密码
	Age          string // 用户年龄
	Six          string // 用户性别 默认：0， 1：男，2：女，3：未知
	Status       string // 用户状态,默认1 1：正常，0：禁用，2：xxx，3：xxx
	Phone        string // 手机号码
	EmpNumber    string // 工号
	Email        string // 邮箱
	Role         string // 用户角色 (系统级)USER | ADMIN | SUPER_ADMIN
	DepartmentId string // 部门 id
	IsUpdate     string // 是否存在更新
	CreatedAt    string //
	UpdatedAt    string //
}

// userColumns holds the columns for table user.
var userColumns = UserColumns{
	Id:           "id",
	Avatar:       "avatar",
	Username:     "username",
	Nickname:     "nickname",
	Password:     "password",
	Age:          "age",
	Six:          "six",
	Status:       "status",
	Phone:        "phone",
	EmpNumber:    "emp_number",
	Email:        "email",
	Role:         "role",
	DepartmentId: "department_id",
	IsUpdate:     "is_update",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao() *UserDao {
	return &UserDao{
		group:   "default",
		table:   "user",
		columns: userColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
