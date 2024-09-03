package v1

import "github.com/gogf/gf/v2/frame/g"

type GetTestCaseReq struct {
	g.Meta    `path:"/test-case/getCaseDetail" method:"get" summary:"测试用例详情" tags:"TestCase"`
	Id        string `json:"id" v:"required" dc:"用例 id"`
	ProjectId string `json:"projectId" v:"required" dc:"项目id"`
}
type GetTestCaseRes struct {
	Id            string   `json:"id" v:"required" dc:"用例id"`
	Name          string   `json:"name" v:"required" dc:"用例名"`
	Description   string   `json:"description" v:"required" dc:"描述说明"`
	Value         string   `json:"value" v:"required" dc:"数据"`
	Progress      int      `json:"progress" v:"required" dc:"进度(草稿: DRAFT 发布:PUBLISHED 测试中:TESTING 测试完成:DONE)"`
	Status        string   `json:"status" v:"required" dc:"状态"`
	CreatorId     string   `json:"creatorId" v:"required" dc:"创建人id"`
	CreatorName   string   `json:"creatorName" v:"required" dc:"创建人名称"`
	CreatorAvatar string   `json:"creatorAvatar" v:"required" dc:"创建人头像"`
	ActorIds      []string `json:"actorIds" v:"required" dc:"参与人"`
	EditorId      string   `json:"editorId" v:"required" dc:"正在编辑的用户（lock）"`
	EditorName    string   `json:"editorName" v:"required" dc:"编辑人名称"`
	EditorAvatar  string   `json:"editorAvatar" v:"required" dc:"编辑人头像"`
}

type CreateTestCaseReq struct {
	g.Meta      `path:"/test-case/createCase" method:"post" summary:"创建测试用例" tags:"TestCase"`
	ProjectId   string `json:"projectId" v:"required" dc:"项目id"`
	Name        string `json:"name" v:"required" dc:"测试用例名"`
	Description string `json:"description" v:"required" dc:"关于测试用例的描述"`
}
type CreateTestCaseRes struct{}

type DelTestCaseReq struct {
	g.Meta    `path:"/test-case/dekCase" method:"post" summary:"删除测试用例" tags:"TestCase"`
	Id        string `json:"Id" v:"required" dc:"用例 id"`
	ProjectId string `json:"projectId" v:"required" dc:"项目id"`
}
type DelTestCaseRes struct{}

type UpdateTestCaseReq struct {
	g.Meta      `path:"/test-case/updateCase" method:"post" summary:"更新测试用例" tags:"TestCase"`
	Id          string  `json:"id" v:"required" dc:"用例id"`
	ProjectId   string  `json:"projectId" v:"required" dc:"项目id"`
	Name        *string `json:"name" dc:"测试用例名"`
	Status      *string `json:"status" dc:"状态(草稿: DRAFT 发布:PUBLISHED 测试中:TESTING 测试完成:DONE)"`
	Progress    *int    `json:"progress" dc:"测试进度(1-100)"`
	Description *string `json:"description" dc:"关于测试用例的描述"`
	Value       *string `json:"value" v:"required" dc:"数据"`
}
type UpdateTestCaseRes struct{}

type GetTestCaseListReq struct {
	g.Meta    `path:"/test-case/getCaseList" method:"post" summary:"获取测试用例列表" tags:"TestCase"`
	ProjectId string `json:"projectId" v:"required" dc:"项目id"`
	CreatorId string `json:"creatorId" dc:"创建人id"`
	PageNo    int    `json:"pageNo" d:"1" dc:"分页号码，默认1"`
	PageSize  int    `json:"pageSize" d:"10" dc:"分页数量，最大1000"`
}
type TestCaseItem struct {
	Id            string   `json:"id" v:"required" dc:"用例id"`
	Name          string   `json:"name" v:"required" dc:"用例名"`
	Description   string   `json:"description" v:"required" dc:"描述说明"`
	Progress      int      `json:"progress" v:"required" dc:"进度(草稿: DRAFT 发布:PUBLISHED 测试中:TESTING 测试完成:DONE)"`
	Status        string   `json:"status" v:"required" dc:"状态"`
	CreatorId     string   `json:"creatorId" v:"required" dc:"创建人id"`
	CreatorName   string   `json:"creatorName" v:"required" dc:"创建人名称"`
	CreatorAvatar string   `json:"creatorAvatar" v:"required" dc:"创建人头像"`
	EditorId      string   `json:"editorId" v:"required" dc:"正在编辑的用户（lock）"`
	EditorName    string   `json:"editorName" v:"required" dc:"编辑人名称"`
	EditorAvatar  string   `json:"editorAvatar" v:"required" dc:"编辑人头像"`
	ActorIds      []string `json:"actorIds" v:"required" dc:"参与人"`
	CreatedAt     string   `json:"createdAt" v:"required" dc:"创建时间"`
	UpdatedAt     string   `json:"updatedAt" v:"required" dc:"更新时间"`
}
type GetTestCaseListRes struct {
	List     []*TestCaseItem `json:"list" v:"required" dc:"用例"`
	PageNo   int             `json:"pageNo" v:"required" dc:"页码"`
	PageSize int             `json:"pageSize" v:"required" dc:"页码大小"`
	Total    int             `json:"total" v:"required" dc:"总数"`
}
type AddTestCaseLockReq struct {
	g.Meta `path:"/test-case/addCaseLock" method:"post" summary:"添加修改锁" tags:"TestCase"`
	Id     string `json:"Id" v:"required" dc:"用例 id"`
}
type AddTestCaseLockRes struct{}

type DelTestCaseLockReq struct {
	g.Meta `path:"/test-case/delCaseLock" method:"post" summary:"删除修改锁" tags:"TestCase"`
	Id     string `json:"Id" v:"required" dc:"用例 id"`
}
type DelTestCaseLockRes struct{}
