package sys_config

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/admin"
	"github.com/shichen437/live-dog/internal/app/admin/dao"
	"github.com/shichen437/live-dog/internal/app/admin/model/do"
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"
	"github.com/shichen437/live-dog/internal/app/admin/service"
	commonConst "github.com/shichen437/live-dog/internal/app/common/consts"
	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	service.RegisterSysConfig(New())
}

func New() *sSysConfig {
	return &sSysConfig{}
}

type sSysConfig struct {
}

// 获取参数配置表列表
func (s *sSysConfig) GetSysConfigList(ctx context.Context, req *v1.GetSysConfigListReq) (res *v1.GetSysConfigListRes, err error) {
	res = &v1.GetSysConfigListRes{}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = commonConst.PageSize
	}
	err = g.Try(ctx, func(ctx context.Context) {
		var list []*entity.SysConfig
		m := dao.SysConfig.Ctx(ctx)
		if req.ConfigName != "" {
			m = m.WhereLike(dao.SysConfig.Columns().ConfigName, "%"+req.ConfigName+"%")
		}
		if req.ConfigKey != "" {
			m = m.Where(dao.SysConfig.Columns().ConfigKey, req.ConfigKey)
		}
		if req.ConfigType != "" {
			m = m.Where(dao.SysConfig.Columns().ConfigType, req.ConfigType)
		}
		if len(req.Params) > 0 {
			m = m.WhereBetween(dao.SysConfig.Columns().CreateTime, req.Params["beginTime"], req.Params["endTime"])
		}
		res.Total, err = m.Count()
		err = m.Page(req.PageNum, req.PageSize).Scan(&list)

		utils.WriteErrLogT(ctx, err, utils.T(ctx, commonConst.GetF))
		res.Rows = list
	})

	return
}

// 添加参数配置表
func (s *sSysConfig) Add(ctx context.Context, req *v1.PostSysConfigReq) (res *v1.PostSysConfigRes, err error) {
	adminName := gconv.String(ctx.Value(commonConst.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.SysConfig.Ctx(ctx).Data(do.SysConfig{
			ConfigName:  req.ConfigName,
			ConfigKey:   req.ConfigKey,
			ConfigValue: req.ConfigValue,
			ConfigType:  req.ConfigType,
			Remark:      req.Remark,
			UpdateTime:  gtime.Now(),
			UpdateBy:    adminName,
			CreateTime:  gtime.Now(),
			CreateBy:    adminName,
		}).Insert()
		utils.WriteErrLogT(ctx, e, commonConst.AddF)
	})

	return
}

// 修改参数配置表
func (s *sSysConfig) Update(ctx context.Context, req *v1.PutSysConfigReq) (res *v1.PutSysConfigRes, err error) {
	adminName := gconv.String(ctx.Value(commonConst.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.SysConfig.Ctx(ctx).WherePri(&req.ConfigId).Update(do.SysConfig{
			ConfigName:  req.ConfigName,
			ConfigKey:   req.ConfigKey,
			ConfigValue: req.ConfigValue,
			ConfigType:  req.ConfigType,
			Remark:      req.Remark,
			UpdateTime:  gtime.Now(),
			UpdateBy:    adminName,
		})
		utils.WriteErrLogT(ctx, e, commonConst.UpdateF)
	})

	return
}

// 删除参数配置表
func (s *sSysConfig) Delete(ctx context.Context, req *v1.DeleteSysConfigReq) (res *v1.DeleteSysConfigRes, err error) {

	err = g.Try(ctx, func(ctx context.Context) {
		postIds := utils.ParamStrToSlice(req.ConfigId, ",")
		_, e := dao.SysConfig.Ctx(ctx).WhereIn(dao.SysConfig.Columns().ConfigId, postIds).Delete()
		utils.WriteErrLogT(ctx, e, commonConst.DeleteF)
	})
	return
}

// 获取参数配置表
func (s *sSysConfig) GetSysConfig(ctx context.Context, req *v1.GetSysConfigReq) (res *v1.GetSysConfigRes, err error) {
	res = &v1.GetSysConfigRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		var table *entity.SysConfig
		err = dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns().ConfigId, req.ConfigId).Scan(&table)
		res.SysConfig = table
		utils.WriteErrLogT(ctx, err, commonConst.GetF)
	})

	return
}
