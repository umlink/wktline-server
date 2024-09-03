package message

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/model/model_task"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"wktline-server/internal/dao"
	"wktline-server/internal/model"
	"wktline-server/internal/service"
	"wktline-server/utility"
)

type sMessage struct{}

func New() *sMessage {
	return &sMessage{}
}

func init() {
	service.RegisterMessage(New())
}

func (s *sMessage) GetUnreadMessageCount(ctx context.Context) (count int, err error) {
	return dao.Message.Ctx(ctx).
		Where("status", 0).
		Where("receiver_id", service.BizCtx().Get(ctx).User.Uid).
		Count()
}
func (s *sMessage) GetMyMessageDetail(ctx context.Context, msgId string) (count int, err error) {
	return dao.Message.Ctx(ctx).WherePri(msgId).Where("receiver_id", service.BizCtx().Get(ctx).User.Uid).Count()
}
func (s *sMessage) CreateMessage(ctx context.Context, in model.AddMessageInput) (err error) {
	in.Id = utility.GenUniIdByGuid()
	_, err = dao.Message.Ctx(ctx).OmitEmpty().Insert(in)
	return err
}
func (s *sMessage) CreateTaskMessage(in model.AddMessageInput) (err error) {
	list := g.Slice{}
	// 给所有任务参与者发送消息
	taskActorIds, err := g.Model("task_actor").Array("user_id", "task_id", in.TaskId)
	for _, v := range taskActorIds {
		list = append(list, g.Map{
			"id":          utility.GenUniIdByGuid(),
			"title":       in.Title,
			"content":     in.Content,
			"task_id":     in.TaskId,
			"project_id":  in.ProjectId,
			"sender_id":   in.SenderId,
			"msg_type":    "TASK",
			"receiver_id": gconv.String(v),
		})
	}
	if _, err := g.Model("message").Data(list).Batch(len(list)).Insert(); err != nil {
		return err
	}
	return err
}
func (s *sMessage) CreateProjectActionUsersMsg(ctx context.Context, in model.CreateProjectActionUsersMsgInput) (err error) {
	list := g.Slice{}
	project, err := service.Project().GetProjectDetail(ctx, model_project.GetProjectDetailInput{Id: in.ProjectId})
	for _, v := range in.UserIds {
		list = append(list, g.Map{
			"id":          utility.GenUniIdByGuid(),
			"title":       fmt.Sprintf(in.Title, project.Name),
			"content":     project.Name,
			"project_id":  in.ProjectId,
			"sender_id":   service.BizCtx().Get(ctx).User.Uid,
			"msg_type":    in.MsgType,
			"receiver_id": gconv.String(v),
		})
	}
	if _, err := g.Model("message").Data(list).Batch(len(list)).Insert(); err != nil {
		return err
	}
	return err
}
func (s *sMessage) CreateTaskActionUserMsg(ctx context.Context, in model.CreateTaskActionUsersMsgInput) (err error) {
	list := g.Slice{}
	task, err := service.Task().GetTaskDetail(model_task.GetTaskDetailInput{Id: in.TaskId})
	if err != nil {
		return err
	}
	for _, v := range in.UserIds {
		list = append(list, g.Map{
			"id":          utility.GenUniIdByGuid(),
			"title":       fmt.Sprintf(in.Title, task.Name),
			"content":     task.Name,
			"project_id":  in.ProjectId,
			"sender_id":   service.BizCtx().Get(ctx).User.Uid,
			"msg_type":    in.MsgType,
			"receiver_id": gconv.String(v),
		})
	}
	if _, err := g.Model("message").Data(list).Batch(len(list)).Insert(); err != nil {
		return err
	}
	return err
}
func (s *sMessage) ReadMessage(ctx context.Context, in model.ReadMessageInput) (err error) {
	_, err = dao.Message.Ctx(ctx).WherePri(in.Id).Update("status=1")
	return err
}
func (s *sMessage) GetMessageList(ctx context.Context, in model.GetMessageListInput) (res *model.GetMessageListOutput, err error) {
	var out = &model.GetMessageListOutput{
		PageSize: in.PageSize,
		PageNo:   in.PageNo,
	}
	d := g.Model("message m")
	if !g.IsEmpty(in.Status) {
		d = d.Where("m.status", in.Status)
	}
	if !g.IsEmpty(in.Type) {
		d = d.Where("m.type", in.Type)
	}
	if !g.IsEmpty(in.ProjectId) {
		d = d.Where("m.project_id", in.ProjectId)
	}

	d.LeftJoin("user ru", "ru.id=m.receiver_id").
		LeftJoin("user su", "su.id=m.sender_id").
		LeftJoin("task t", "t.id=m.task_id").
		Fields("m.id,m.title,m.msg_type,m.status,m.created_at,m.sender_id,m.content").
		Fields("su.username as senderName,su.avatar as senderAvatar").
		Fields("t.id as taskId,t.name as taskName").
		Where("m.receiver_id", service.BizCtx().Get(ctx).User.Uid).
		Group("m.id").OrderDesc("m.id")

	err = d.Page(in.PageNo, in.PageSize).ScanAndCount(&out.List, &out.Total, false)
	return out, err
}
func (s *sMessage) GetMessageDetail(ctx context.Context, msgId string) (res *model.GetMessageDetailOutput, err error) {
	var out = &model.GetMessageDetailOutput{}
	d := g.Model("message m").Where("m.id", msgId).
		LeftJoin("user ru", "ru.id=m.receiver_id").
		LeftJoin("user su", "su.id=m.sender_id").
		LeftJoin("task t", "t.id=m.task_id").
		Fields("m.*").
		Fields("t.id as taskId,t.name as taskName").
		Fields("ru.username as receiverName,ru.avatar as receiverAvatar").
		Fields("su.username as senderName,su.avatar as senderAvatar")
	if err = d.Scan(&out); err != nil {
		return nil, err
	}
	if out.ReceiverId != service.BizCtx().Get(ctx).User.Uid {
		return nil, gerror.New("您没有这条消息")
	}
	return out, err
}
