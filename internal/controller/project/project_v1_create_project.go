package project

import (
	"context"
	"wktline-server/utility"

	"wktline-server/api/project/v1"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"
)

func (c *ControllerV1) CreateProject(ctx context.Context, req *v1.CreateProjectReq) (res *v1.CreateProjectRes, err error) {
	currentUserId := service.BizCtx().Get(ctx).User.Uid
	ProjectId := utility.GenUniIdByGuid()
	if err != nil {
		return nil, err
	}
	err = service.Project().CreateProject(ctx, model_project.CreateProjectInput{
		Id:          ProjectId,
		Name:        req.Name,
		Description: req.Description,
		HeaderImg:   req.HeaderImg,
		ShowType:    req.ShowType,
		GroupId:     req.GroupId,
		OperatorId:  currentUserId,
		CreatorId:   currentUserId,
		OwnerId:     currentUserId,
	})
	return &v1.CreateProjectRes{
		Id: ProjectId,
	}, err
}
