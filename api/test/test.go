// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package test

import (
	"context"
	
	"wktline-server/api/test/v1"
)

type ITestV1 interface {
	TestConnect(ctx context.Context, req *v1.TestConnectReq) (res *v1.TestConnectRes, err error)
}


