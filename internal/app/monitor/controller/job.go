package monitor

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/monitor"
	"github.com/shichen437/live-dog/internal/app/monitor/service"
)

type sysJobController struct {
}

var SysJob = sysJobController{}

func (o *sysJobController) GetJobList(ctx context.Context, req *v1.GetJobListReq) (res *v1.GetJobListRes, err error) {
	res, err = service.SysJob().List(ctx, req)
	return
}

func (o *sysJobController) GetJobLogList(ctx context.Context, req *v1.GetJobLogListReq) (res *v1.GetJobLogListRes, err error) {
	res, err = service.SysJob().LogList(ctx, req)
	return
}
