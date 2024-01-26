package main

import (
	_ "SheeDrive/internal/packed"

	_ "SheeDrive/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"SheeDrive/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
