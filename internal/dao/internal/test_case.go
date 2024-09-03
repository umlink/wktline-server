// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TestCaseDao is the data access object for table test_case.
type TestCaseDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns TestCaseColumns // columns contains all the column names of Table for convenient usage.
}

// TestCaseColumns defines and stores column names for table test_case.
type TestCaseColumns struct {
	Id          string // id
	Name        string // 测试用例名称
	Description string // 说明
	Value       string // 测试用例数据
	Progress    string // 测试进度（1-100）
	ProjectId   string // 项目 id
	CreatorId   string // 创建人
	ActorIds    string // 参与人ids
	Status      string // 状态(草稿: DRAFT 发布:PUBLISHED 测试中:TESTING 测试完成:DONE)
	EditorId    string // 正在编辑的用户 id，lock
	CreatedAt   string //
	DeletedAt   string // 删除时间
}

// testCaseColumns holds the columns for table test_case.
var testCaseColumns = TestCaseColumns{
	Id:          "id",
	Name:        "name",
	Description: "description",
	Value:       "value",
	Progress:    "progress",
	ProjectId:   "project_id",
	CreatorId:   "creator_id",
	ActorIds:    "actor_ids",
	Status:      "status",
	EditorId:    "editor_id",
	CreatedAt:   "created_at",
	DeletedAt:   "deleted_at",
}

// NewTestCaseDao creates and returns a new DAO object for table data access.
func NewTestCaseDao() *TestCaseDao {
	return &TestCaseDao{
		group:   "default",
		table:   "test_case",
		columns: testCaseColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TestCaseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TestCaseDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TestCaseDao) Columns() TestCaseColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TestCaseDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TestCaseDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TestCaseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
