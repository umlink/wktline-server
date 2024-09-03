package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type ProjectInfoItem struct {
	Id          string `json:"id" v:"required" dc:"项目id"`
	Name        string `json:"name" v:"required" dc:"项目名"`
	Description string `json:"description" v:"required" dc:"项目描述"`
	HeaderImg   string `json:"headerImg" v:"required" dc:"项目头图"`
	OwnerId     string `json:"ownerId" v:"required" dc:"项目负责人id"`
	OwnerName   string `json:"ownerName" v:"required" dc:"项目负责人name"`
	OwnerAvatar string `json:"ownerAvatar" v:"required" dc:"项目负责人头像"`
	GroupId     string `json:"groupId" v:"required" dc:"项目分组id"`
	GroupName   string `json:"groupName" v:"required" dc:"项目分组name"`
	ShowType    string `json:"showType" v:"required" dc:"项目类型 PRIVATE: 私有  PUBLIC：公开"`
}

type GetProjectInfoReq struct {
	g.Meta `path:"/project/getProjectInfo" method:"get" summary:"获取项目详情" tags:"Project"`
	Id     string `json:"id" v:"required#项目id不能为空" dc:"项目 id"`
}

type GetProjectInfoRes struct {
	Id           string    `json:"id" v:"required" dc:"项目id"`
	Name         string    `json:"name" v:"required" dc:"项目名"`
	Description  string    `json:"description" v:"required" dc:"项目描述"`
	HeaderImg    string    `json:"headerImg" v:"required" dc:"项目头图"`
	OperatorId   string    `json:"operatorId" v:"required" dc:"操作者id"`
	OperatorName string    `json:"operatorName" v:"required" dc:"操作者name"`
	OwnerId      string    `json:"ownerId" v:"required" dc:"项目负责人id"`
	OwnerName    string    `json:"ownerName" v:"required" dc:"项目负责人name"`
	OwnerAvatar  string    `json:"ownerAvatar" v:"required" dc:"项目负责人avatar"`
	GroupId      string    `json:"groupId" v:"required" dc:"项目分组id"`
	GroupName    string    `json:"groupName" v:"required" dc:"项目分组name"`
	ShowType     string    `json:"showType" v:"required" dc:"项目类型 PRIVATE: 私有  PUBLIC：公开"`
	Status       int       `json:"status" v:"required" dc:"项目状态 1. 正常 2：回收站 3：已删除"`
	IsJoined     bool      `json:"isJoined" d:"false" v:"required" dc:"是否加入项目"`
	CanEdit      bool      `json:"canEdit" d:"false" v:"required" dc:"是否可编辑"`
	CreatedAt    time.Time `json:"createdAt" v:"required" dc:"创建时间"`
	UpdatedAt    time.Time `json:"updatedAt" v:"required" dc:"更新时间"`
}

type CreateProjectReq struct {
	g.Meta      `path:"/project/createProject" method:"post" summary:"新建项目" tags:"Project"`
	Name        string `json:"name" v:"required#项目名不能为空" dc:"项目名"`
	Description string `json:"description" dc:"项目描述"`
	HeaderImg   string `json:"headerImg" dc:"项目头图"`
	GroupId     string `json:"groupId" v:"required#项目分组不能为空" dc:"项目分组id"`
	ShowType    string `json:"showType" d:"PUBLIC" dc:"项目类型 PRIVATE: 私有  PUBLIC：公开"`
}
type CreateProjectRes struct {
	Id string `json:"id" v:"required" dc:"项目 id"`
}

type DelProjectReq struct {
	g.Meta    `path:"/project/delProject" method:"post" summary:"删除项目" tags:"Project"`
	ProjectId string `json:"projectId" v:"required#项目id不能为空" dc:"项目 id"`
}
type DelProjectRes struct{}

type UpdateProjectReq struct {
	g.Meta      `path:"/project/updateProject" method:"post" summary:"更新项目" tags:"Project"`
	ProjectId   string  `json:"projectId" v:"required#项目id不能为空" dc:"项目 id"`
	Name        *string `json:"name" dc:"项目名"`
	Description *string `json:"description" dc:"项目描述"`
	HeaderImg   *string `json:"headerImg" dc:"项目头图"`
	OwnerId     *string `json:"ownerId" dc:"项目负责人id"`
	GroupId     *string `json:"groupId" dc:"项目分组id"`
	ShowType    *string `json:"showType" dc:"项目类型 PRIVATE: 私有  PUBLIC：公开"`
	Status      *int    `json:"status" dc:"项目状态 1. 正常 2：回收站 3：已删除"`
}
type UpdateProjectRes struct{}

type GetProjectListReq struct {
	g.Meta    `path:"/project/getProjectList" method:"post" summary:"获取项目列表" tags:"Project"`
	Keywords  *string `json:"keywords" dc:"项目名-支持模糊搜索"`
	IsOwner   bool    `json:"isOwner" d:"false" dc:"我负责的"`
	IsCreator bool    `json:"isCreator" d:"false" dc:"我创建的"`
	IsJoined  bool    `json:"isJoined" d:"false" dc:"是否加入项目"`
	GroupId   *string `json:"groupId" dc:"项目分组id"`
	ShowType  *string `json:"showType" dc:"显示类型-只有系统管理员和我的全部（全部包含已加入的私有项目)）"`
	Status    *int    `json:"status" dc:"项目状态 1. 正常 2：回收站 3：已删除"`
	PageNo    int     `json:"pageNo" d:"1" dc:"分页号码，默认1"`
	PageSize  int     `json:"pageSize" d:"10" dc:"分页数量，最大100"`
}
type GetProjectListRes struct {
	List     []*ProjectInfoItem `json:"list" v:"required" dc:"项目列表"`
	PageNo   int                `json:"pageNo" v:"required" dc:"页码"`
	PageSize int                `json:"pageSize" v:"required" dc:"页码大小"`
	Total    int                `json:"total" v:"required" dc:"总数"`
}
