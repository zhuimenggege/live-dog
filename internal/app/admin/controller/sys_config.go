package admin

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/admin"
	"github.com/shichen437/live-dog/internal/app/admin/service"
)

type sysConfigController struct {
}

var SysConfig = sysConfigController{}

func (s *sysConfigController) GetSysConfigList(ctx context.Context, req *v1.GetSysConfigListReq) (res *v1.GetSysConfigListRes, err error) {
	res, err = service.SysConfig().GetSysConfigList(ctx, req)
	return
}

func (s *sysConfigController) Add(ctx context.Context, req *v1.PostSysConfigReq) (res *v1.PostSysConfigRes, err error) {
	res, err = service.SysConfig().Add(ctx, req)
	return
}

func (s *sysConfigController) Update(ctx context.Context, req *v1.PutSysConfigReq) (res *v1.PutSysConfigRes, err error) {
	res, err = service.SysConfig().Update(ctx, req)
	return
}

func (s *sysConfigController) Delete(ctx context.Context, req *v1.DeleteSysConfigReq) (res *v1.DeleteSysConfigRes, err error) {
	res, err = service.SysConfig().Delete(ctx, req)
	return
}

func (s *sysConfigController) GetSysConfig(ctx context.Context, req *v1.GetSysConfigReq) (res *v1.GetSysConfigRes, err error) {
	res, err = service.SysConfig().GetSysConfig(ctx, req)
	return
}
