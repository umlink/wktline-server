package model

import v1 "wktline-server/api/message/v1"

type GetMyMessageDetailInput struct {
	Id string
}
type AddMessageInput struct {
	Id          string
	Title       string
	Content     string
	ContentType string
	Type        string
	ProjectId   string
	TaskId      string
	ReceiverId  string
	SenderId    string
}
type CreateProjectActionUsersMsgInput struct {
	Title     string
	ProjectId string
	UserIds   []string
	MsgType   string
}
type CreateTaskActionUsersMsgInput struct {
	Title     string
	TaskId    string
	ProjectId string
	UserIds   []string
	MsgType   string
}
type ReadMessageInput struct {
	Id string
}
type GetMessageListInput struct {
	Type      string
	Status    string
	ProjectId string
	PageNo    int
	PageSize  int
}
type GetMessageListOutput v1.GetMessageListRes

type GetMessageDetailOutput v1.GetMessageDetailRes
