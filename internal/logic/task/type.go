package task

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"time"

	"github.com/gogf/gf/v2/frame/g"

	"wktline-server/internal/dao"
	"wktline-server/internal/model/model_task"
)

func (s *sTask) InitProjectTaskType(ctx context.Context, in []model_task.InitTaskTypeInput) (err error) {
	_, err = dao.TaskType.Ctx(ctx).Data(in).Batch(len(in)).Insert()
	return err
}

func (s *sTask) CreateTaskType(ctx context.Context, in model_task.CreateTaskTypeInput) (err error) {
	if _, err = dao.TaskType.Ctx(ctx).Data(in).OmitEmptyData().Insert(); err != nil {
		return nil
	}
	_, err = gcache.Remove(ctx, in.ProjectId+"_types")
	return err
}

func (s *sTask) DelTaskType(ctx context.Context, in model_task.DelTaskTypeInput) (err error) {
	taskCount, err := dao.Task.Ctx(ctx).Where("type_id", in.Id).Count()
	if taskCount > 0 {
		return gerror.New("该类型已被使用，无法删除")
	}
	if _, err = dao.TaskType.Ctx(ctx).WherePri(in.Id).Delete(); err != nil {
		return nil
	}
	_, err = gcache.Remove(ctx, in.ProjectId+"_types")
	return err
}

func (s *sTask) UpdateTaskType(ctx context.Context, in model_task.UpdateTaskTypeInput) (err error) {
	if _, err = dao.TaskType.Ctx(ctx).Data(in).WherePri(in.Id).Update(); err != nil {
		return nil
	}
	_, err = gcache.Remove(ctx, in.ProjectId+"_types")
	return err
}

func (s *sTask) GetTaskType(ctx context.Context, in model_task.GetTaskTypeInput) (res *model_task.GetTaskTypeOutput, err error) {
	cacheKey := in.ProjectId + "_types"
	cache, _ := gcache.Get(ctx, cacheKey)
	if !g.IsEmpty(cache) {
		var cRet model_task.GetTaskTypeOutput
		if err = gconv.Struct(cache, &cRet); err == nil {
			return &cRet, nil
		}
	}
	out := model_task.GetTaskTypeOutput{
		PageNo:   in.PageNo,
		PageSize: in.PageSize,
	}
	d := dao.TaskType.Ctx(ctx).Where("project_id", in.ProjectId)
	if !g.IsEmpty(in.Keywords) {
		likePattern := `%` + in.Keywords + `%`
		d = d.WhereLike("group_name", likePattern)
	}
	if err = d.Page(in.PageNo, in.PageSize).
		Fields("id,name,color").
		ScanAndCount(&out.List, &out.Total, false); err != nil {
		return nil, err
	}
	// 缓存一小时
	if err = gcache.Set(ctx, cacheKey, out, time.Hour); err != nil {
		return nil, gerror.New("项目详情设置缓存失效")
	}
	return &out, err
}
