package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TaskAttachmentItem struct {
	Id   string `json:"id" v:"required" dc:"附件id"`
	Name string `json:"name" v:"required" dc:"附件名"`
	Type string `json:"type" v:"required" dc:"附件类型"`
	Url  string `json:"url" v:"required" dc:"附件url"`
	Size string `json:"size" v:"required" dc:"附件大小（kb）"`
}

type AddTaskAttachmentReq struct {
	g.Meta     `path:"/task/addTaskAttachment" method:"post" summary:"添加任务附件" tags:"TaskAttachment"`
	ProjectId  string `json:"projectId" v:"required#项目id不能为空" dc:"项目id"`
	TaskId     string `json:"taskId" v:"required#任务id不能为空" dc:"任务 id"`
	ResourceId string `json:"resourceId" v:"required#资源id不能为空" dc:"类型"`
}
type AddTaskAttachmentRes struct {
	Id string `json:"id" v:"required" dc:"任务类型 id"`
}

type DelTaskAttachmentReq struct {
	g.Meta    `path:"/task/delTaskAttachment" method:"post" summary:"删除" tags:"TaskAttachment"`
	ProjectId string `json:"projectId" v:"required#项目id不能为空" dc:"项目id"`
	Id        string `json:"id" v:"required#附件id不能为空" dc:"附件id"`
}
type DelTaskAttachmentRes struct{}

type GetTaskAttachmentListReq struct {
	g.Meta    `path:"/task/getTaskAttachmentList" method:"get" summary:"获取附件（项目｜任务｜创建者）" tags:"TaskAttachment"`
	ProjectId *string `json:"projectId"  dc:"项目id"`
	TaskId    string  `json:"taskId" dc:"任务 id"`
	CreatorId *string `json:"creatorId" dc:"创建用户 id"`
	PageNo    int     `json:"pageNo" d:"1" dc:"分页号码，默认1"`
	PageSize  int     `json:"pageSize" d:"10" dc:"分页数量，最大1000"`
}
type GetTaskAttachmentListRes struct {
	List     []*TaskAttachmentItem `json:"list" v:"required" dc:"任务附件列表"`
	PageNo   int                   `json:"pageNo" v:"required" dc:"页码"`
	PageSize int                   `json:"pageSize" v:"required" dc:"页码大小"`
	Total    int                   `json:"total" v:"required" dc:"总数"`
}
