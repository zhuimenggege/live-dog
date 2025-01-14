package live_manage

import (
	"context"
	"net/url"
	"sync"

	v1 "github.com/shichen437/live-dog/api/v1/live"
	"github.com/shichen437/live-dog/internal/app/common/consts"
	lConst "github.com/shichen437/live-dog/internal/app/live/consts"
	"github.com/shichen437/live-dog/internal/app/live/dao"
	"github.com/shichen437/live-dog/internal/app/live/model/do"
	"github.com/shichen437/live-dog/internal/app/live/model/entity"
	"github.com/shichen437/live-dog/internal/app/live/service"

	"github.com/shichen437/live-dog/internal/pkg/crons"
	"github.com/shichen437/live-dog/internal/pkg/listeners"
	"github.com/shichen437/live-dog/internal/pkg/lives"
	"github.com/shichen437/live-dog/internal/pkg/recorders"
	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	service.RegisterLiveManage(New())
}

func New() *sLiveManage {
	return &sLiveManage{}
}

type sLiveManage struct {
	lock sync.RWMutex
}

func (s *sLiveManage) AddLiveManage(ctx context.Context, req *v1.PostLiveManageReq) (res *v1.PostLiveManageRes, err error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	live, err := dataValid(ctx, req)
	if err != nil {
		return
	}
	info, err := live.GetInfo()
	if err != nil || info.Anchor == "" {
		return
	}
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	result, err := dao.LiveManage.Ctx(ctx).Insert(do.LiveManage{
		RoomUrl:      req.RoomUrl,
		Format:       req.Format,
		Interval:     30,
		EnableNotice: req.EnableNotice,
		MonitorType:  req.MonitorType,
		MonitorStart: req.MonitorStart,
		MonitorStop:  req.MonitorStop,
		CreateTime:   gtime.Now(),
		ActionTime:   gtime.Now(),
		CreateBy:     adminName,
	})
	utils.WriteErrLogT(ctx, err, consts.AddF)
	info.Status = req.MonitorType
	liveId, _ := result.LastInsertId()
	_, err = dao.RoomInfo.Ctx(ctx).Insert(do.RoomInfo{
		Anchor:     info.Anchor,
		RoomName:   info.RoomName,
		Platform:   info.Platform,
		Status:     info.Status,
		LiveId:     liveId,
		CreateTime: gtime.Now(),
		ActionTime: gtime.Now(),
		CreateBy:   adminName,
	})
	utils.WriteErrLogT(ctx, err, consts.AddF)
	go listenerForNew(int(liveId), info.Status)
	return
}

func (s *sLiveManage) UpdateLiveManage(ctx context.Context, req *v1.PutLiveManageReq) (res *v1.PutLiveManageRes, err error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if req.Id == 0 {
		err = utils.TError(ctx, consts.IDEmpty)
		return
	}
	if req.MonitorType == 1 {
		if ok, err := validMonitor(ctx, req.MonitorStart, req.MonitorStop); !ok {
			return nil, err
		}
	}
	var om *entity.LiveManage
	err = g.Try(ctx, func(ctx context.Context) {
		dao.LiveManage.Ctx(ctx).Where(dao.LiveManage.Columns().Id, req.Id).Scan(&om)
	})
	if err != nil {
		err = utils.TError(ctx, consts.DataNotFound)
		return
	}
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		//更新用户信息
		_, e := dao.LiveManage.Ctx(ctx).WherePri(req.Id).Update(do.LiveManage{
			Interval:     req.Interval,
			Format:       req.Format,
			EnableNotice: req.EnableNotice,
			MonitorType:  req.MonitorType,
			MonitorStart: req.MonitorStart,
			MonitorStop:  req.MonitorStop,
			Remark:       req.Remark,
			ActionTime:   gtime.Now(),
			ActionBy:     adminName,
		})
		utils.WriteErrLogT(ctx, e, consts.UpdateF)
	})
	if req.MonitorType != om.MonitorType {
		err = g.Try(ctx, func(ctx context.Context) {
			//更新用户信息
			_, e := dao.RoomInfo.Ctx(ctx).Where(dao.RoomInfo.Columns().LiveId, req.Id).Update(do.RoomInfo{
				Status:     req.MonitorType,
				ActionTime: gtime.Now(),
			})
			utils.WriteErrLogT(ctx, e, consts.UpdateF)
		})
	}
	go listenerForUpdate(req, om)
	return
}

func (s *sLiveManage) GetRoomInfoList(ctx context.Context, req *v1.GetRoomInfoListReq) (res *v1.GetRoomInfoListRes, err error) {
	res = &v1.GetRoomInfoListRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		var list []*entity.RoomInfo
		m := dao.RoomInfo.Ctx(ctx)
		if req.Anchor != "" {
			m = m.WhereLike(dao.RoomInfo.Columns().Anchor, "%"+req.Anchor+"%")
		}
		if req.RoomName != "" {
			m = m.WhereLike(dao.RoomInfo.Columns().RoomName, "%"+req.RoomName+"%")
		}
		if req.Platform != "" {
			m = m.WhereLike(dao.RoomInfo.Columns().Platform, "%"+req.Platform+"%")
		}
		m = m.OrderDesc(dao.RoomInfo.Columns().CreateTime)
		res.Total, err = m.Count()
		utils.WriteErrLogT(ctx, err, consts.ListF)
		if res.Total > 0 {
			err = m.Page(req.PageNum, req.PageSize).Scan(&list)
			utils.WriteErrLogT(ctx, err, consts.ListF)
			gconv.Structs(list, &res.Rows)
			global := utils.GetGlobalDefault()
			for _, v := range res.Rows {
				if v.Status != 0 {
					v.Recording = global.RecorderManager.(recorders.Manager).HasRecorder(gctx.GetInitCtx(), v.LiveId)
				}
			}
		}
	})
	return
}

func (s *sLiveManage) GetLiveManage(ctx context.Context, req *v1.GetLiveManageReq) (res *v1.GetLiveManageRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.LiveManage.Ctx(ctx).Where(dao.LiveManage.Columns().Id, req.Id).Scan(&res)
		utils.WriteErrLogT(ctx, err, consts.GetF)
	})
	return
}

func (s *sLiveManage) DeleteLiveManage(ctx context.Context, req *v1.DeleteLiveManageReq) (res *v1.DeleteLiveManageRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.LiveManage.Ctx(ctx).Where(dao.LiveManage.Columns().Id, req.Id).Delete()
	})
	utils.WriteErrLogT(ctx, err, consts.DeleteF)
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.RoomInfo.Ctx(ctx).Where(dao.RoomInfo.Columns().LiveId, req.Id).Delete()
	})
	utils.WriteErrLogT(ctx, err, consts.DeleteF)
	go listenerForDelete(req.Id)
	return
}

func (s *sLiveManage) GetLiveModels4Init(ctx context.Context) (liveModels []*lives.LiveModel, err error) {
	var mList []*entity.LiveManage
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.LiveManage.Ctx(ctx)
		m = m.WhereNot(dao.LiveManage.Columns().MonitorType, 0)
		err = m.Scan(&mList)
	})
	if len(mList) == 0 {
		return
	}
	liveModels = make([]*lives.LiveModel, 0, len(mList))
	var rList []*entity.RoomInfo
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.RoomInfo.Ctx(ctx)
		err = m.Scan(&rList)
	})
	if len(rList) > 0 {
		rMap := make(map[int]*entity.RoomInfo)
		for _, r := range rList {
			rMap[r.LiveId] = r
		}
		for _, m := range mList {
			r, ok := rMap[m.Id]
			if !ok {
				continue
			}
			lm, err := assembleLiveModel(m, r)
			if err != nil {
				continue
			}
			liveModels = append(liveModels, lm)
		}
	}
	return
}

func dataValid(ctx context.Context, req *v1.PostLiveManageReq) (lives.Live, error) {
	if req.RoomUrl == "" {
		return nil, utils.TError(ctx, lConst.RoomUrlEmpty)
	}
	sUrl, err := url.Parse(req.RoomUrl)
	if err != nil {
		return nil, utils.TError(ctx, lConst.RoomUrlParseF)
	}
	live, err := lives.New(sUrl, 0)
	if err != nil {
		return nil, utils.TError(ctx, lConst.RoomUrlNotAllowed)
	}
	if req.MonitorType == 1 {
		ok, err1 := validMonitor(ctx, req.MonitorStart, req.MonitorStop)
		if !ok {
			return nil, err1
		}
	}
	return live, nil
}

func validMonitor(ctx context.Context, start, stop string) (bool, error) {
	if start == "" || stop == "" {
		return false, utils.TError(ctx, lConst.MonitorTimeEmpty)
	}
	if start == stop {
		return false, utils.TError(ctx, lConst.MonitorTimeSame)
	}
	if utils.IsWithinCustomTimes(start, stop, 2) {
		return false, utils.TError(ctx, lConst.MonitorTimeShort)
	}
	return true, nil
}

func assembleLiveModel(m *entity.LiveManage, r *entity.RoomInfo) (model *lives.LiveModel, err error) {
	var lm lives.LiveManage
	var info lives.RoomInfo
	if err := gconv.Struct(m, &lm); err != nil {
		return nil, err
	}
	if err := gconv.Struct(r, &info); err != nil {
		return nil, err
	}
	return &lives.LiveModel{
		LiveManage: lm,
		RoomInfo:   info,
	}, nil
}

func listenerForNew(liveId int, status int) {
	if status == 0 {
		return
	}
	global := utils.GetGlobalDefault()
	// 添加监听器
	liveManage, roomInfo, err := getLiveManageAndInfo(liveId)
	if err != nil {
		g.Log().Error(gctx.GetInitCtx(), "listenerForNew getLiveManageAndInfo error:", err)
		return
	}
	lm, err := assembleLiveModel(liveManage, roomInfo)
	if err != nil {
		g.Log().Error(gctx.GetInitCtx(), "listenerForNew assembleLiveModel error:", err)
		return
	}
	u, err := lm.ParseUrl()
	if err != nil {
		g.Log().Error(gctx.GetInitCtx(), "listenerForNew ParseUrl error:", err)
		return
	}
	global.ModelsMap[liveId] = lm
	l, _ := lives.New(u, liveId)
	global.Lives[liveId] = l
	if status == 1 {
		crons.AddCron(liveId)
		return
	}
	global.ListenerManager.(listeners.Manager).AddListener(gctx.GetInitCtx(), l)
}

func listenerForUpdate(req *v1.PutLiveManageReq, om *entity.LiveManage) {
	r := judgeRestartListener(req, om)
	if r {
		listenerForDelete(req.Id)
		listenerForNew(req.Id, req.MonitorType)
		return
	}
	if req.EnableNotice != om.EnableNotice {
		global := utils.GetGlobalDefault()
		_, ok := global.ModelsMap[req.Id]
		if ok {
			global.ModelsMap[req.Id].LiveManage.EnableNotice = req.EnableNotice
		}
	}
}

func judgeRestartListener(req *v1.PutLiveManageReq, om *entity.LiveManage) bool {
	base := req.Interval != int(om.Interval) || req.Format != om.Format || req.MonitorType != om.MonitorType
	if base {
		return base
	}
	if req.MonitorType == 1 {
		return req.MonitorStart != om.MonitorStart || req.MonitorStop != om.MonitorStop
	}
	return false
}

func listenerForDelete(liveId int) {
	global := utils.GetGlobalDefault()
	global.ListenerManager.(listeners.Manager).RemoveListener(gctx.GetInitCtx(), liveId)
	crons.RemoveCron(liveId)
	delete(global.ModelsMap, liveId)
	delete(global.Lives, liveId)
}

func getLiveManageAndInfo(liveId int) (liveManage *entity.LiveManage, roomInfo *entity.RoomInfo, err error) {
	err = dao.LiveManage.Ctx(gctx.GetInitCtx()).Where(dao.LiveManage.Columns().Id, liveId).Scan(&liveManage)
	if err != nil {
		return
	}
	model, err := dao.RoomInfo.Ctx(gctx.GetInitCtx()).Where(dao.LiveManage.Columns().Id, liveId).One()
	if err != nil {
		return
	}
	model.Struct(&roomInfo)
	return
}
