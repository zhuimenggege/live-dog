package sys_job

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
	service.RegisterSysJob(New())
}

func New() *sSysJob {
	return &sSysJob{}
}

type sSysJob struct{}

func (s *sSysJob) List(ctx context.Context, req *v1.GetJobListReq) (result *v1.GetJobListRes, err error) {
	result = &v1.GetJobListRes{}

	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}

	err = g.Try(ctx, func(ctx context.Context) {
		var list []*entity.SysJob
		m := dao.SysJob.Ctx(ctx)
		if req.JobName != "" {
			m = m.WhereLike(dao.SysJob.Columns().JobName, "%"+req.JobName+"%")
		}
		if req.JobGroup != "" {
			m = m.WhereLike(dao.SysJob.Columns().JobGroup, "%"+req.JobGroup+"%")
		}
		m = m.Where(dao.SysJob.Columns().Status, req.Status)
		m = m.OrderDesc(dao.SysJob.Columns().CreateTime)
		result.Total, err = m.Count()
		err = m.Page(req.PageNum, req.PageSize).Scan(&list)
		result.Rows = list
	})
	return
}

func (s *sSysJob) LogList(ctx context.Context, req *v1.GetJobLogListReq) (result *v1.GetJobLogListRes, err error) {
	result = &v1.GetJobLogListRes{}

	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}

	err = g.Try(ctx, func(ctx context.Context) {
		var list []*entity.SysJobLog
		m := dao.SysJobLog.Ctx(ctx)
		if req.JobName != "" {
			m = m.WhereLike(dao.SysJobLog.Columns().JobName, "%"+req.JobName+"%")
		}
		if req.JobGroup != "" {
			m = m.WhereLike(dao.SysJobLog.Columns().JobGroup, "%"+req.JobGroup+"%")
		}
		m = m.Where(dao.SysJobLog.Columns().Status, req.Status)
		m = m.OrderDesc(dao.SysJobLog.Columns().CreateTime)
		result.Total, err = m.Count()
		err = m.Page(req.PageNum, req.PageSize).Scan(&list)
		result.Rows = list
	})
	return
}
