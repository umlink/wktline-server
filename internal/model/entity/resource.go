// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Resource is the golang structure for table resource.
type Resource struct {
	Id         string      `json:"id"         orm:"id"          description:""`
	Name       string      `json:"name"       orm:"name"        description:""`
	Url        string      `json:"url"        orm:"url"         description:""`
	Type       string      `json:"type"       orm:"type"        description:""`
	Size       int64       `json:"size"       orm:"size"        description:""`
	Hash       string      `json:"hash"       orm:"hash"        description:""`
	BucketHash string      `json:"bucketHash" orm:"bucket_hash" description:""`
	CreatorId  int64       `json:"creatorId"  orm:"creator_id"  description:""`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:""`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:""`
}
