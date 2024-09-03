package test_case

import (
	"context"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"wktline-server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"

	"wktline-server/api/test_case/v1"
)

func (c *ControllerV1) AddTestCaseLock(ctx context.Context, req *v1.AddTestCaseLockReq) (res *v1.AddTestCaseLockRes, err error) {
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
	actorIds := garray.NewStrArrayFrom(gconv.Strings(detail.ActorIds), true)
	if !actorIds.Contains(myId) {
		return nil, gerror.New("无权修改")
	}
	return nil, service.TestCase().AddTestCaseLock(ctx, req.Id)
}
