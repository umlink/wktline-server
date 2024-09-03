package project

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"wktline-server/api/project/v1"
	"wktline-server/internal/consts"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"
)

func (c *ControllerV1) GetProjectInfo(ctx context.Context, req *v1.GetProjectInfoReq) (res *v1.GetProjectInfoRes, err error) {
	out, err := service.Project().GetProjectDetail(ctx, model_project.GetProjectDetailInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, gerror.New("获取项目详情失败")
	}
	var ret = &v1.GetProjectInfoRes{
		Id:           out.Id,
		Name:         out.Name,
		Description:  out.Description,
		HeaderImg:    out.HeaderImg,
		ShowType:     out.ShowType,
		Status:       out.Status,
		GroupId:      out.GroupId,
		GroupName:    out.GroupName,
		OwnerId:      out.OwnerId,
		OwnerName:    out.OwnerName,
		OwnerAvatar:  out.OwnerAvatar,
		OperatorId:   out.OperatorId,
		OperatorName: out.OperatorName,
		CreatedAt:    out.CreatedAt,
		UpdatedAt:    out.UpdatedAt,
		CanEdit:      false,
	}
	user, err := service.Project().GetProjectUserDetail(ctx, model_project.GetProjectUserDetailInput{
		ProjectId: req.Id,
		UserId:    service.BizCtx().Get(ctx).User.Uid,
	})
	if !g.IsEmpty(user) {
		ret.IsJoined = true
		// 项目负责人或项目管理员
		if user.Role != consts.User {
			ret.CanEdit = true
		}
		// 系统 超级管理员｜管理员
		if service.BizCtx().Get(ctx).User.Role != consts.User {
			ret.CanEdit = true
		}
	}
	return ret, err
}
