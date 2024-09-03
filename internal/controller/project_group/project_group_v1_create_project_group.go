package project_group

import (
	"context"

	"wktline-server/utility"

	"wktline-server/api/project_group/v1"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"
)

func (c *ControllerV1) CreateProjectGroup(ctx context.Context, req *v1.CreateProjectGroupReq) (res *v1.CreateProjectGroupRes, err error) {
	groupId := utility.GenUniIdByGuid()
	if err != nil {
		return nil, err
	}
	err = service.Project().CreateProjectGroup(ctx, model_project.CreateGroupInput{
		Id:          groupId,
		Name:        req.Name,
		Description: req.Description,
		OperatorId:  service.BizCtx().Get(ctx).User.Uid,
	})
	return &v1.CreateProjectGroupRes{
		Id: groupId,
	}, err
}
