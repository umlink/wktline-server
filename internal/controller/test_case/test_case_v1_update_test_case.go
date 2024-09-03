package test_case

import (
	"context"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"wktline-server/internal/model"
	"wktline-server/internal/service"

	"wktline-server/api/test_case/v1"
)

func (c *ControllerV1) UpdateTestCase(ctx context.Context, req *v1.UpdateTestCaseReq) (res *v1.UpdateTestCaseRes, err error) {
	userRole, err := service.Task().CheckUpdateForProjectUserAuth(ctx, req.ProjectId)
	if err != nil {
		return nil, err
	}
	detail, err := service.TestCase().GetTestCaseBaseInfo(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if g.IsEmpty(detail) {
		return nil, gerror.New("测试用例不存在")
	}
	myId := service.BizCtx().Get(ctx).User.Uid
	// lock 机制
	if !g.IsEmpty(detail.EditorId) && detail.EditorId != myId {
		return nil, gerror.New("当前用例正在编辑中")
	}
	// 只有项目普通成员时，检测是否是参与者，管理员可以直接操作
	if userRole == "P-USER" {
		actorIds := garray.NewStrArrayFrom(gconv.Strings(detail.ActorIds), true)
		if myId != detail.CreatorId && !actorIds.Contains(myId) {
			return nil, gerror.New("无更新")
		}
	}
	return nil, service.TestCase().UpdateTestCase(ctx, model.UpdateTestCaseInput{
		Id:          req.Id,
		Name:        req.Name,
		Value:       req.Value,
		Status:      req.Status,
		Progress:    req.Progress,
		Description: req.Description,
	})
}
