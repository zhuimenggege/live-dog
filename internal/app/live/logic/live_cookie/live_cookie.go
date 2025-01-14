package live_cookie

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/live"
	"github.com/shichen437/live-dog/internal/app/common/consts"
	"github.com/shichen437/live-dog/internal/app/live/dao"
	"github.com/shichen437/live-dog/internal/app/live/model/do"
	"github.com/shichen437/live-dog/internal/app/live/model/entity"
	"github.com/shichen437/live-dog/internal/app/live/service"
	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func init() {
	service.RegisterLiveCookie(New())
}

func New() *sLiveCookie {
	return &sLiveCookie{}
}

type sLiveCookie struct {
}

func (s *sLiveCookie) List(ctx context.Context, req *v1.GetCookieListReq) (res *v1.GetCookieListRes, err error) {
	res = &v1.GetCookieListRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.LiveCookie.Ctx(ctx)
		if req.Platform != "" {
			m = m.Where(dao.LiveCookie.Columns().Platform, req.Platform)
		}
		m = m.OrderDesc(dao.LiveCookie.Columns().CreateTime)
		res.Total, err = m.Count()
		err = m.Page(req.PageNum, req.PageSize).Scan(&res.Rows)
		utils.WriteErrLogT(ctx, err, consts.ListF)
	})
	return
}

func (s *sLiveCookie) Add(ctx context.Context, req *v1.PostLiveCookieReq) (res *v1.PostLiveCookieRes, err error) {
	if req.Cookie == "" {
		err = gerror.Newf(`主播名称不能为空`)
		return
	}
	if req.Platform == "" {
		err = gerror.Newf(`节目名称不能为空`)
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		//添加角色信息
		_, err := dao.LiveCookie.Ctx(ctx).Insert(do.LiveCookie{
			Platform:   req.Platform,
			Cookie:     req.Cookie,
			Remark:     req.Remark,
			CreateTime: gtime.Now(),
		})
		utils.WriteErrLogT(ctx, err, consts.AddF)
	})
	utils.GetGlobalDefault().CookieMap[req.Platform] = req.Cookie
	return
}

func (s *sLiveCookie) Get(ctx context.Context, req *v1.GetLiveCookieReq) (res *v1.GetLiveCookieRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.LiveCookie.Ctx(ctx).Where(dao.LiveCookie.Columns().Id, req.Id).Scan(&res)
		utils.WriteErrLogT(ctx, err, consts.GetF)
	})
	return
}

func (s *sLiveCookie) Update(ctx context.Context, req *v1.PutLiveCookieReq) (res *v1.PutLiveCookieRes, err error) {
	if req.Id == 0 {
		return nil, utils.TError(ctx, consts.IDEmpty)
	}
	err = g.Try(ctx, func(ctx context.Context) {
		//更新用户信息
		_, e := dao.LiveCookie.Ctx(ctx).WherePri(req.Id).Update(do.LiveCookie{
			Cookie:     req.Cookie,
			Remark:     req.Remark,
			ActionTime: gtime.Now(),
		})
		utils.WriteErrLogT(ctx, e, consts.AddF)
	})
	utils.GetGlobalDefault().CookieMap[req.Platform] = req.Cookie
	return
}

func (s *sLiveCookie) Delete(ctx context.Context, req *v1.DeleteLiveCookieReq) (res *v1.DeleteLiveCookieRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		ids := utils.ParamStrToSlice(req.Id, ",")
		//删除用户信息
		_, e := dao.LiveCookie.Ctx(ctx).WhereIn(dao.LiveCookie.Columns().Id, ids).Delete()
		utils.WriteErrLogT(ctx, e, consts.DeleteF)
	})
	return
}

func (s *sLiveCookie) GetAllCookie4Init(ctx context.Context) (res []*entity.LiveCookie, err error) {
	err = dao.LiveCookie.Ctx(ctx).WhereNot(dao.LiveCookie.Columns().Platform, "").Scan(&res)
	return
}
