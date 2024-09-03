package project_invite

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"wktline-server/internal/dao"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"
	"wktline-server/utility/response"

	"github.com/gogf/gf/v2/errors/gerror"

	"wktline-server/api/project_invite/v1"
)

func (c *ControllerV1) InJoinInviteProject(ctx context.Context, req *v1.InJoinInviteProjectReq) (res *v1.InJoinInviteProjectRes, err error) {
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
		return nil, gerror.New("不合法的邀请码")
	}
	if inviteData.JoinedCount >= inviteData.MaxInviteCount {
		response.JsonExit(r, 5001, fmt.Sprintf("该邀请链接已达到加入人数上限 %s 人", gconv.String(inviteData.MaxInviteCount)))
	}
	// 判断是否已加入
	uId, err := dao.ProjectUser.Ctx(ctx).
		Where("project_id", inviteData.ProjectId).
		Where("user_id", service.BizCtx().Get(ctx).User.Uid).Value("id")

	if !g.IsEmpty(uId) {
		// 提示加入成功，不做任何处理
		response.JsonExit(r, 2001, "已加入该项目")
	}
	err = service.Project().InJoinProjectInvite(ctx, model_project.InJoinProjectInviteInput{
		Id:        req.Code,
		ProjectId: inviteData.ProjectId,
	})
	return nil, err
}
