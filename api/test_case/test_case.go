// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package test_case

import (
	"context"
	
	"wktline-server/api/test_case/v1"
)

type ITestCaseV1 interface {
	AddTestCaseActors(ctx context.Context, req *v1.AddTestCaseActorsReq) (res *v1.AddTestCaseActorsRes, err error)
	DelTestCaseActors(ctx context.Context, req *v1.DelTestCaseActorsReq) (res *v1.DelTestCaseActorsRes, err error)
	GetTestCase(ctx context.Context, req *v1.GetTestCaseReq) (res *v1.GetTestCaseRes, err error)
	CreateTestCase(ctx context.Context, req *v1.CreateTestCaseReq) (res *v1.CreateTestCaseRes, err error)
	DelTestCase(ctx context.Context, req *v1.DelTestCaseReq) (res *v1.DelTestCaseRes, err error)
	UpdateTestCase(ctx context.Context, req *v1.UpdateTestCaseReq) (res *v1.UpdateTestCaseRes, err error)
	GetTestCaseList(ctx context.Context, req *v1.GetTestCaseListReq) (res *v1.GetTestCaseListRes, err error)
	AddTestCaseLock(ctx context.Context, req *v1.AddTestCaseLockReq) (res *v1.AddTestCaseLockRes, err error)
	DelTestCaseLock(ctx context.Context, req *v1.DelTestCaseLockReq) (res *v1.DelTestCaseLockRes, err error)
}


