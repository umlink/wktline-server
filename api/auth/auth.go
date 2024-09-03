// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"wktline-server/api/auth/v1"
)

type IAuthV1 interface {
	Auth(ctx context.Context, req *v1.AuthReq) (res *v1.AuthRes, err error)
}
