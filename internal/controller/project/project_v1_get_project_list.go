package project

import (
	"context"

	"wktline-server/api/project/v1"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"
)

func (c *ControllerV1) GetProjectList(ctx context.Context, req *v1.GetProjectListReq) (res *v1.GetProjectListRes, err error) {

	out, err := service.Project().GetProjectList(ctx, model_project.GetProjectListInput{
		Keywords:  req.Keywords,
		GroupId:   req.GroupId,
		IsCreator: req.IsCreator,
		IsOwner:   req.IsOwner,
		IsJoined:  req.IsJoined,
		ShowType:  req.ShowType,
		PageNo:    req.PageNo,
		PageSize:  req.PageSize,
	})
	return &v1.GetProjectListRes{
		List:     out.List,
		PageSize: out.PageSize,
		PageNo:   out.PageNo,
		Total:    out.Total,
	}, err
}
