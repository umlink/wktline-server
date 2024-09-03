package v1

import "github.com/gogf/gf/v2/frame/g"

type GenProjectInviteCodeReq struct {
	g.Meta         `path:"/project/genInviteCode" method:"post" summary:"生成项目邀请链接" tags:"ProjectInvite"`
	ProjectId      string `json:"projectId" v:"required" dc:"项目id"`
	Deadline       string `json:"deadline" v:"required:datetime" dc:"链接生效截止时间"`
	MaxInviteCount int    `json:"maxInviteCount" dc:"最大可邀请人数"`
}
type GenProjectInviteCodeRes struct {
	Id string `json:"code" v:"required" dc:"邀请码"`
}

type GetProjectInviteInfoReq struct {
	g.Meta `path:"/project/getInviteInfo" method:"get" summary:"根据邀请码获取邀请信息" tags:"ProjectInvite"`
	Code   string `json:"code" v:"required" dc:"邀请码"`
}
type GetProjectInviteInfoRes struct {
	Id             string `json:"code" v:"required" dc:"邀请码"`
	ProjectId      string `json:"projectId" v:"required" dc:"项目id"`
	ProjectName    string `json:"projectName" v:"required" dc:"项目名"`
	InviterId      string `json:"inviterId" v:"required" dc:"发起邀请的用户id"`
	InviterName    string `json:"inviterName" v:"required" dc:"发起邀请的用户名"`
	InviterAvatar  string `json:"inviterAvatar" v:"required" dc:"发起邀请的用户头像"`
	JoinedCount    int    `json:"joinedCount" v:"required" dc:"已通过链接加入的人数"`
	Deadline       string `json:"deadline" v:"required:datetime" dc:"链接生效截止时间"`
	MaxInviteCount int    `json:"maxInviteCount" dc:"最大可邀请人数"`
	Joined         bool   `json:"joined" dc:"是否已加入"`
}

type InJoinInviteProjectReq struct {
	g.Meta `path:"/project/inJoinInvite" method:"post" summary:"加入项目" tags:"ProjectInvite"`
	Code   string `json:"code" v:"required" dc:"邀请码"`
}
type InJoinInviteProjectRes struct{}

type DelInviteCodeReq struct {
	g.Meta `path:"/project/delInviteCode" method:"post" summary:"删除邀请码" tags:"ProjectInvite"`
	Code   string `json:"code" v:"required" dc:"邀请码"`
}
type DelInviteCodeRes struct{}

type GetProjectMyInviteListReq struct {
	g.Meta    `path:"/project/getInviteList" method:"post" summary:"获取项目下的邀请地址列表" tags:"ProjectInvite"`
	ProjectId string `json:"projectId" v:"required" dc:"项目 id"`
}
type ProjectMyInviteItem struct {
	Id             string `json:"code" v:"required" dc:"邀请码"`
	JoinedCount    int    `json:"joinedCount" v:"required" dc:"已通过链接加入的人数"`
	Deadline       string `json:"deadline" v:"required:datetime" dc:"链接生效截止时间"`
	MaxInviteCount int    `json:"maxInviteCount" dc:"最大可邀请人数"`
}
type GetProjectMyInviteListRes []*ProjectMyInviteItem
