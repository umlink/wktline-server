// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"wktline-server/internal/model"
	"wktline-server/internal/model/entity"
)

type (
	ITestCase interface {
		AddTestCaseUser(ctx context.Context, in model.AddTestCaseActorInput) (err error)
		DelTestCaseUser(ctx context.Context, in model.DelTestCaseActorInput) (err error)
		AddTestCaseLock(ctx context.Context, id string) (err error)
		DelTestCaseLock(ctx context.Context, id string) (err error)
		GetTestCaseBaseInfo(ctx context.Context, id string) (out entity.TestCase, err error)
		CreateTestCase(ctx context.Context, in model.CreateTestCaseInput) (err error)
		UpdateTestCase(ctx context.Context, in model.UpdateTestCaseInput) (err error)
		DelTestCase(ctx context.Context, id string) (err error)
		GetTestCaseList(ctx context.Context, in model.GetTestCaseListInput) (res *model.GetTestCaseListOutput, err error)
		GetTestCaseDetail(id string) (res *model.GetTestCaseDetailOutput, err error)
	}
)

var (
	localTestCase ITestCase
)

func TestCase() ITestCase {
	if localTestCase == nil {
		panic("implement not found for interface ITestCase, forgot register?")
	}
	return localTestCase
}

func RegisterTestCase(i ITestCase) {
	localTestCase = i
}
