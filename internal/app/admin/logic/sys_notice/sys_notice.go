package sys_notice

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/admin"
	"github.com/shichen437/live-dog/internal/app/admin/consts"
	"github.com/shichen437/live-dog/internal/app/admin/dao"
	"github.com/shichen437/live-dog/internal/app/admin/model/do"
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"
	"github.com/shichen437/live-dog/internal/app/admin/service"
	commonConst "github.com/shichen437/live-dog/internal/app/common/consts"
	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	service.RegisterSysNotice(New())
}

func New() *sSysNotice {
	return &sSysNotice{}
}

type sSysNotice struct {
}

func (s *sSysNotice) GetSysNoticeList(ctx context.Context, req *v1.GetSysNoticeListReq) (res *v1.GetSysNoticeListRes, err error) {
	res = &v1.GetSysNoticeListRes{}
	m := dao.SysNotice.Ctx(ctx)
	if req.NoticeTitle != "" {
		m = m.WhereLike(dao.SysNotice.Columns().NoticeTitle, "%"+req.NoticeTitle+"%")
	}
	if req.NoticeType != "" {
		m = m.Where(dao.SysNotice.Columns().NoticeType, req.NoticeType)
	}
	if req.Status != "" {
		m = m.Where(dao.SysNotice.Columns().Status, req.Status)
	}
	err = g.Try(ctx, func(ctx context.Context) {
		var result []*entity.SysNotice

		res.Total, err = m.Count()
		err = m.Page(req.PageNum, req.PageSize).Scan(&result)
		utils.WriteErrLogT(ctx, err, commonConst.ListF)
		res.Data = result
	})
	return
}

func (s *sSysNotice) GetNoticeData(ctx context.Context, req *v1.GetSysNoticeDetailReq) (res *v1.GetSysNoticeDetailRes, err error) {
	if req.NoticeId == 0 {
		err = utils.TError(ctx, commonConst.IDEmpty)
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysNotice.Ctx(ctx).WherePri(req.NoticeId).Scan(&res)
		utils.WriteErrLogT(ctx, err, commonConst.GetF)
	})
	return
}

func (s *sSysNotice) Add(ctx context.Context, req *v1.PostSysNoticeReq) (res *v1.PostSysNoticeRes, err error) {
	adminName := gconv.String(ctx.Value(commonConst.CtxAdminName))
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			//添加角色信息
			_, err := dao.SysNotice.Ctx(ctx).TX(tx).Insert(do.SysNotice{
				NoticeTitle:   req.NoticeTitle,
				NoticeType:    req.NoticeType,
				NoticeContent: req.NoticeContent,
				Status:        req.Status,
				Remark:        req.Remark,
				CreateTime:    gtime.Now(),
				CreateBy:      adminName,
				UpdateTime:    gtime.Now(),
				UpdateBy:      adminName,
			})
			utils.WriteErrLogT(ctx, err, commonConst.AddF)
		})
		return err
	})
	return
}

func (s *sSysNotice) Update(ctx context.Context, req *v1.PutSysNoticeReq) (res *v1.PutSysNoticeRes, err error) {
	adminName := gconv.String(ctx.Value(commonConst.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.SysNotice.Ctx(ctx).WherePri(&req.NoticeId).Update(do.SysNotice{
			NoticeId:      req.NoticeId,
			NoticeTitle:   req.NoticeTitle,
			NoticeType:    req.NoticeType,
			NoticeContent: req.NoticeContent,
			Status:        req.Status,
			Remark:        req.Remark,
			UpdateTime:    gtime.Now(),
			UpdateBy:      adminName,
		})
		utils.WriteErrLogT(ctx, e, commonConst.UpdateF)
	})
	return
}

func (s *sSysNotice) Delete(ctx context.Context, req *v1.DeleteSysNoticeReq) (res *v1.DeleteSysNoticeRes, err error) {
	adminName := gconv.String(ctx.Value(commonConst.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		ids := utils.ParamStrToSlice(req.NoticeId, ",")
		_, e := dao.SysNotice.Ctx(ctx).WhereIn(dao.SysNotice.Columns().NoticeId, ids).Update(do.SysNotice{
			Status:     consts.SysDictTypeStatusNo,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utils.WriteErrLogT(ctx, e, commonConst.DeleteF)
	})
	return
}
