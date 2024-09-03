// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Project is the golang structure for table project.
type Project struct {
	Id          string      `json:"id"          orm:"id"          description:"项目 id"`
	Name        string      `json:"name"        orm:"name"        description:"项目名"`
	Description string      `json:"description" orm:"description" description:"项目描述"`
	HeaderImg   string      `json:"headerImg"   orm:"header_img"  description:"项目头图"`
	CreatorId   string      `json:"creatorId"   orm:"creator_id"  description:"创建人 id"`
	OwnerId     string      `json:"ownerId"     orm:"owner_id"    description:"项目所有者 id"`
	OperatorId  int64       `json:"operatorId"  orm:"operator_id" description:"项目操作者 id"`
	GroupId     string      `json:"groupId"     orm:"group_id"    description:"项目分组 id"`
	Status      int         `json:"status"      orm:"status"      description:"项目状态 1: 正常 2: 删除"`
	ShowType    string      `json:"showType"    orm:"show_type"   description:"显示类型 PUBLIC: 公开 PRIVATE: 私有"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:""`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:"删除shi"`
}
