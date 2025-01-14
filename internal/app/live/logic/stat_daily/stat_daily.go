package stat_daily

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/live"
	"github.com/shichen437/live-dog/internal/app/common/consts"
	"github.com/shichen437/live-dog/internal/app/live/dao"
	"github.com/shichen437/live-dog/internal/app/live/model"
	"github.com/shichen437/live-dog/internal/app/live/model/do"
	"github.com/shichen437/live-dog/internal/app/live/model/entity"
	"github.com/shichen437/live-dog/internal/app/live/service"
	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	service.RegisterStatDaily(New())
}

func New() *sStatDaily {
	return &sStatDaily{}
}

type sStatDaily struct {
}

func (s *sStatDaily) GetStatDailyById(ctx context.Context, id int64) (dailies *entity.StatDaily, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.StatDaily.Ctx(ctx).Where(dao.StatDaily.Columns().Id, id).Scan(&dailies)
		utils.WriteErrLogT(ctx, err, consts.GetF)
	})
	return
}

func GetStatDaily(ctx context.Context, queryParams *entity.StatDaily) (result *entity.StatDaily, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.StatDaily.Ctx(ctx)
		m = m.Where(dao.StatDaily.Columns().Anchor, queryParams.Anchor)
		m = m.Where(dao.StatDaily.Columns().DisplayName, queryParams.DisplayName)
		m = m.Where(dao.StatDaily.Columns().DisplayDate, queryParams.DisplayDate)
		m = m.Where(dao.StatDaily.Columns().DisplayType, queryParams.DisplayType)
		m = m.Where(dao.StatDaily.Columns().Remark, queryParams.Remark)
		m = m.WhereNot(dao.StatDaily.Columns().Action, 2)
		err = m.Scan(&result)
		utils.WriteErrLogT(ctx, err, consts.GetF)
	})
	return
}

func (s *sStatDaily) GetForUpdate(ctx context.Context, req *v1.GetStatDailyReq) (statDaily *v1.GetStatDailyRes, err error) {
	if req.Id == 0 {
		err = utils.TError(ctx, consts.IDEmpty)
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.StatDaily.Ctx(ctx).Where(dao.StatDaily.Columns().Id, req.Id).Scan(&statDaily)
		utils.WriteErrLogT(ctx, err, consts.GetF)
	})
	return
}

func (s *sStatDaily) GetStatDailyList(ctx context.Context, req *v1.GetStatDailyListReq) (dailyList *v1.GetStatDailyListRes, err error) {
	dailyList = &v1.GetStatDailyListRes{}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}

	err = g.Try(ctx, func(ctx context.Context) {
		var list []*entity.StatDaily
		m := dao.StatDaily.Ctx(ctx)
		if req.Anchor != "" {
			m = m.WhereLike(dao.StatDaily.Columns().Anchor, "%"+req.Anchor+"%")
		}
		if req.DisplayName != "" {
			m = m.WhereLike(dao.StatDaily.Columns().DisplayName, "%"+req.DisplayName+"%")
		}
		if req.DisplayDate != "" {
			m = m.Where(dao.StatDaily.Columns().DisplayDate, req.DisplayDate)
		}
		if req.DisplayType != 0 {
			m = m.Where(dao.StatDaily.Columns().DisplayType, req.DisplayType)
		}
		m = m.WhereNot(dao.StatDaily.Columns().Action, 2)
		m = m.OrderDesc(dao.StatDaily.Columns().CreateTime)
		dailyList.Total, err = m.Count()
		err = m.Page(req.PageNum, req.PageSize).Scan(&list)
		rows := make([]*model.StatDailyList, len(list))
		for k, value := range list {
			ul := &model.StatDailyList{}
			ul.StatDaily = value
			rows[k] = ul
		}
		dailyList.Rows = rows
	})
	utils.WriteErrLogT(ctx, err, consts.ListF)
	return
}

func (s *sStatDaily) Add(ctx context.Context, req *v1.PostStatDailyReq) (res *v1.PostStatDailyRes, err error) {
	if req.Anchor == "" {
		err = gerror.Newf(`主播名称不能为空`)
		return
	}
	if req.DisplayName == "" {
		err = gerror.Newf(`节目名称不能为空`)
		return
	}
	if req.DisplayType == 0 {
		err = gerror.Newf(`节目类型不能为空`)
		return
	}
	if req.DisplayDate == "" {
		err = gerror.Newf(`节目日期不能为空`)
		return
	}
	if req.Count == 0 {
		req.Count = 1
	}
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	var entity entity.StatDaily
	entity.Anchor = req.Anchor
	entity.DisplayName = req.DisplayName
	entity.DisplayType = req.DisplayType
	entity.DisplayDate = req.DisplayDate
	entity.Remark = req.Remark
	target, err := GetStatDaily(ctx, &entity)
	if target != nil && target.Id > 0 {
		target.Count += 1
		err = g.Try(ctx, func(ctx context.Context) {
			//更新用户信息
			_, e := dao.StatDaily.Ctx(ctx).WherePri(target.Id).Update(do.StatDaily{
				Count:      target.Count,
				Action:     1,
				UpdateTime: gtime.Now(),
				UpdateBy:   adminName,
			})
			utils.WriteErrLogT(ctx, e, consts.AddF)
		})

	} else {
		err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			err = g.Try(ctx, func(ctx context.Context) {
				//添加角色信息
				_, err := dao.StatDaily.Ctx(ctx).TX(tx).Insert(do.StatDaily{
					Anchor:      req.Anchor,
					DisplayName: req.DisplayName,
					DisplayType: req.DisplayType,
					DisplayDate: req.DisplayDate,
					Count:       req.Count,
					Remark:      req.Remark,
					CreateTime:  gtime.Now(),
					CreateBy:    adminName,
					UpdateTime:  gtime.Now(),
					UpdateBy:    adminName,
				})
				utils.WriteErrLogT(ctx, err, consts.AddF)
			})
			return err
		})
	}
	return
}

func (s *sStatDaily) Update(ctx context.Context, req *v1.PutStatDailyReq) (res *v1.PutStatDailyRes, err error) {
	if req.Id == 0 {
		err = utils.TError(ctx, consts.IDEmpty)
		return
	}
	if req.Anchor == "" {
		err = gerror.Newf(`主播名称不能为空`)
		return
	}
	if req.DisplayName == "" {
		err = gerror.Newf(`节目名称不能为空`)
		return
	}
	if req.DisplayType == 0 {
		err = gerror.Newf(`节目类型不能为空`)
		return
	}
	if req.DisplayDate == "" {
		err = gerror.Newf(`节目日期不能为空`)
		return
	}
	if req.Count == 0 {
		req.Count = 1
	}
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		//更新用户信息
		_, e := dao.StatDaily.Ctx(ctx).WherePri(req.Id).Update(do.StatDaily{
			Anchor:      req.Anchor,
			DisplayName: req.DisplayName,
			DisplayType: req.DisplayType,
			DisplayDate: req.DisplayDate,
			Count:       req.Count,
			Remark:      req.Remark,
			Action:      1,
			UpdateTime:  gtime.Now(),
			UpdateBy:    adminName,
		})
		utils.WriteErrLogT(ctx, e, consts.UpdateF)
	})
	return
}

func (s *sStatDaily) Delete(ctx context.Context, req *v1.DeleteStatDailyReq) (res *v1.DeleteStatDailyRes, err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		ids := utils.ParamStrToSlice(req.Id, ",")
		//删除用户信息
		_, e := dao.StatDaily.Ctx(ctx).WhereIn(dao.StatDaily.Columns().Id, ids).Update(do.StatDaily{
			Action:     2,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utils.WriteErrLogT(ctx, e, consts.DeleteF)
	})
	return
}
