package task

import (
	"context"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"

	"wktline-server/internal/service"
	"wktline-server/utility"

	"github.com/gogf/gf/v2/frame/g"

	"wktline-server/internal/dao"
	"wktline-server/internal/model/model_task"
)

func (s *sTask) GetTaskActorList(ctx context.Context, in model_task.GetTaskActorListInput) (res *model_task.GetTaskActorListOutput, err error) {
	out := &model_task.GetTaskActorListOutput{
		PageSize: in.PageSize,
		PageNo:   in.PageNo,
	}
	d := g.Model("task_actor tu").Where("tu.task_id", in.TaskId)
	d = d.InnerJoin("user u", "u.id=tu.user_id")
	if !g.IsEmpty(in.Keywords) {
		likePattern := `%` + in.Keywords + `%`
		d = d.WhereLike("u.username", likePattern).WhereOrLike("u.nickname", likePattern)
	}
	d = d.Fields("u.id,u.username,u.nickname,u.avatar")
	err = d.Page(in.PageNo, in.PageSize).ScanAndCount(&out.List, &out.Total, false)
	return out, err
}

func (s *sTask) AddTaskActor(ctx context.Context, in model_task.AddTaskActorInput) (err error) {
	var userIds []string
	if in.IsNew {
		userIds = in.UserIds
	} else {
		ids, err := dao.TaskActor.Ctx(ctx).Fields("user_id").Where("task_id", in.TaskId).Array()
		if err != nil {
			return gerror.New("查询任务参与者异常")
		}
		array := garray.NewStrArrayFrom(gconv.Strings(ids), true)
		for _, v := range in.UserIds {
			if !array.Contains(v) {
				userIds = append(userIds, v)
			}
		}
	}
	list := g.Slice{}
	for _, v := range userIds {
		list = append(list, g.Map{
			"id":          utility.GenUniIdByGuid(),
			"task_id":     in.TaskId,
			"user_id":     v,
			"operator_id": service.BizCtx().Get(ctx).User.Uid,
		})
	}
	if _, err := dao.TaskActor.Ctx(ctx).Data(list).Batch(len(list)).Insert(); err != nil {
		return err
	}
	return nil
}

func (s *sTask) DelTaskActor(ctx context.Context, in model_task.DelTaskActorInput) (err error) {
	_, err = dao.TaskActor.Ctx(ctx).Where("task_id", in.TaskId).WhereIn("user_id", in.UserIds).Delete()
	return err
}

func (s *sTask) FindTaskActor(ctx context.Context, in model_task.FindTaskActorInput) (count int, err error) {
	return dao.TaskActor.Ctx(ctx).Where(in).Count()
}
