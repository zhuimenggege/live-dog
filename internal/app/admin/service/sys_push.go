// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "github.com/shichen437/live-dog/api/v1/admin"
	"github.com/shichen437/live-dog/internal/app/admin/model"
)

type (
	IPushChannel interface {
		Add(ctx context.Context, req *v1.PostPushChannelReq) (res *v1.PostPushChannelRes, err error)
		Delete(ctx context.Context, req *v1.DeletePushChannelReq) (res *v1.DeletePushChannelRes, err error)
		Update(ctx context.Context, req *v1.PutPushChannelReq) (res *v1.PutPushChannelRes, err error)
		List(ctx context.Context, req *v1.GetPushChannelListReq) (res *v1.GetPushChannelListRes, err error)
		Get(ctx context.Context, req *v1.GetPushChannelReq) (res *v1.GetPushChannelRes, err error)
		ListAll(ctx context.Context) (res []*model.PushChannel, err error)
	}
)

var (
	localPushChannel IPushChannel
)

func PushChannel() IPushChannel {
	if localPushChannel == nil {
		panic("implement not found for interface IPushChannel, forgot register?")
	}
	return localPushChannel
}

func RegisterPushChannel(i IPushChannel) {
	localPushChannel = i
}
