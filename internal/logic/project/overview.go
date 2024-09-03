package project

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"sync"
	"wktline-server/internal/consts"
	"wktline-server/internal/dao"
	"wktline-server/internal/model/model_project"
)

func (s *sProject) GetProjectUserTaskStatusStatistics(ctx context.Context, in model_project.GetProjectUserTaskStatusStatisticsInput) (res *model_project.GetProjectUserTaskStatusStatisticsOutput, err error) {
	var out model_project.GetProjectUserTaskStatusStatisticsOutput
	d := g.Model("project_user pu").
		InnerJoin("task t", "pu.user_id=t.handler_id").
		LeftJoin("task_status ts", "t.status_id=ts.id").
		LeftJoin("user u", "u.id=pu.user_id").
		Where("pu.project_id", in.ProjectId).
		Where("t.project_id", in.ProjectId).
		Group("pu.user_id, ts.id").
		Order("pu.user_id, ts.id").
		Fields("u.id as userId,u.userName,u.avatar as userAvatar,ts.id as statusId,ts.name as statusName,count(t.id) as taskCount")

	if !g.IsEmpty(in.GroupId) {
		d.Where("t.group_id", in.GroupId)
	}
	if !g.IsEmpty(in.StartTime) {
		d.Where("t.created_at>=", in.StartTime)
	}
	if !g.IsEmpty(in.EndTime) {
		d.Where("t.created_at<=", in.EndTime)
	}
	if err := d.Scan(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *sProject) GetProjectUserTaskTypeStatistics(ctx context.Context, in model_project.GetProjectUserTaskTypeStatisticsInput) (res *model_project.GetProjectUserTaskTypeStatisticsOutput, err error) {
	var out model_project.GetProjectUserTaskTypeStatisticsOutput
	d := g.Model("project_user pu").
		InnerJoin("task t", "pu.user_id=t.handler_id AND pu.project_id=t.project_id").
		LeftJoin("task_type tt", "t.type_id=tt.id").
		LeftJoin("user u", "u.id=pu.user_id").
		Fields("u.id as userId,u.userName,u.avatar as userAvatar,tt.id as typeId,tt.name as typeName,count(t.id) as taskCount").
		Where("pu.project_id", in.ProjectId).
		Group("pu.user_id, tt.id").
		Order("pu.user_id, tt.id")

	if !g.IsEmpty(in.GroupId) {
		d.Where("t.group_id", in.GroupId)
	}
	if !g.IsEmpty(in.StartTime) {
		d.Where("t.created_at>=", in.StartTime)
	}
	if !g.IsEmpty(in.EndTime) {
		d.Where("t.created_at<=", in.EndTime)
	}
	if err := d.Scan(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *sProject) GetProjectTaskTypeStatistics(ctx context.Context, in model_project.GetProjectTaskTypeStatInput) (res *model_project.GetProjectTaskTypeStatOutput, err error) {
	var out model_project.GetProjectTaskTypeStatOutput
	d := g.Model("task_type tt").
		LeftJoin("task t", "t.type_id=tt.id").
		Fields("tt.id as typeId,tt.name as typeName, count(t.id) as taskCount").
		Where("tt.project_id", in.ProjectId).
		Group("tt.id").
		Order("tt.id")

	if !g.IsEmpty(in.GroupId) {
		d.Where("t.group_id", in.GroupId)
	}
	if !g.IsEmpty(in.StartTime) {
		d.Where("t.created_at>=", in.StartTime)
	}
	if !g.IsEmpty(in.EndTime) {
		d.Where("t.created_at<=", in.EndTime)
	}
	if err := d.Scan(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *sProject) GetProjectTaskStatusStatistics(ctx context.Context, in model_project.GetProjectTaskStatusStatInput) (res *model_project.GetProjectTaskStatusStatOutput, err error) {
	var out model_project.GetProjectTaskStatusStatOutput
	d := g.Model("task_status ts").
		LeftJoin("task t", "t.status_id=ts.id").
		Fields("ts.id as statusId,ts.name as statusName, count(t.id) as taskCount").
		Where("ts.project_id", in.ProjectId).
		Group("ts.id").
		Order("ts.id")

	if !g.IsEmpty(in.GroupId) {
		d.Where("t.group_id", in.GroupId)
	}
	if !g.IsEmpty(in.StartTime) {
		d.Where("t.created_at>=", in.StartTime)
	}
	if !g.IsEmpty(in.EndTime) {
		d.Where("t.created_at<=", in.EndTime)
	}
	if err := d.Scan(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *sProject) GetProjectStatistics(ctx context.Context, in model_project.GetProjectStatisticsInput) (res *model_project.GetProjectStatisticsOut, err error) {
	var out model_project.GetProjectStatisticsOut
	m := dao.Task.Ctx(ctx).Where("project_id", in.ProjectId)
	if !g.IsEmpty(in.GroupId) {
		m = m.Where("group_id", in.GroupId)
		out.GroupCount = 1
	} else {
		out.GroupCount, err = dao.TaskGroup.Ctx(ctx).Count("project_id", in.ProjectId)
	}
	if err != nil {
		return nil, err
	}
	if !g.IsEmpty(in.StartTime) {
		m = m.Where("created_at>=?", in.StartTime)
	}
	if !g.IsEmpty(in.EndTime) {
		m = m.Where("created_at<=?", in.EndTime)
	}
	// 分开查询，后续扩展
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		out.UserCount, err = dao.ProjectUser.Ctx(ctx).Count("project_id", in.ProjectId)
	}()
	go func() {
		defer wg.Done()
		record, err := m.Fields("Count(1) as taskCount,Sum(labor_hour) as laborHour,Sum(plan_hour) as planHour").One()
		fmt.Println(err)
		gMap := record.GMap()
		out.TaskCount = gconv.Int(gMap.Get("taskCount"))
		out.LaborHour = float64(gconv.Int(gMap.Get("laborHour")))
		out.PlanHour = float64(gconv.Int(gMap.Get("planHour")))
	}()
	// 延期完成数
	go func() {
		defer wg.Done()
		// 延期完成：任务状态已结束且，完成时间大于结束时间
		out.OverTimeDoneCount, err = g.Model("task t").
			LeftJoin("task_status ts", "ts.id=t.status_id").
			Where("ts.enum", consts.TaskDone).
			WhereNotNull("t.finished_time").
			Where("t.project_id", in.ProjectId).
			Where("t.finished_time>t.end_time").Count()
		if err != nil {
			fmt.Println(err)
		}
	}()
	// 已延期，未完成
	go func() {
		defer wg.Done()
		out.OverTimeNoDoneCount, err = g.Model("task t").
			LeftJoin("task_status ts", "ts.id=t.status_id").
			WhereNot("ts.enum", consts.TaskDone).
			Where("t.project_id", in.ProjectId).
			Where("t.end_time<", gtime.Date()).Count()
		if err != nil {
			fmt.Println(err)
		}
	}()
	wg.Wait()
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (s *sProject) GetProjectUserLaborHourStat(ctx context.Context, in model_project.GetProjectUserLaborHourStatInput) (res *model_project.GetProjectUserLaborHourStatOutput, err error) {
	var out model_project.GetProjectUserLaborHourStatOutput
	d := g.Model("project_user pu").
		LeftJoin("task_labor_hour tlh", "tlh.project_id=pu.project_id AND pu.user_id=tlh.user_id").
		LeftJoin("user u", "u.id=pu.user_id").
		Fields("u.id as userId,u.username,u.avatar as userAvatar, sum(tlh.hour) as laborHour").
		Where("tlh.project_id", in.ProjectId).
		Group("tlh.user_id").
		Order("tlh.user_id")

	if !g.IsEmpty(in.GroupId) {
		d.Where("tlh.task_id in ?", dao.Task.Ctx(ctx).Fields("id").Where("group_id", in.GroupId))
	}
	if !g.IsEmpty(in.StartTime) {
		d.Where("tlh.date>=", in.StartTime)
	}
	if !g.IsEmpty(in.EndTime) {
		d.Where("tlh.date<=", in.EndTime)
	}
	if err := d.Scan(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *sProject) GetProjectUserTaskCountStat(ctx context.Context, in model_project.GetProjectUserTaskCountStatInput) (res *model_project.GetProjectUserTaskCountStatOutput, err error) {
	var out model_project.GetProjectUserTaskCountStatOutput
	d := g.Model("project_user pu").
		LeftJoin("user u", "u.id=pu.user_id").
		LeftJoin("task t", "t.project_id=pu.project_id AND t.handler_id=pu.user_id").
		Fields("u.id as userId,u.username,u.avatar as userAvatar, count(t.id) as taskCount").
		Where("pu.project_id", in.ProjectId).
		Group("pu.user_id").
		Order("pu.user_id")

	if !g.IsEmpty(in.GroupId) {
		d.Where("t.group_id", in.GroupId)
	}
	if !g.IsEmpty(in.StartTime) {
		d.Where("t.created_at>=", in.StartTime)
	}
	if !g.IsEmpty(in.EndTime) {
		d.Where("t.created_at<=", in.EndTime)
	}
	if err := d.Scan(&out); err != nil {
		return nil, err
	}
	return &out, nil
}
