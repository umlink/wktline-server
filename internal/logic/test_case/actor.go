package test_case

import (
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"wktline-server/internal/dao"
	"wktline-server/internal/model"
)

func (s *sTestCase) AddTestCaseUser(ctx context.Context, in model.AddTestCaseActorInput) (err error) {
	var actorIds []interface{}
	if err := dao.TestCase.Ctx(ctx).Fields("actor_ids").WherePri(in.Id).Scan(&actorIds); err != nil {
		return err
	}
	// 去重
	ids := garray.NewArrayFrom(actorIds, true)
	var removalIds []string
	for _, id := range in.ActorIds {
		if !ids.Contains(id) {
			removalIds = append(removalIds, id)
		}
	}
	if len(removalIds) == 0 {
		return nil
	}
	_, err = dao.TestCase.Ctx(ctx).Data(g.Map{"actor_ids": removalIds}).WherePri(in.Id).Update()
	return err
}
func (s *sTestCase) DelTestCaseUser(ctx context.Context, in model.DelTestCaseActorInput) (err error) {
	var actorIds []string
	if err := dao.TestCase.Ctx(ctx).Fields("actor_ids").WherePri(in.Id).Scan(&actorIds); err != nil {
		return err
	}
	ids := garray.NewArrayFrom(gconv.Interfaces(in.ActorIds), true)
	var saveIds []string
	for _, id := range actorIds {
		if !ids.Contains(id) {
			saveIds = append(saveIds, id)
		}
	}
	_, err = dao.TestCase.Ctx(ctx).Data(g.Map{"actor_ids": saveIds}).WherePri(in.Id).Update()
	return err
}
