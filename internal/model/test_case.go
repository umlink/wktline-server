package model

import v1 "wktline-server/api/test_case/v1"

type CreateTestCaseInput struct {
	Id          string
	Name        string
	Description string
	ProjectId   string
}

type UpdateTestCaseInput struct {
	Id          string
	Name        *string
	Description *string
	Value       *string
	Progress    *int
	Status      *string
}

type AddTestCaseActorInput struct {
	Id       string
	ActorIds []string
}
type DelTestCaseActorInput struct {
	Id       string
	ActorIds []string
}
type GetTestCaseListInput struct {
	ProjectId string
	CreatorId string
	PageSize  int
	PageNo    int
}
type GetTestCaseListOutput v1.GetTestCaseListRes
type GetTestCaseDetailOutput v1.GetTestCaseRes
