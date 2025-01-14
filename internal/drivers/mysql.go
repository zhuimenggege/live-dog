package drivers

import (
	"time"

	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/v2/database/gdb"
)

func InitDockerDatabaseConfig() {
	n := gdb.ConfigNode{}
	n.Link = utils.DbLink
	n.Debug = false
	n.DryRun = false
	n.MaxIdleConnCount = 10
	n.MaxOpenConnCount = 100
	n.MaxConnLifeTime = time.Duration(30) * time.Second
	gdb.AddConfigNode("default", n)
}
