// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "github.com/shichen437/live-dog/api/v1/monitor"
)

type (
	ISysJob interface {
		List(ctx context.Context, req *v1.GetJobListReq) (result *v1.GetJobListRes, err error)
		LogList(ctx context.Context, req *v1.GetJobLogListReq) (result *v1.GetJobLogListRes, err error)
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
