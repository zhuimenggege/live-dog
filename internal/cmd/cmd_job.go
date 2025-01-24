package cmd

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func Job() {
	g.Log().Info(gctx.GetInitCtx(), "job monitor start!")
}
