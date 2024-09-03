// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package project

import (
	"context"

	"wktline-server/api/project/v1"
)

type IProjectV1 interface {
	GetProjectInfo(ctx context.Context, req *v1.GetProjectInfoReq) (res *v1.GetProjectInfoRes, err error)
	CreateProject(ctx context.Context, req *v1.CreateProjectReq) (res *v1.CreateProjectRes, err error)
	DelProject(ctx context.Context, req *v1.DelProjectReq) (res *v1.DelProjectRes, err error)
	UpdateProject(ctx context.Context, req *v1.UpdateProjectReq) (res *v1.UpdateProjectRes, err error)
	GetProjectList(ctx context.Context, req *v1.GetProjectListReq) (res *v1.GetProjectListRes, err error)
}
