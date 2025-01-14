package live

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/live"
	"github.com/shichen437/live-dog/internal/app/live/service"
)

type liveCookieController struct {
}

var LiveCookie = liveCookieController{}

func (c *liveCookieController) List(ctx context.Context, req *v1.GetCookieListReq) (res *v1.GetCookieListRes, err error) {
	res, err = service.LiveCookie().List(ctx, req)
	return
}

func (c *liveCookieController) Add(ctx context.Context, req *v1.PostLiveCookieReq) (res *v1.PostLiveCookieRes, err error) {
	res, err = service.LiveCookie().Add(ctx, req)
	return
}

func (c *liveCookieController) Delete(ctx context.Context, req *v1.DeleteLiveCookieReq) (res *v1.DeleteLiveCookieRes, err error) {
	res, err = service.LiveCookie().Delete(ctx, req)
	return
}

func (c *liveCookieController) Update(ctx context.Context, req *v1.PutLiveCookieReq) (res *v1.PutLiveCookieRes, err error) {
	res, err = service.LiveCookie().Update(ctx, req)
	return
}

func (c *liveCookieController) Get(ctx context.Context, req *v1.GetLiveCookieReq) (res *v1.GetLiveCookieRes, err error) {
	res, err = service.LiveCookie().Get(ctx, req)
	return
}
