package drivers

import (
	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {
	if !utils.IsDocker() {
		return
	}
	g.Log().Info(gctx.GetInitCtx(), "init docker environment")
	InitDockerDatabaseConfig()
	g.Log().Info(gctx.GetInitCtx(), "init docker environment finished")
}
