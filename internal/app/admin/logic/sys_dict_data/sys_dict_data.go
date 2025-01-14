package sys_dict_data

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/admin"
	"github.com/shichen437/live-dog/internal/app/admin/consts"
	"github.com/shichen437/live-dog/internal/app/admin/dao"
	"github.com/shichen437/live-dog/internal/app/admin/model/do"
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"
	"github.com/shichen437/live-dog/internal/app/admin/service"
	cConsts "github.com/shichen437/live-dog/internal/app/common/consts"
	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	service.RegisterSysDictData(New())
}

func New() *sSysDictData {
	return &sSysDictData{}
}

type sSysDictData struct {
}

func (s *sSysDictData) GetDictDataByType(ctx context.Context, dictType string) (dictdata []*entity.SysDictData, err error) {
	if utils.IsInternalDictData(dictType) {
		data := utils.GetDictDataByType(dictType)
		gconv.Structs(data, &dictdata)
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		//字典数据表
		err = dao.SysDictData.Ctx(ctx).Where(dao.SysDictData.Columns().DictType, dictType).Scan(&dictdata)
		utils.WriteErrLogT(ctx, err, cConsts.GetF)
	})
	return
}

func (s *sSysDictData) GetDictDataList(ctx context.Context, req *v1.GetDictDataListReq) (res *v1.GetDictDataListRes, err error) {
	res = &v1.GetDictDataListRes{}
	m := dao.SysDictData.Ctx(ctx)
	if req.DictLabel != "" {
		m = m.WhereLike(dao.SysDictData.Columns().DictLabel, "%"+req.DictLabel+"%")
	}
	if req.DictType != "" {
		m = m.Where(dao.SysDictData.Columns().DictType, req.DictType)
	}
	if req.DictValue != "" {
		m = m.WhereLike(dao.SysDictData.Columns().DictValue, "%"+req.DictValue+"%")
	}
	if req.Status != "" {
		m = m.Where(dao.SysDictData.Columns().Status, req.Status)
	}
	err = g.Try(ctx, func(ctx context.Context) {
		var result []*entity.SysDictData
		res.Total, err = m.Count()
		err = m.Page(req.PageNum, req.PageSize).Scan(&result)
		utils.WriteErrLogT(ctx, err, cConsts.GetF)
		res.Data = result
	})
	return
}

func (s *sSysDictData) GetDictData(ctx context.Context, req *v1.GetDictDataDetailReq) (res *v1.GetDictDataDetailRes, err error) {
	if req.DictCode == 0 {
		err = utils.TError(ctx, cConsts.DataCodeEmpty)
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysDictData.Ctx(ctx).WherePri(req.DictCode).Scan(&res)
		utils.WriteErrLogT(ctx, err, cConsts.GetF)
	})
	return
}

func (s *sSysDictData) Add(ctx context.Context, req *v1.PostDictDataReq) (res *v1.PostDictDataRes, err error) {
	adminName := gconv.String(ctx.Value(cConsts.CtxAdminName))
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			//添加角色信息
			_, err := dao.SysDictData.Ctx(ctx).TX(tx).Insert(do.SysDictData{
				DictSort:   req.DictSort,
				DictLabel:  req.DictLabel,
				DictType:   req.DictType,
				DictValue:  req.DictValue,
				CssClass:   req.CssClass,
				ListClass:  req.ListClass,
				IsDefault:  req.IsDefault,
				Status:     req.Status,
				Remark:     req.Remark,
				CreateTime: gtime.Now(),
				CreateBy:   adminName,
				UpdateTime: gtime.Now(),
				UpdateBy:   adminName,
			})
			utils.WriteErrLogT(ctx, err, cConsts.AddF)
		})
		return err
	})
	return
}

func (s *sSysDictData) Update(ctx context.Context, req *v1.PutDictDataReq) (res *v1.PutDictDataRes, err error) {
	adminName := gconv.String(ctx.Value(cConsts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.SysDictData.Ctx(ctx).WherePri(&req.DictCode).Update(do.SysDictData{
			DictSort:   req.DictSort,
			DictLabel:  req.DictLabel,
			DictType:   req.DictType,
			DictValue:  req.DictValue,
			CssClass:   req.CssClass,
			ListClass:  req.ListClass,
			IsDefault:  req.IsDefault,
			Status:     req.Status,
			Remark:     req.Remark,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utils.WriteErrLogT(ctx, e, cConsts.UpdateF)
	})
	return
}

func (s *sSysDictData) Delete(ctx context.Context, req *v1.DeleteDictDataReq) (res *v1.DeleteDictDataRes, err error) {
	adminName := gconv.String(ctx.Value(cConsts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		ids := utils.ParamStrToSlice(req.DictCode, ",")
		_, e := dao.SysDictData.Ctx(ctx).WhereIn(dao.SysDictData.Columns().DictCode, ids).Update(do.SysDictData{
			Status:     consts.SysDictTypeStatusNo,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utils.WriteErrLogT(ctx, e, cConsts.DeleteF)
	})

	return
}
