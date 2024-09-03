// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProjectDao is the data access object for table project.
type ProjectDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns ProjectColumns // columns contains all the column names of Table for convenient usage.
}

// ProjectColumns defines and stores column names for table project.
type ProjectColumns struct {
	Id          string // 项目 id
	Name        string // 项目名
	Description string // 项目描述
	HeaderImg   string // 项目头图
	CreatorId   string // 创建人 id
	OwnerId     string // 项目所有者 id
	OperatorId  string // 项目操作者 id
	GroupId     string // 项目分组 id
	Status      string // 项目状态 1: 正常 2: 删除
	ShowType    string // 显示类型 PUBLIC: 公开 PRIVATE: 私有
	CreatedAt   string //
	UpdatedAt   string //
	DeletedAt   string // 删除shi
}

// projectColumns holds the columns for table project.
var projectColumns = ProjectColumns{
	Id:          "id",
	Name:        "name",
	Description: "description",
	HeaderImg:   "header_img",
	CreatorId:   "creator_id",
	OwnerId:     "owner_id",
	OperatorId:  "operator_id",
	GroupId:     "group_id",
	Status:      "status",
	ShowType:    "show_type",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewProjectDao creates and returns a new DAO object for table data access.
func NewProjectDao() *ProjectDao {
	return &ProjectDao{
		group:   "default",
		table:   "project",
		columns: projectColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProjectDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProjectDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProjectDao) Columns() ProjectColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProjectDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProjectDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProjectDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
