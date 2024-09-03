package test_case

import (
	"context"
	"wktline-server/internal/model"
	"wktline-server/internal/service"

	"wktline-server/api/test_case/v1"
)

func (c *ControllerV1) GetTestCaseList(ctx context.Context, req *v1.GetTestCaseListReq) (res *v1.GetTestCaseListRes, err error) {
	out, err := service.TestCase().GetTestCaseList(ctx, model.GetTestCaseListInput{
		ProjectId: req.ProjectId,
		CreatorId: req.CreatorId,
		PageSize:  req.PageSize,
		PageNo:    req.PageSize,
	})
	return (*v1.GetTestCaseListRes)(out), err
}
