package test_case

import (
	"context"

	"wktline-server/api/test_case/v1"
	"wktline-server/internal/model"
	"wktline-server/internal/service"
)

func (c *ControllerV1) DelTestCaseActors(ctx context.Context, req *v1.DelTestCaseActorsReq) (res *v1.DelTestCaseActorsRes, err error) {
	return nil, service.TestCase().DelTestCaseUser(ctx, model.DelTestCaseActorInput{
		Id:       req.Id,
		ActorIds: req.ActorIds,
	})
}
