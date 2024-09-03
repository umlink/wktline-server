package project_group

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"wktline-server/api/project_group/v1"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"
)

func (c *ControllerV1) DelProjectGroup(ctx context.Context, req *v1.DelProjectGroupReq) (res *v1.DelProjectGroupRes, err error) {
	count, err := service.Project().GetProjectCountFromGroupId(ctx, req.Id)
	if count > 0 {
		return nil, gerror.New("该分组下已存在项目，无法删除")
	}
	err = service.Project().DeleteGroupById(ctx, model_project.DelGroupInput{
		Id: req.Id,
	})
	return nil, err
}
