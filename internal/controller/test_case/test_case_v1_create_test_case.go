package test_case

import (
	"context"
	"wktline-server/internal/model"
	"wktline-server/internal/service"

	"wktline-server/api/test_case/v1"
)

func (c *ControllerV1) CreateTestCase(ctx context.Context, req *v1.CreateTestCaseReq) (res *v1.CreateTestCaseRes, err error) {
	// 非项目成员无法添加测试用例
	if _, err := service.Task().CheckUpdateForProjectUserAuth(ctx, req.ProjectId); err != nil {
		return nil, err
	}
	return nil, service.TestCase().CreateTestCase(ctx, model.CreateTestCaseInput{
		Name:        req.Name,
		ProjectId:   req.ProjectId,
		Description: req.Description,
	})
}
