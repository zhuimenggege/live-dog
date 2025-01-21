package live

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/live"
	"github.com/shichen437/live-dog/internal/app/live/service"
)

type liveHistoryController struct {
}

var LiveHistory = liveHistoryController{}

func (c *liveHistoryController) List(ctx context.Context, req *v1.GetLiveHistoryListReq) (res *v1.GetLiveHistoryListRes, err error) {
	res, err = service.LiveHistory().List(ctx, req)
	return
}

func (c *liveHistoryController) Get(ctx context.Context, req *v1.GetLiveHistoryReq) (res *v1.GetLiveHistoryRes, err error) {
	res, err = service.LiveHistory().Get(ctx, req)
	return
}

func (c *liveHistoryController) Add(ctx context.Context, req *v1.PostLiveHistoryReq) (res *v1.PostLiveHistoryRes, err error) {
	res, err = service.LiveHistory().Add(ctx, req)
	return
}

func (c *liveHistoryController) Update(ctx context.Context, req *v1.PutLiveHistoryReq) (res *v1.PutLiveHistoryRes, err error) {
	res, err = service.LiveHistory().Update(ctx, req)
	return
}

func (c *liveHistoryController) Delete(ctx context.Context, req *v1.DeleteLiveHistoryReq) (res *v1.DeleteLiveHistoryRes, err error) {
	res, err = service.LiveHistory().Delete(ctx, req)
	return
}
