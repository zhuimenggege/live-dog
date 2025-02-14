// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/monitor"
	"github.com/shichen437/live-dog/internal/app/monitor/model/entity"
)

type (
	ISysJob interface {
		List(ctx context.Context, req *v1.GetJobListReq) (result *v1.GetJobListRes, err error)
		LogList(ctx context.Context, req *v1.GetJobLogListReq) (result *v1.GetJobLogListRes, err error)
		Get(ctx context.Context, req *v1.GetJobDetailReq) (result *v1.GetJobDetailRes, err error)
		Add(ctx context.Context, req *v1.PostJobReq) (result *v1.PostJobRes, err error)
		Update(ctx context.Context, req *v1.PutJobReq) (result *v1.PutJobRes, err error)
		Delete(ctx context.Context, req *v1.DeleteJobReq) (result *v1.DeleteJobRes, err error)
		DeleteLog(ctx context.Context, req *v1.DeleteJobLogReq) (result *v1.DeleteJobLogRes, err error)
		RunJob(ctx context.Context, req *v1.PutJobRunReq) (result *v1.PutJobRunRes, err error)
		ChangeStatus(ctx context.Context, req *v1.PutJobStatusReq) (result *v1.PutJobStatusRes, err error)
		ListAll4Init(ctx context.Context) (result []*entity.SysJob, err error)
		GetJobDetail(ctx context.Context, jobId int64) (result *entity.SysJob, err error)
		AddLog(ctx context.Context, jobLog *entity.SysJobLog) (err error)
	}
)

var (
	localSysJob ISysJob
)

func SysJob() ISysJob {
	if localSysJob == nil {
		panic("implement not found for interface ISysJob, forgot register?")
	}
	return localSysJob
}

func RegisterSysJob(i ISysJob) {
	localSysJob = i
}
