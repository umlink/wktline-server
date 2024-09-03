package common

import (
	"wktline-server/internal/service"
)

type sCommon struct{}

func init() {
	service.RegisterCommon(New())
}

func New() *sCommon {
	return &sCommon{}
}
