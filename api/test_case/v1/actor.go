package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AddTestCaseActorsReq struct {
	g.Meta    `path:"/test-case/addCaseActors" method:"post" summary:"添加测试用例参与人" tags:"TestCase"`
	Id        string   `json:"Id" v:"required" dc:"用例 id"`
	ProjectId string   `json:"projectId" v:"required" dc:"项目id"`
	ActorIds  []string `json:"actorIds" v:"required" dc:"参与人id列表"`
}
type AddTestCaseActorsRes struct{}

type DelTestCaseActorsReq struct {
	g.Meta    `path:"/test-case/delCaseActors" method:"post" summary:"添加测试用例参与人" tags:"TestCase"`
	Id        string   `json:"Id" v:"required" dc:"用例 id"`
	ProjectId string   `json:"projectId" v:"required" dc:"项目id"`
	ActorIds  []string `json:"actorIds" v:"required" dc:"参与人id列表"`
}
type DelTestCaseActorsRes struct{}
