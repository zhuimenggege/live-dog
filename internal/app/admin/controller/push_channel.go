package admin

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/admin"
	"github.com/shichen437/live-dog/internal/app/admin/service"
)

type pushChannelController struct {
}

var PushChannel = pushChannelController{}

func (c *pushChannelController) Add(ctx context.Context, req *v1.PostPushChannelReq) (res *v1.PostPushChannelRes, err error) {
	return service.PushChannel().Add(ctx, req)
}

func (c *pushChannelController) Delete(ctx context.Context, req *v1.DeletePushChannelReq) (res *v1.DeletePushChannelRes, err error) {
	return service.PushChannel().Delete(ctx, req)
}

func (c *pushChannelController) List(ctx context.Context, req *v1.GetPushChannelListReq) (res *v1.GetPushChannelListRes, err error) {
	return service.PushChannel().List(ctx, req)
}

func (c *pushChannelController) Update(ctx context.Context, req *v1.PutPushChannelReq) (res *v1.PutPushChannelRes, err error) {
	return service.PushChannel().Update(ctx, req)
}

func (c *pushChannelController) Get(ctx context.Context, req *v1.GetPushChannelReq) (res *v1.GetPushChannelRes, err error) {
	return service.PushChannel().Get(ctx, req)
}
