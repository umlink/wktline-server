package project

import (
	"context"
	"github.com/gogf/gf/v2/os/gcache"
	"time"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"

	"wktline-server/api/project/v1"
)

func (c *ControllerV1) UpdateProject(ctx context.Context, req *v1.UpdateProjectReq) (res *v1.UpdateProjectRes, err error) {
	if err = service.Project().UpdateProject(ctx, model_project.UpdateProjectInput{
		Id:          req.ProjectId,
		Name:        req.Name,
		HeaderImg:   req.HeaderImg,
		Description: req.Description,
		GroupId:     req.GroupId,
		OwnerId:     req.OwnerId,
		ShowType:    req.ShowType,
		Status:      req.Status,
		OperatorId:  service.BizCtx().Get(ctx).User.Uid,
	}); err != nil {
		return nil, err
	}
	// 异步更新缓存
	go func() {
		out, _ := service.Project().GetProjectDetail(ctx, model_project.GetProjectDetailInput{
			Id: req.ProjectId,
		})
		_ = gcache.Set(ctx, out.Id, out, time.Minute)
	}()
	return
}
