package project

import (
	"context"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
	"wktline-server/utility"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"wktline-server/internal/consts"
	"wktline-server/internal/dao"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"
)

type sProject struct{}

func init() {
	service.RegisterProject(New())
}

func New() *sProject {
	return &sProject{}
}

func (s *sProject) CreateProject(ctx context.Context, in model_project.CreateProjectInput) (err error) {
	return dao.Project.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 1. 创建项目
		if _, err := dao.Project.Ctx(ctx).Data(in).Insert(); err != nil {
			return err
		}
		// 2. 初始化项目下的任务状态
		if err := service.Task().InitProjectTaskStatus(ctx, consts.GetDefaultTaskStatusList(in.Id, in.OperatorId)); err != nil {
			return err
		}
		// 3. 初始化项目下任务类型列表
		if err := service.Task().InitProjectTaskType(ctx, consts.GetDefaultTaskTypeList(in.Id, in.OperatorId)); err != nil {
			return err
		}
		// 添加项目用户为项目负责人
		err = service.Project().AddProjectUser(ctx, model_project.AddProjectUserInput{
			Id:        utility.GenUniIdByGuid(),
			ProjectId: in.Id,
			UserId:    service.BizCtx().Get(ctx).User.Uid,
			Role:      consts.SuperAdmin,
		})
		return err
	})
}

func (s *sProject) DelProjectById(ctx context.Context, in model_project.DelProjectInput) (err error) {
	_, err = dao.Project.Ctx(ctx).WherePri(in.Id).Delete()
	return err
}

func (s *sProject) UpdateProject(ctx context.Context, in model_project.UpdateProjectInput) (err error) {
	return dao.Project.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		out, err := s.GetProjectDetail(ctx, model_project.GetProjectDetailInput{
			Id: in.Id,
		})
		if err != nil {
			return nil
		}
		// 普通用户只能为项目负责人时，才可以修改项目负责人，系统管理员不受限
		role := service.BizCtx().Get(ctx).User.Role
		if role == consts.User && out.OwnerId != service.BizCtx().Get(ctx).User.Uid {
			return gerror.New("项目负责人只能由当前项目负责人或系统管理员修改")
		}
		// 上一个项目负责人改为普通用户
		if err := service.Project().UpdateProjectUserRole(ctx, model_project.UpdateProjectUserRoleInput{
			ProjectId: out.Id,
			UserId:    out.OwnerId,
			Role:      consts.User,
		}); err != nil {
			return nil
		}
		if _, err = dao.Project.Ctx(ctx).Data(in).OmitEmptyData().WherePri(in.Id).Update(); err != nil {
			return nil
		}
		// 最新用户改为项目负责人
		return service.Project().UpdateProjectUserRole(ctx, model_project.UpdateProjectUserRoleInput{
			ProjectId: in.Id,
			UserId:    *in.OwnerId,
			Role:      consts.SuperAdmin,
		})
	})
}

func (s *sProject) GetProjectDetail(ctx context.Context, in model_project.GetProjectDetailInput) (project *model_project.GetProjectDetailOutput, err error) {
	cache, _ := gcache.Get(ctx, in.Id)
	if !g.IsEmpty(cache) {
		var cRet model_project.GetProjectDetailOutput
		if err = gconv.Struct(cache, &cRet); err == nil {
			return &cRet, nil
		}
	}
	err = g.Model("project p").Where("p.id", in.Id).
		InnerJoin("user u1", "u1.id=p.owner_id").
		InnerJoin("user u2", "u2.id=p.operator_id").
		InnerJoin("project_group pg", "pg.id=p.group_id").
		Fields("p.*, u1.username as owner_name, u1.avatar as owner_avatar, u2.username as operator_name, pg.name as group_name").
		Scan(&project)

	// 这是缓存一小时
	if err = gcache.Set(ctx, project.Id, project, time.Hour); err != nil {
		return nil, gerror.New("项目详情设置缓存失效")
	}
	return project, err
}

func (s *sProject) GetProjectCountFromGroupId(ctx context.Context, in string) (count int, err error) {
	return dao.Project.Ctx(ctx).Count("group_id", in)
}

func (s *sProject) GetProjectList(ctx context.Context, in model_project.GetProjectListInput) (res *model_project.GetProjectListOutput, err error) {
	out := &model_project.GetProjectListOutput{
		PageSize: in.PageSize,
		PageNo:   in.PageNo,
	}
	userId := service.BizCtx().Get(ctx).User.Uid
	userRole := service.BizCtx().Get(ctx).User.Role

	d := g.Model("project p")

	if !g.IsEmpty(in.Keywords) {
		likePattern := `%` + *in.Keywords + `%`
		d = d.WhereLike("p.name", likePattern)
	}
	if !g.IsEmpty(in.GroupId) {
		d = d.Where("p.group_id", in.GroupId)
	}
	if in.IsOwner {
		d = d.Where("p.owner_id", userId)
	}
	if in.IsCreator {
		d = d.Where("p.creator_id", userId)
	}
	if in.IsJoined {
		d = d.WhereIn("p.id", dao.ProjectUser.Ctx(ctx).Fields("project_id").
			Where("user_id", userId))
	}
	// 非拥有，创建，加入，则查询时需检测所有已加入和开放类型的项目
	if !in.IsOwner && !in.IsCreator && !in.IsJoined {
		if !g.IsEmpty(in.ShowType) {
			// 非普通用户，即管理员和超管都可以查看所有类型的项目,系统管理员无需加入即可查询
			if userRole != consts.User {
				d = d.Where("p.show_type", *in.ShowType)
			} else {
				// 非管理员查询
				if *in.ShowType == consts.ProjectPrivate {
					d = d.WhereIn("p.id", dao.ProjectUser.Ctx(ctx).Fields("project_id").
						Where("user_id", userId)).Where("p.show_type", consts.ProjectPrivate)
				} else {
					// 公开包括 我加入的和未加入的
					d = d.Where(d.Builder().
						WhereIn("p.id", dao.ProjectUser.Ctx(ctx).Fields("project_id").
							Where("user_id", userId)).WhereOr("p.show_type", consts.ProjectPublic))
				}
			}
		} else {
			// 普通用户只能查询已加入的和公开的 公开包括 我加入的和未加入的
			if userRole == consts.User {
				d = d.Where(d.Builder().
					WhereIn("p.id", dao.ProjectUser.Ctx(ctx).Fields("project_id").
						Where("user_id", userId)).WhereOr("p.show_type", consts.ProjectPublic))
			}
		}
	} else {
		// 带查询条件则筛选显示类型
		if !g.IsEmpty(in.ShowType) {
			d = d.Where("p.show_type", *in.ShowType)
		}
	}
	err = d.InnerJoin("user u1", "u1.id=p.owner_id").
		InnerJoin("project_group pg", "pg.id=p.group_id").
		Fields("p.*, u1.username as owner_name, pg.name as group_name").
		OrderDesc("p.id").
		Page(in.PageNo, in.PageSize).ScanAndCount(&out.List, &out.Total, false)
	return out, err
}
