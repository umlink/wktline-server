package task

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"time"

	"wktline-server/internal/dao"
	"wktline-server/internal/model/model_task"
)

var cacheKey = "project_priority"

func (s *sTask) AddTaskPriority(ctx context.Context, in model_task.AddTaskPriorityInput) (err error) {
	if _, err = dao.TaskPriority.Ctx(ctx).Data(in).Insert(); err != nil {
		return nil
	}
	_, err = gcache.Remove(ctx, cacheKey)
	return err
}
func (s *sTask) UpdateTaskPriority(ctx context.Context, in model_task.UpdateTaskPriorityInput) (err error) {
	if _, err = dao.TaskPriority.Ctx(ctx).Data(in).WherePri(in.Id).Update(); err != nil {
		return nil
	}
	_, err = gcache.Remove(ctx, cacheKey)
	return err
}
func (s *sTask) DelTaskPriority(ctx context.Context, id string) (err error) {
	if _, err = dao.TaskPriority.Ctx(ctx).WherePri(id).Delete(); err != nil {
		return nil
	}
	_, err = gcache.Remove(ctx, cacheKey)
	return err
}
func (s *sTask) GetTaskPriority(ctx context.Context, in model_task.GetTaskPriorityInput) (res *model_task.GetTaskPriorityOutput, err error) {
	cache, _ := gcache.Get(ctx, cacheKey)
	if !g.IsEmpty(cache) {
		var cRet model_task.GetTaskPriorityOutput
		if err = gconv.Struct(cache, &cRet); err == nil {
			return &cRet, nil
		}
	}
	out := model_task.GetTaskPriorityOutput{
		PageNo:   in.PageNo,
		PageSize: in.PageSize,
	}
	d := dao.TaskPriority.Ctx(ctx)
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
