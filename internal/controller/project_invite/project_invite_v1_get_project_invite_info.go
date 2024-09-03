package project_invite

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"wktline-server/internal/dao"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"
	"wktline-server/utility/response"

	"wktline-server/api/project_invite/v1"
)

func (c *ControllerV1) GetProjectInviteInfo(ctx context.Context, req *v1.GetProjectInviteInfoReq) (res *v1.GetProjectInviteInfoRes, err error) {
	r := g.RequestFromCtx(ctx)
	// 合法性校验
	result, err := dao.ProjectInvite.Ctx(ctx).WherePri(req.Code).One()
	if err != nil {
		return nil, err
	}
	var inviteData model_project.ProjectInviteDataItem
	if err = gconv.Scan(result, &inviteData); err != nil {
		return nil, err
	}
	if g.IsEmpty(inviteData) {
		response.JsonExit(r, 5001, "不合法的邀请码")
	}
	out, err := service.Project().GetProjectInviteInfo(ctx, model_project.GetProjectInviteInfoInput{
		Id: req.Code,
	})
	// 判断是否已加入
	uId, err := dao.ProjectUser.Ctx(ctx).
		Where("project_id", inviteData.ProjectId).
		Where("user_id", service.BizCtx().Get(ctx).User.Uid).Value("id")
	if !g.IsEmpty(uId) {
		out.Joined = true
	}
	return (*v1.GetProjectInviteInfoRes)(out), err
}
