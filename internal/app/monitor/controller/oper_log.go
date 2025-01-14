package monitor

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/monitor"
	"github.com/shichen437/live-dog/internal/app/monitor/service"
)

type operLogController struct {
}

var OperLog = operLogController{}

func (o *operLogController) GetOperLogList(ctx context.Context, req *v1.GetOperLogListReq) (res *v1.GetOperLogListRes, err error) {
	res, err = service.SysOperLog().List(ctx, req)
	return
}
