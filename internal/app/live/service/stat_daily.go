// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/live"
	"github.com/shichen437/live-dog/internal/app/live/model/entity"
)

type (
	IStatDaily interface {
		GetStatDailyById(ctx context.Context, id int64) (dailies *entity.StatDaily, err error)
		GetForUpdate(ctx context.Context, req *v1.GetStatDailyReq) (statDaily *v1.GetStatDailyRes, err error)
		GetStatDailyList(ctx context.Context, req *v1.GetStatDailyListReq) (dailyList *v1.GetStatDailyListRes, err error)
		Add(ctx context.Context, req *v1.PostStatDailyReq) (res *v1.PostStatDailyRes, err error)
		Update(ctx context.Context, req *v1.PutStatDailyReq) (res *v1.PutStatDailyRes, err error)
		Delete(ctx context.Context, req *v1.DeleteStatDailyReq) (res *v1.DeleteStatDailyRes, err error)
	}
)

var (
	localStatDaily IStatDaily
)

func StatDaily() IStatDaily {
	if localStatDaily == nil {
		panic("implement not found for interface IStatDaily, forgot register?")
	}
	return localStatDaily
}

func RegisterStatDaily(i IStatDaily) {
	localStatDaily = i
}
