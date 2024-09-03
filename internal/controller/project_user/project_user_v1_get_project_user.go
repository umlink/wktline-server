package project_user

import (
	"context"

	"wktline-server/api/project_user/v1"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"
)

func (c *ControllerV1) GetProjectUser(ctx context.Context, req *v1.GetProjectUserReq) (res *v1.GetProjectUserRes, err error) {
	out, err := service.Project().GetProjectUserList(ctx, model_project.GetProjectUserListInput{
		ProjectId: req.ProjectId,
		PageSize:  req.PageSize,
		Role:      req.Role,
		PageNo:    req.PageNo,
		Keywords:  req.Keywords,
	})
	return &v1.GetProjectUserRes{
		List:     out.List,
		Total:    out.Total,
		PageSize: out.PageSize,
		PageNo:   out.PageNo,
	}, err
}
