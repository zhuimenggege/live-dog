package live_history

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	v1 "github.com/shichen437/live-dog/api/v1/live"
	"github.com/shichen437/live-dog/internal/app/common/consts"
	"github.com/shichen437/live-dog/internal/app/live/dao"
	"github.com/shichen437/live-dog/internal/app/live/model"
	"github.com/shichen437/live-dog/internal/app/live/model/do"
	"github.com/shichen437/live-dog/internal/app/live/service"
	"github.com/shichen437/live-dog/internal/pkg/utils"
)

func init() {
	service.RegisterLiveHistory(New())
}

func New() *sLiveHistory {
	return &sLiveHistory{}
}

type sLiveHistory struct {
}

func (s *sLiveHistory) List(ctx context.Context, req *v1.GetLiveHistoryListReq) (res *v1.GetLiveHistoryListRes, err error) {
	res = &v1.GetLiveHistoryListRes{}
	m := g.Model("live_history h").
		InnerJoin("live_manage l", "h.live_id=l.id").
		InnerJoin("room_info r", "l.id=r.live_id").
		OrderDesc("h.start_time")
	if req.LiveId != 0 {
		m = m.Where("l.id", req.LiveId)
	}
	if req.Anchor != "" {
		m = m.WhereLike("r.anchor", "%"+req.Anchor+"%")
	}
	err = g.Try(ctx, func(ctx context.Context) {
		var result []*model.LiveHistory
		res.Total, err = m.Count()
		err = m.Fields("h.*, r.anchor").Page(req.PageNum, req.PageSize).Scan(&result)
		utils.WriteErrLogT(ctx, err, consts.ListF)
		res.Rows = result
	})
	return
}

func (s *sLiveHistory) Get(ctx context.Context, req *v1.GetLiveHistoryReq) (res *v1.GetLiveHistoryRes, err error) {
	res = &v1.GetLiveHistoryRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		err = g.Model("live_history h").
			InnerJoin("live_manage l", "h.live_id=l.id").
			InnerJoin("room_info r", "l.id=r.live_id").
			Fields("h.*, r.anchor").
			Where("h.id", req.Id).
			Scan(&res)
		utils.WriteErrLogT(ctx, err, consts.GetF)
	})
	return
}

func (s *sLiveHistory) Add(ctx context.Context, req *v1.PostLiveHistoryReq) (res *v1.PostLiveHistoryRes, err error) {
	if req.LiveId == 0 {
		err = gerror.Newf(`直播名称不能为空`)
		return
	}
	if err = validTime(req.StartTime, req.EndTime); err != nil {
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		//添加角色信息
		_, err := dao.LiveHistory.Ctx(ctx).Insert(do.LiveHistory{
			LiveId:    req.LiveId,
			StartTime: req.StartTime,
			EndTime:   req.EndTime,
			Duration:  fmt.Sprintf("%.2f", req.EndTime.Sub(req.StartTime).Hours()),
		})
		utils.WriteErrLogT(ctx, err, consts.AddF)
	})
	return
}

func (s *sLiveHistory) Update(ctx context.Context, req *v1.PutLiveHistoryReq) (res *v1.PutLiveHistoryRes, err error) {
	res = &v1.PutLiveHistoryRes{}
	if req.Id == 0 {
		return nil, utils.TError(ctx, consts.IDEmpty)
	}
	if err = validTime(req.StartTime, req.EndTime); err != nil {
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.LiveHistory.Ctx(ctx).WherePri(req.Id).Update(do.LiveHistory{
			StartTime: req.StartTime,
			EndTime:   req.EndTime,
			Duration:  fmt.Sprintf("%.2f", req.EndTime.Sub(req.StartTime).Hours()),
		})
		utils.WriteErrLogT(ctx, e, consts.UpdateF)
	})
	return
}

func (s *sLiveHistory) DeleteHistory(ctx context.Context, req *v1.DeleteLiveHistoryReq) (res *v1.DeleteLiveHistoryRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		ids := utils.ParamStrToSlice(req.Id, ",")
		_, e := dao.LiveHistory.Ctx(ctx).WhereIn(dao.LiveHistory.Columns().Id, ids).Delete()
		utils.WriteErrLogT(ctx, e, consts.DeleteF)
	})
	return
}

func (s *sLiveHistory) AddHistory(liveId int) {
	if liveId == 0 {
		return
	}
	global := utils.GetGlobalDefault()
	m, ok := global.StartTimeMap[liveId]
	if !ok || m == nil {
		return
	}
	endTime := gtime.Now()
	dao.LiveHistory.Ctx(gctx.New()).Insert(do.LiveHistory{
		LiveId:    liveId,
		StartTime: m,
		EndTime:   endTime,
		Duration:  fmt.Sprintf("%.2f", endTime.Sub(m).Hours()),
	})
}

func validTime(s, e *gtime.Time) error {
	if s == nil || e == nil {
		return gerror.Newf(`开播时间或下播时间不能为空`)
	}
	if s.After(e) {
		return gerror.Newf(`开播时间不能早于下播时间`)
	}
	return nil
}
