package test_case

import (
	"context"
	"wktline-server/internal/service"

	"wktline-server/api/test_case/v1"
)

func (c *ControllerV1) GetTestCase(ctx context.Context, req *v1.GetTestCaseReq) (res *v1.GetTestCaseRes, err error) {
	out, err := service.TestCase().GetTestCaseDetail(req.Id)
	return (*v1.GetTestCaseRes)(out), err
}
