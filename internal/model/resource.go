package model

import (
	"time"
)

type ResourceItem struct {
	Id          string
	Name        string
	Hash        string
	Size        int64
	Url         string
	Type        string
	CreatorId   string
	CreatedTime time.Time
	UpdatedTime time.Time
}

type CreateResourceInput struct {
	Id         string
	Name       string
	Hash       string
	Size       int64
	Url        string
	Type       string
	BucketHash string
	CreatorId  string
}
type CreateResourceOutput struct {
	Id string
}

type CheckHasResourceOutput = ResourceItem
