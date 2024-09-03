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

// InitProjectTaskStatus 项目创建时同步调用一次 初始化项目下任务状态列表
func (s *sTask) InitProjectTaskStatus(ctx context.Context, in []model_task.InitTaskStatusInput) (err error) {
	if _, err = dao.TaskStatus.Ctx(ctx).Data(in).Batch(len(in)).Insert(); err != nil {
		return nil
	}
	return
}

func (s *sTask) CreateTaskStatus(ctx context.Context, in model_task.CreateTaskStatusInput) (err error) {
	in.Sort, err = dao.TaskStatus.Ctx(ctx).Where("project_id", in.ProjectId).Count()
	in.Sort += 1
	if err != nil {
		return err
	}
	if _, err = dao.TaskStatus.Ctx(ctx).Data(in).OmitEmptyData().Insert(); err != nil {
		return nil
	}
	_, err = gcache.Remove(ctx, in.ProjectId+"_status")
	return err
}

func (s *sTask) DelTaskStatus(ctx context.Context, in model_task.DelTaskStatusInput) (err error) {
	if _, err = dao.TaskStatus.Ctx(ctx).WherePri(in.Id).Delete(); err != nil {
		return nil
	}
	_, err = gcache.Remove(ctx, in.ProjectId+"_status")
	return err
}

func (s *sTask) GetTaskStatusById(ctx context.Context, statusId string) (res *model_task.GetTaskStatusByIdOutput, err error) {
	err = dao.TaskStatus.Ctx(ctx).WherePri(statusId).Scan(&res)
	return res, err
}

func (s *sTask) UpdateTaskStatus(ctx context.Context, in model_task.UpdateTaskStatusInput) (err error) {
	if _, err = dao.TaskStatus.Ctx(ctx).Data(in).WherePri(in.Id).Update(); err != nil {
		return nil
	}
	_, err = gcache.Remove(ctx, in.ProjectId+"_status")
	return err
}

func (s *sTask) UpdateTaskStatusSort(ctx context.Context, in model_task.UpdateTaskStatusSortInput) (err error) {
	for _, v := range in.SortMapList {
		_, err := dao.TaskStatus.Ctx(ctx).Data(g.Map{
			"sort": v.Sort,
		}).WherePri(v.Id).Update()
		if err != nil {
			return err
		}
	}
	_, err = gcache.Remove(ctx, in.ProjectId+"_status")
	return nil
}

func (s *sTask) GetTaskStatus(ctx context.Context, in model_task.GetTaskStatusInput) (res *model_task.GetTaskStatusOutput, err error) {
	cacheKey := in.ProjectId + "_status"
	cache, _ := gcache.Get(ctx, cacheKey)
	if !g.IsEmpty(cache) {
		var cRet model_task.GetTaskStatusOutput
		if err = gconv.Struct(cache, &cRet); err == nil {
			return &cRet, nil
		}
	}
	out := model_task.GetTaskStatusOutput{
		PageNo:   in.PageNo,
		PageSize: in.PageSize,
	}
	d := dao.TaskStatus.Ctx(ctx).Where("project_id", in.ProjectId)
	if err = d.Page(in.PageNo, in.PageSize).
		Fields("id,name,color,enum,sort,default").OrderAsc("sort").
		ScanAndCount(&out.List, &out.Total, false); err != nil {
		return nil, err
	}
	// 这是缓存一小时
	if err = gcache.Set(ctx, cacheKey, out, time.Hour); err != nil {
		return nil, gerror.New("项目详情设置缓存失效")
	}
	return &out, err
}

func (s *sTask) GetTaskCountByStatus(ctx context.Context, id string) (res int, err error) {
	return dao.Task.Ctx(ctx).Where("status_id", id).Count()
}
