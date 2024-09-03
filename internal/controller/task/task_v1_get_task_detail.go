package task

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"wktline-server/internal/consts"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"

	"wktline-server/api/task/v1"
)

func (c *ControllerV1) GetTaskDetail(ctx context.Context, req *v1.GetTaskDetailReq) (res *v1.GetTaskDetailRes, err error) {
	task, err := service.Task().GetTaskDetail(model_task.GetTaskDetailInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, gerror.New("查询任务失败")
	}
	if g.IsEmpty(task) {
		return nil, gerror.New("任务不存在")
	}
	project, err := service.Project().GetProjectDetail(ctx, model_project.GetProjectDetailInput{
		Id: task.ProjectId,
	})
	if err != nil {
		return nil, gerror.New("项目不存在")
	}
	// 私有项目则必须用户在项目中才可获取任务
	if project.ShowType == consts.ProjectPrivate {
		user := service.BizCtx().Get(ctx).User
		pUser, err := service.Project().GetProjectUserDetail(ctx, model_project.GetProjectUserDetailInput{
			ProjectId: task.ProjectId,
			UserId:    user.Uid,
		})
		if err != nil {
			return nil, gerror.New("查询用户归属项目失败")
		}
		// 普通用户必须在项目中才可新建任务
		if g.IsEmpty(pUser) {
			return nil, gerror.New("私有项目中非本项目成员无权查看")
		}
	}
	return (*v1.GetTaskDetailRes)(task), err
}
