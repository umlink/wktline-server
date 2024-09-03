// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package project_invite

import (
	"context"

	"wktline-server/api/project_invite/v1"
)

type IProjectInviteV1 interface {
	GenProjectInviteCode(ctx context.Context, req *v1.GenProjectInviteCodeReq) (res *v1.GenProjectInviteCodeRes, err error)
	GetProjectInviteInfo(ctx context.Context, req *v1.GetProjectInviteInfoReq) (res *v1.GetProjectInviteInfoRes, err error)
	InJoinInviteProject(ctx context.Context, req *v1.InJoinInviteProjectReq) (res *v1.InJoinInviteProjectRes, err error)
	DelInviteCode(ctx context.Context, req *v1.DelInviteCodeReq) (res *v1.DelInviteCodeRes, err error)
	GetProjectMyInviteList(ctx context.Context, req *v1.GetProjectMyInviteListReq) (res *v1.GetProjectMyInviteListRes, err error)
}
