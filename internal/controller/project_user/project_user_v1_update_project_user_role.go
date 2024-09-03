package project_user

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"

	"wktline-server/api/project_user/v1"
	"wktline-server/internal/consts"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"
)

func (c *ControllerV1) UpdateProjectUserRole(ctx context.Context, req *v1.UpdateProjectUserRoleReq) (res *v1.UpdateProjectUserRoleRes, err error) {
	out, err := service.Project().GetProjectUserDetail(ctx, model_project.GetProjectUserDetailInput{
		ProjectId: req.ProjectId,
		UserId:    req.UserId,
	})
	if err != nil {
		return nil, err
	}
	if out.Role == consts.SuperAdmin {
		return nil, gerror.New("当前用户为项目负责人，请刷新后重试")
	}
	err = service.Project().UpdateProjectUserRole(ctx, model_project.UpdateProjectUserRoleInput{
		UserId:    req.UserId,
		ProjectId: req.ProjectId,
		Role:      req.Role,
	})
	return nil, err
}
