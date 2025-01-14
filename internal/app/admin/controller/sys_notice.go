package admin

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/admin"
	"github.com/shichen437/live-dog/internal/app/admin/service"
)

type sysNoticeController struct {
}

var SysNotice = sysNoticeController{}

func (s *sysNoticeController) GetSysNotice(ctx context.Context, req *v1.GetSysNoticeDetailReq) (res *v1.GetSysNoticeDetailRes, err error) {
	res, err = service.SysNotice().GetNoticeData(ctx, req)
	return
}

func (s *sysNoticeController) GetSysNoticeList(ctx context.Context, req *v1.GetSysNoticeListReq) (res *v1.GetSysNoticeListRes, err error) {
	res, err = service.SysNotice().GetSysNoticeList(ctx, req)
	return
}

func (s *sysNoticeController) Add(ctx context.Context, req *v1.PostSysNoticeReq) (res *v1.PostSysNoticeRes, err error) {
	res, err = service.SysNotice().Add(ctx, req)
	return
}

func (s *sysNoticeController) Update(ctx context.Context, req *v1.PutSysNoticeReq) (res *v1.PutSysNoticeRes, err error) {
	res, err = service.SysNotice().Update(ctx, req)
	return
}

func (s *sysNoticeController) Delete(ctx context.Context, req *v1.DeleteSysNoticeReq) (res *v1.DeleteSysNoticeRes, err error) {
	res, err = service.SysNotice().Delete(ctx, req)
	return
}
