package recorders

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shichen437/live-dog/internal/app/live/dao"
	"github.com/shichen437/live-dog/internal/app/live/model/do"
	"github.com/shichen437/live-dog/internal/pkg/lives"
	"github.com/shichen437/live-dog/internal/pkg/message_push"
	"github.com/shichen437/live-dog/internal/pkg/utils"
)

func liveStartBiz(ctx context.Context, liveId int) {
	go message_push.LivePush(ctx, liveId)
	utils.GetGlobal(ctx).StartTimeMap[liveId] = gtime.Now()
}

func liveEndBiz(ctx context.Context, liveId int) {
	addHistory(ctx, liveId)
}

func addHistory(ctx context.Context, liveId int) {
	if liveId == 0 {
		return
	}
	global := utils.GetGlobal(ctx)
	m, ok := global.StartTimeMap[liveId]
	if !ok || m == nil {
		return
	}
	endTime := gtime.Now()
	dao.LiveHistory.Ctx(ctx).Insert(do.LiveHistory{
		LiveId:    liveId,
		StartTime: m,
		EndTime:   endTime,
		Duration:  fmt.Sprintf("%.2f", endTime.Sub(m).Hours()),
	})
}

func (*manager) updateName(ctx context.Context, live lives.Live) {
	roomInfo, err := live.GetInfo()
	if err == nil {
		g.Try(ctx, func(ctx context.Context) {
			//更新房间信息
			_, e := dao.RoomInfo.Ctx(ctx).Where(dao.RoomInfo.Columns().LiveId, live.GetLiveId()).Update(do.RoomInfo{
				Anchor:     roomInfo.Anchor,
				RoomName:   roomInfo.RoomName,
				ActionTime: gtime.Now(),
			})
			if e == nil {
				global := utils.GetGlobal(ctx)
				_, ok := global.ModelsMap[live.GetLiveId()]
				if ok {
					global.ModelsMap[live.GetLiveId()].RoomInfo.Anchor = roomInfo.Anchor
					global.ModelsMap[live.GetLiveId()].RoomInfo.RoomName = roomInfo.RoomName
				}
			}
		})
	}
}
