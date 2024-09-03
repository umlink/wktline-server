package test_case

import (
	"context"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"wktline-server/internal/service"

	"wktline-server/api/test_case/v1"
)

func (c *ControllerV1) DelTestCase(ctx context.Context, req *v1.DelTestCaseReq) (res *v1.DelTestCaseRes, err error) {
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
	// 只有项目普通成员时，检测是否是参与者，管理员可以直接操作
	if userRole == "P-USER" {
		myId := service.BizCtx().Get(ctx).User.Uid
		actorIds := garray.NewStrArrayFrom(gconv.Strings(detail.ActorIds), true)
		if myId != detail.CreatorId && !actorIds.Contains(myId) {
			return nil, gerror.New("无权限")
		}
	}
	return nil, service.TestCase().DelTestCase(ctx, req.Id)
}
