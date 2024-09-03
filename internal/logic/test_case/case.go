package test_case

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"wktline-server/internal/dao"
	"wktline-server/internal/model"
	"wktline-server/internal/model/entity"
	"wktline-server/internal/service"
	"wktline-server/utility"
)

type sTestCase struct{}

func init() {
	service.RegisterTestCase(New())
}

func New() *sTestCase {
	return &sTestCase{}
}
func (s *sTestCase) AddTestCaseLock(ctx context.Context, id string) (err error) {
	_, err = dao.TestCase.Ctx(ctx).Data(g.Map{
		"editor_id": service.BizCtx().Get(ctx).User.Uid,
	}).WherePri(id).Update()
	return err
}
func (s *sTestCase) DelTestCaseLock(ctx context.Context, id string) (err error) {
	_, err = dao.TestCase.Ctx(ctx).Data(g.Map{
		"editor_id": nil,
	}).WherePri(id).Update()
	return err
}
func (s *sTestCase) GetTestCaseBaseInfo(ctx context.Context, id string) (out entity.TestCase, err error) {
	err = dao.TestCase.Ctx(ctx).WherePri(id).Scan(&out)
	return out, err
}
func (s *sTestCase) CreateTestCase(ctx context.Context, in model.CreateTestCaseInput) (err error) {
	in.Id = utility.GenUniIdByGuid()
	_, err = dao.TestCase.Ctx(ctx).Data(in).Insert()
	return err
}
func (s *sTestCase) UpdateTestCase(ctx context.Context, in model.UpdateTestCaseInput) (err error) {
	_, err = dao.TestCase.Ctx(ctx).Data(in).OmitEmptyData().WherePri(in.Id).Update()
	return err
}
func (s *sTestCase) DelTestCase(ctx context.Context, id string) (err error) {
	_, err = dao.TestCase.Ctx(ctx).WherePri(id).Delete()
	return err
}
func (s *sTestCase) GetTestCaseList(ctx context.Context, in model.GetTestCaseListInput) (res *model.GetTestCaseListOutput, err error) {
	var out = &model.GetTestCaseListOutput{
		PageNo:   in.PageNo,
		PageSize: in.PageSize,
	}
	d := g.Model("test_case tc")
	if !g.IsEmpty(in.CreatorId) {
		d = d.Where("creator_id", in.CreatorId)
	}
	if err = d.LeftJoin("user u1", "u1.id=tc.creator_id").
		LeftJoin("user u2", "u2.id=tc.editor_id").
		Fields("tc.id,tc.name,tc.project_id,tc.actor_ids,tc.creator_id,tc.editor_id,tc.progress,tc.status,tc.created_at,tc.updated_at").
		Fields("u1.username as creatorName,u1.avatar as creatorAvatar").
		Fields("u2.username as editorName,u2.avatar as editorAvatar").
		OrderDesc("tc.id").
		ScanAndCount(&out.List, &out.Total, false); err != nil {
		return nil, err
	}
	return out, err
}
func (s *sTestCase) GetTestCaseDetail(id string) (res *model.GetTestCaseDetailOutput, err error) {
	var out model.GetTestCaseDetailOutput
	err = g.Model("test_case tc").
		LeftJoin("user u1", "u1.id=tc.creator_id").
		LeftJoin("user u2", "u2.id=tc.editor_id").
		Fields("tc.id,tc.name,tc.project_id,tc.actor_ids,tc.value,tc.creator_id,tc.editor_id,tc.progress,tc.status,tc.created_at,tc.updated_at").
		Fields("u1.username as creatorName,u1.avatar as creatorAvatar").
		Fields("u2.username as editorName,u2.avatar as editorAvatar").
		OrderDesc("tc.id").
		Where("tc.id", id).Scan(&out)
	return &out, nil
}
