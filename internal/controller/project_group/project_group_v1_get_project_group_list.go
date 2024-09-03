package project_group

import (
	"context"

	"wktline-server/api/project_group/v1"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"
)

func (c *ControllerV1) GetProjectGroupList(ctx context.Context, req *v1.GetProjectGroupListReq) (res *v1.GetProjectGroupListRes, err error) {
	out, err := service.Project().GetGroupList(ctx, model_project.GetGroupListInput{
		Keywords: req.Keywords,
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
	})
	return &v1.GetProjectGroupListRes{
		List:     out.List,
		Total:    out.Total,
		PageSize: out.PageSize,
		PageNo:   out.PageNo,
	}, err
}
