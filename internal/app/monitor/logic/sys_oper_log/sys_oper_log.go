package sys_oper_log

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/monitor"
	"github.com/shichen437/live-dog/internal/app/common/consts"
	"github.com/shichen437/live-dog/internal/app/monitor/dao"
	"github.com/shichen437/live-dog/internal/app/monitor/model/entity"
	"github.com/shichen437/live-dog/internal/app/monitor/service"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterSysOperLog(New())
}

func New() *sSysOperLog {
	return &sSysOperLog{}
}

type sSysOperLog struct{}

func (s *sSysOperLog) List(ctx context.Context, req *v1.GetOperLogListReq) (result *v1.GetOperLogListRes, err error) {
	result = &v1.GetOperLogListRes{}

	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}

	err = g.Try(ctx, func(ctx context.Context) {
		var list []*entity.SysOperLog
		m := dao.SysOperLog.Ctx(ctx)
		if req.Title != "" {
			m = m.WhereLike(dao.SysOperLog.Columns().Title, "%"+req.Title+"%")
		}
		if req.OperName != "" {
			m = m.WhereLike(dao.SysOperLog.Columns().OperName, "%"+req.OperName+"%")
		}
		m = m.Where(dao.SysOperLog.Columns().BusinessType, req.BusinessType)
		m = m.Where(dao.SysOperLog.Columns().Status, req.Status)
		m = m.OrderDesc(dao.SysOperLog.Columns().OperTime)
		result.Total, err = m.Count()
		err = m.Page(req.PageNum, req.PageSize).Scan(&list)
		result.Rows = list
	})
	return
}
