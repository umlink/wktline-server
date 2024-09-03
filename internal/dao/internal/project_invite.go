// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProjectInviteDao is the data access object for table project_invite.
type ProjectInviteDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns ProjectInviteColumns // columns contains all the column names of Table for convenient usage.
}

// ProjectInviteColumns defines and stores column names for table project_invite.
type ProjectInviteColumns struct {
	Id             string // 邀请 code
	ProjectId      string // 项目 id
	InviterId      string // 发起邀请的用户 id
	Deadline       string // 生效截止时间
	MaxInviteCount string // z
	JoinedCount    string // 已加入人数
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	DeletedAt      string // 删除时间
}

// projectInviteColumns holds the columns for table project_invite.
var projectInviteColumns = ProjectInviteColumns{
	Id:             "id",
	ProjectId:      "project_id",
	InviterId:      "inviter_id",
	Deadline:       "deadline",
	MaxInviteCount: "max_invite_count",
	JoinedCount:    "joined_count",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewProjectInviteDao creates and returns a new DAO object for table data access.
func NewProjectInviteDao() *ProjectInviteDao {
	return &ProjectInviteDao{
		group:   "default",
		table:   "project_invite",
		columns: projectInviteColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProjectInviteDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProjectInviteDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProjectInviteDao) Columns() ProjectInviteColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProjectInviteDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProjectInviteDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProjectInviteDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
