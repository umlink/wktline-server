package task

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"wktline-server/internal/consts"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"

	"wktline-server/api/task/v1"
)

func (c *ControllerV1) GetTaskList(ctx context.Context, req *v1.GetTaskListReq) (res *v1.GetTaskListRes, err error) {
	if !g.IsEmpty(req.ProjectId) {
		project, err := service.Project().GetProjectDetail(ctx, model_project.GetProjectDetailInput{
			Id: *req.ProjectId,
		})
		if err != nil {
			return nil, gerror.New("项目不存在")
		}
		// 私有项目则必须用户在项目中才可获取任务
		if project.ShowType == consts.ProjectPrivate {
			user := service.BizCtx().Get(ctx).User
			pUser, err := service.Project().GetProjectUserDetail(ctx, model_project.GetProjectUserDetailInput{
				ProjectId: *req.ProjectId,
				UserId:    user.Uid,
			})
			if err != nil {
				return nil, gerror.New("查询用户归属项目失败")
			}
			// 普通用户必须在项目中才可新建任务
			if g.IsEmpty(pUser) {
				return nil, gerror.New("私有项目中非本项目成员无权获取任务列表")
			}
		}
	}
	out, err := service.Task().GetTaskList(model_task.GetTaskListInput{
		ProjectId: req.ProjectId,
		Keywords:  req.Keywords,
		SortMode:  req.SortMode,
		StatusId:  req.StatusId,
		GroupId:   req.GroupId,
		Priority:  req.Priority,
		TypeId:    req.TypeId,
		ParentId:  req.ParentId,
		HandlerId: req.HandlerId,
		PageSize:  req.PageSize,
		PageNo:    req.PageNo,
	})
	return (*v1.GetTaskListRes)(out), err
}
