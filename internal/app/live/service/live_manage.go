// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "github.com/shichen437/live-dog/api/v1/live"
	"github.com/shichen437/live-dog/internal/pkg/lives"
)

type (
	ILiveManage interface {
		AddLiveManage(ctx context.Context, req *v1.PostLiveManageReq) (res *v1.PostLiveManageRes, err error)
		UpdateLiveManage(ctx context.Context, req *v1.PutLiveManageReq) (res *v1.PutLiveManageRes, err error)
		GetRoomInfoList(ctx context.Context, req *v1.GetRoomInfoListReq) (res *v1.GetRoomInfoListRes, err error)
		GetLiveManage(ctx context.Context, req *v1.GetLiveManageReq) (res *v1.GetLiveManageRes, err error)
		DeleteLiveManage(ctx context.Context, req *v1.DeleteLiveManageReq) (res *v1.DeleteLiveManageRes, err error)
		GetLiveModels4Init(ctx context.Context) (liveModels []*lives.LiveModel, err error)
	}
)

var (
	localLiveManage ILiveManage
)

func LiveManage() ILiveManage {
	if localLiveManage == nil {
		panic("implement not found for interface ILiveManage, forgot register?")
	}
	return localLiveManage
}

func RegisterLiveManage(i ILiveManage) {
	localLiveManage = i
}
