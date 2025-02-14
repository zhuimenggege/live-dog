package monitor

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/monitor"
	"github.com/shichen437/live-dog/internal/app/monitor/service"
)

type sysJobController struct {
}

var SysJob = sysJobController{}

func (o *sysJobController) GetJobList(ctx context.Context, req *v1.GetJobListReq) (res *v1.GetJobListRes, err error) {
	res, err = service.SysJob().List(ctx, req)
	return
}

func (o *sysJobController) GetJobLogList(ctx context.Context, req *v1.GetJobLogListReq) (res *v1.GetJobLogListRes, err error) {
	res, err = service.SysJob().LogList(ctx, req)
	return
}

func (o *sysJobController) GetJobDetail(ctx context.Context, req *v1.GetJobDetailReq) (res *v1.GetJobDetailRes, err error) {
	res, err = service.SysJob().Get(ctx, req)
	return
}

func (o *sysJobController) AddJob(ctx context.Context, req *v1.PostJobReq) (res *v1.PostJobRes, err error) {
	res, err = service.SysJob().Add(ctx, req)
	return
}

func (o *sysJobController) UpdateJob(ctx context.Context, req *v1.PutJobReq) (res *v1.PutJobRes, err error) {
	res, err = service.SysJob().Update(ctx, req)
	return
}

func (o *sysJobController) DeleteJob(ctx context.Context, req *v1.DeleteJobReq) (res *v1.DeleteJobRes, err error) {
	res, err = service.SysJob().Delete(ctx, req)
	return
}

func (o *sysJobController) RunJob(ctx context.Context, req *v1.PutJobRunReq) (res *v1.PutJobRunRes, err error) {
	res, err = service.SysJob().RunJob(ctx, req)
	return
}

func (o *sysJobController) ChangeJobStatus(ctx context.Context, req *v1.PutJobStatusReq) (res *v1.PutJobStatusRes, err error) {
	res, err = service.SysJob().ChangeStatus(ctx, req)
	return
}

func (o *sysJobController) DeleteJobLog(ctx context.Context, req *v1.DeleteJobLogReq) (res *v1.DeleteJobLogRes, err error) {
	res, err = service.SysJob().DeleteLog(ctx, req)
	return
}
