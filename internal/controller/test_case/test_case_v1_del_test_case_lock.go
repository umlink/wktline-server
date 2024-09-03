package test_case

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"wktline-server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"

	"wktline-server/api/test_case/v1"
)

func (c *ControllerV1) DelTestCaseLock(ctx context.Context, req *v1.DelTestCaseLockReq) (res *v1.DelTestCaseLockRes, err error) {
	detail, err := service.TestCase().GetTestCaseBaseInfo(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if g.IsEmpty(detail) {
		return nil, gerror.New("测试用例不存在")
	}
	if g.IsEmpty(detail.EditorId) {
		return nil, nil
	}
	myId := service.BizCtx().Get(ctx).User.Uid
	if detail.EditorId != myId {
		return nil, gerror.New("需等待他人编辑完毕")
	}
	return nil, service.TestCase().DelTestCaseLock(ctx, req.Id)
}
