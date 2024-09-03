// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package project_group

import (
	"context"
	
	"wktline-server/api/project_group/v1"
)

type IProjectGroupV1 interface {
	CreateProjectGroup(ctx context.Context, req *v1.CreateProjectGroupReq) (res *v1.CreateProjectGroupRes, err error)
	DelProjectGroup(ctx context.Context, req *v1.DelProjectGroupReq) (res *v1.DelProjectGroupRes, err error)
	UpdateProjectGroup(ctx context.Context, req *v1.UpdateProjectGroupReq) (res *v1.UpdateProjectGroupRes, err error)
	GetProjectGroupList(ctx context.Context, req *v1.GetProjectGroupListReq) (res *v1.GetProjectGroupListRes, err error)
}


