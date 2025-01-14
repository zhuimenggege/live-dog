package monitor

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/monitor"
	"github.com/shichen437/live-dog/internal/app/monitor/service"
)

type serverController struct {
}

var ServerInfo = serverController{}

func (o *serverController) GetServerInfo(ctx context.Context, req *v1.GetServerInfoReq) (res *v1.GetServerInfoRes, err error) {
	return service.ServerInfo().GetServerInfo(ctx, req)
}
