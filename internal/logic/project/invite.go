package project

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"wktline-server/internal/consts"
	"wktline-server/internal/dao"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"
	"wktline-server/utility"
)

func (s *sProject) GenProjectInviteCode(ctx context.Context, in model_project.GenProjectInviteCodeInput) (code string, err error) {
	inviteId := utility.GenUniIdByGuid()
	_, err = dao.ProjectInvite.Ctx(ctx).OmitEmptyData().Insert(g.Map{
		"id":               inviteId,
		"project_id":       in.ProjectId,
		"max_invite_count": in.MaxInviteCount,
		"deadline":         in.Deadline,
		"inviter_id":       service.BizCtx().Get(ctx).User.Uid,
	})
	return inviteId, err
}

func (s *sProject) GetProjectInviteInfo(ctx context.Context, in model_project.GetProjectInviteInfoInput) (res *model_project.GetProjectInviteInfoOutput, err error) {
	var out model_project.GetProjectInviteInfoOutput
	err = g.Model("project_invite pi").
		LeftJoin("project p", "p.id=pi.project_id").
		LeftJoin("user u", "u.id=pi.inviter_id").
		Fields("pi.id,pi.project_id,pi.inviter_id,pi.max_invite_count,pi.joined_count,pi.deadline").
		Fields("u.username as inviterName,u.avatar as inviterAvatar").
		Fields("p.name as projectName").
		Where("pi.id", in.Id).Scan(&out)
	return &out, err
}

func (s *sProject) InJoinProjectInvite(ctx context.Context, in model_project.InJoinProjectInviteInput) (err error) {
	return dao.Task.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 添加项目用户 - 默认为普通用户
		if err = service.Project().AddProjectUser(ctx, model_project.AddProjectUserInput{
			Id:        utility.GenUniIdByGuid(),
			ProjectId: in.ProjectId,
			UserId:    service.BizCtx().Get(ctx).User.Uid,
			Role:      consts.User,
		}); err != nil {
			return err
		}
		// 更新加入人数
		_, err = dao.ProjectInvite.Ctx(ctx).WherePri(in.Id).Increment("joined_count", 1)
		return err
	})
}

func (s *sProject) GetProjectMyInviteList(ctx context.Context, projectId string) (res *model_project.GetProjectMyInviteListOutput, err error) {
	out := model_project.GetProjectMyInviteListOutput{}
	err = dao.ProjectInvite.Ctx(ctx).
		Where("project_id", projectId).
		Where("inviter_id", service.BizCtx().Get(ctx).User.Uid).
		Fields("id", "max_invite_count", "joined_count", "deadline").
		Scan(&out)
	return &out, err
}
func (s *sProject) CheckProjectMyInvite(ctx context.Context, code string) (count int, err error) {
	return dao.ProjectInvite.Ctx(ctx).WherePri(code).
		Where("inviter_id", service.BizCtx().Get(ctx).User.Uid).
		Count()
}

func (s *sProject) DelProjectMyInvite(ctx context.Context, code string) (err error) {
	_, err = dao.ProjectInvite.Ctx(ctx).WherePri(code).Delete()
	return err
}
