package cmd

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/shichen437/live-dog/internal/app/monitor/service"
	"github.com/shichen437/live-dog/internal/pkg/crons"
)

func JobInit() {
	g.Log().Info(gctx.GetInitCtx(), "job monitor start!")

	jobs, err := service.SysJob().ListAll4Init(gctx.New())
	if err != nil {
		g.Log().Error(gctx.GetInitCtx(), "job monitor start error!", err)
		return
	}
	for _, job := range jobs {
		crons.AddSystemJob(job.JobId)
	}

	g.Log().Info(gctx.GetInitCtx(), "job monitor shutdown!")
}
