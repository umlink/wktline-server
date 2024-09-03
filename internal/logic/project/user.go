package project

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
	"wktline-server/internal/model"
	"wktline-server/internal/service"
	"wktline-server/utility"

	"github.com/gogf/gf/v2/frame/g"

	"wktline-server/internal/dao"
	"wktline-server/internal/model/model_project"
)

func (s *sProject) GetProjectUserList(ctx context.Context, in model_project.GetProjectUserListInput) (res *model_project.GetProjectUserListOutput, err error) {
	out := &model_project.GetProjectUserListOutput{
		PageSize: in.PageSize,
		PageNo:   in.PageNo,
	}
	d := g.Model("project_user pu").Where("pu.project_id", in.ProjectId)
	d = d.InnerJoin("user u", "u.id=pu.user_id")
	if !g.IsEmpty(in.Keywords) {
		likePattern := `%` + in.Keywords + `%`
		d = d.WhereLike("u.username", likePattern).WhereOrLike("u.nickname", likePattern)
	}
	if !g.IsEmpty(in.Role) {
		d = d.Where("pu.role", in.Role)
	}
	d = d.Fields("pu.id,u.id as user_id,u.username,u.nickname,u.avatar,u.email,pu.role")
	err = d.Page(in.PageNo, in.PageSize).ScanAndCount(&out.List, &out.Total, false)
	return out, err
}

func (s *sProject) GetProjectUserDetail(ctx context.Context, in model_project.GetProjectUserDetailInput) (res *model_project.GetProjectUserDetailOutput, err error) {
	cacheKey := in.ProjectId + in.UserId
	cache, _ := gcache.Get(ctx, cacheKey)
	if !g.IsEmpty(cache) {
		var cRet model_project.GetProjectUserDetailOutput
		if err = gconv.Struct(cache, &cRet); err == nil {
			return &cRet, nil
		}
	}
	err = dao.ProjectUser.Ctx(ctx).Where("project_id", in.ProjectId).Where("user_id", in.UserId).Scan(&res)
	// 这是缓存一小时
	if err = gcache.Set(ctx, cacheKey, res, time.Hour); err != nil {
		return nil, gerror.New("项目详情设置缓存失效")
	}
	return res, err
}

func (s *sProject) GetProjectUserIds(ctx context.Context, projectId string) (ids []string, err error) {
	userIds, err := dao.ProjectUser.Ctx(ctx).Array("user_id", "project_id=", projectId)
	ids = gconv.Strings(userIds)
	return ids, err
}

func (s *sProject) AddProjectUserByIds(ctx context.Context, in model_project.AddProjectUserByIdsInput) (err error) {
	userIds, err := s.GetProjectUserIds(ctx, in.ProjectId)
	if err != nil {
		return gerror.New("查询项目用户异常")
	}
	array := garray.NewStrArrayFrom(userIds, true)
	list := g.Slice{}
	// 去重
	var newUserIds []string
	for _, v := range in.UserIds {
		if !array.Contains(v) {
			list = append(list, g.Map{
				"id":         utility.GenUniIdByGuid(),
				"project_id": in.ProjectId,
				"user_id":    v,
			})
			newUserIds = append(newUserIds, v)
		}
	}
	if len(list) == 0 {
		return nil
	}
	// 批量插入 batch 默认是10，需和实际插入值数量保 持一致
	if _, err := dao.ProjectUser.Ctx(ctx).Data(list).Batch(len(list)).Insert(); err != nil {
		return err
	}
	go func() {
		err = service.Message().CreateProjectActionUsersMsg(ctx, model.CreateProjectActionUsersMsgInput{
			ProjectId: in.ProjectId,
			UserIds:   newUserIds,
			Title:     "你被添加到项目【%s】",
			MsgType:   "PROJECT_ADD_USER",
		})
		fmt.Println(err)
	}()
	return nil
}

func (s *sProject) AddProjectUser(ctx context.Context, in model_project.AddProjectUserInput) (err error) {
	if _, err := dao.ProjectUser.Ctx(ctx).Data(in).Insert(); err != nil {
		return err
	}
	return nil
}

func (s *sProject) UpdateProjectUserRole(ctx context.Context, in model_project.UpdateProjectUserRoleInput) (err error) {
	if _, err := dao.ProjectUser.Ctx(ctx).Data(g.Map{
		"role": in.Role,
	}).Where("user_id", in.UserId).Where("project_id", in.ProjectId).Update(); err != nil {
		return err
	}
	return nil
}

func (s *sProject) DelProjectUser(ctx context.Context, in model_project.DelProjectUserInput) (err error) {
	if _, err = dao.ProjectUser.Ctx(ctx).Where("project_id", in.ProjectId).WhereIn("user_id", in.UserIds).Delete(); err != nil {
		return err
	}
	// 用户移出，则删除缓存
	var cacheKeys []string
	for _, v := range in.UserIds {
		cacheKeys = append(cacheKeys, in.ProjectId+v)
	}
	_, err = gcache.Remove(ctx, cacheKeys)
	return err
}
