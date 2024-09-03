package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	_ "wktline-server/internal/packed"

	_ "wktline-server/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	_ "wktline-server/internal/boot"
	"wktline-server/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
