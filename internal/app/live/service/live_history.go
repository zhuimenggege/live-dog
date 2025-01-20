// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/live"
)

type (
	ILiveHistory interface {
		List(ctx context.Context, req *v1.GetLiveHistoryListReq) (res *v1.GetLiveHistoryListRes, err error)
		Get(ctx context.Context, req *v1.GetLiveHistoryReq) (res *v1.GetLiveHistoryRes, err error)
		Add(ctx context.Context, req *v1.PostLiveHistoryReq) (res *v1.PostLiveHistoryRes, err error)
		Update(ctx context.Context, req *v1.PutLiveHistoryReq) (res *v1.PutLiveHistoryRes, err error)
		DeleteHistory(ctx context.Context, req *v1.DeleteLiveHistoryReq) (res *v1.DeleteLiveHistoryRes, err error)
		AddHistory(liveId int)
	}
)

var (
	localLiveHistory ILiveHistory
)

func LiveHistory() ILiveHistory {
	if localLiveHistory == nil {
		panic("implement not found for interface ILiveHistory, forgot register?")
	}
	return localLiveHistory
}

func RegisterLiveHistory(i ILiveHistory) {
	localLiveHistory = i
}
