// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package project_user

import (
	"context"
	
	"wktline-server/api/project_user/v1"
)

type IProjectUserV1 interface {
	GetProjectUser(ctx context.Context, req *v1.GetProjectUserReq) (res *v1.GetProjectUserRes, err error)
	AddProjectUserByIds(ctx context.Context, req *v1.AddProjectUserByIdsReq) (res *v1.AddProjectUserByIdsRes, err error)
	DelProjectUserByIds(ctx context.Context, req *v1.DelProjectUserByIdsReq) (res *v1.DelProjectUserByIdsRes, err error)
	UpdateProjectUserRole(ctx context.Context, req *v1.UpdateProjectUserRoleReq) (res *v1.UpdateProjectUserRoleRes, err error)
}


