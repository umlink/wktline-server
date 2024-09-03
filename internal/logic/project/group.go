package project

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"wktline-server/internal/dao"
	"wktline-server/internal/model/model_project"
)

func (s *sProject) CreateProjectGroup(ctx context.Context, in model_project.CreateGroupInput) (err error) {
	_, err = dao.ProjectGroup.Ctx(ctx).Insert(in)
	return err
}

func (s *sProject) DeleteGroupById(ctx context.Context, in model_project.DelGroupInput) (err error) {
	_, err = dao.ProjectGroup.Ctx(ctx).WherePri(in.Id).Delete()
	return err
}

func (s *sProject) UpdateGroup(ctx context.Context, in model_project.UpdateGroupInput) (err error) {
	_, err = dao.ProjectGroup.Ctx(ctx).Data(in).WherePri(in.Id).Update()
	return err
}

func (s *sProject) GetGroupDetail(ctx context.Context, in model_project.GetGroupDetailInput) (res *model_project.GetGroupDetailOutput, err error) {
	err = dao.ProjectGroup.Ctx(ctx).WherePri(in.Id).Scan(&res)
	return res, err
}

func (s *sProject) GetGroupList(ctx context.Context, in model_project.GetGroupListInput) (res *model_project.GetGroupListOutput, err error) {
	out := model_project.GetGroupListOutput{
		PageNo:   in.PageNo,
		PageSize: in.PageSize,
	}
	d := dao.ProjectGroup.Ctx(ctx)
	if !g.IsEmpty(in.Keywords) {
		likePattern := `%` + in.Keywords + `%`
		d = d.WhereLike("name", likePattern)
	}
	if err = d.Page(in.PageNo, in.PageSize).ScanAndCount(&out.List, &out.Total, false); err != nil {
		return nil, err
	}
	return &out, err
}
