package test_case

import (
	"context"

	"wktline-server/api/test_case/v1"
	"wktline-server/internal/model"
	"wktline-server/internal/service"
)

func (c *ControllerV1) AddTestCaseActors(ctx context.Context, req *v1.AddTestCaseActorsReq) (res *v1.AddTestCaseActorsRes, err error) {
	return nil, service.TestCase().AddTestCaseUser(ctx, model.AddTestCaseActorInput{
		Id:       req.Id,
		ActorIds: req.ActorIds,
	})
}
