package live

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/live"
	"github.com/shichen437/live-dog/internal/app/live/service"
)

type roomManageController struct {
}

var LiveManage = roomManageController{}

func (c *roomManageController) List(ctx context.Context, req *v1.GetRoomInfoListReq) (res *v1.GetRoomInfoListRes, err error) {
	res, err = service.LiveManage().GetRoomInfoList(ctx, req)
	return
}

func (c *roomManageController) Add(ctx context.Context, req *v1.PostLiveManageReq) (res *v1.PostLiveManageRes, err error) {
	res, err = service.LiveManage().AddLiveManage(ctx, req)
	return
}

func (c *roomManageController) Delete(ctx context.Context, req *v1.DeleteLiveManageReq) (res *v1.DeleteLiveManageRes, err error) {
	res, err = service.LiveManage().DeleteLiveManage(ctx, req)
	return
}

func (c *roomManageController) Update(ctx context.Context, req *v1.PutLiveManageReq) (res *v1.PutLiveManageRes, err error) {
	res, err = service.LiveManage().UpdateLiveManage(ctx, req)
	return
}

func (c *roomManageController) Get(ctx context.Context, req *v1.GetLiveManageReq) (res *v1.GetLiveManageRes, err error) {
	res, err = service.LiveManage().GetLiveManage(ctx, req)
	return
}
