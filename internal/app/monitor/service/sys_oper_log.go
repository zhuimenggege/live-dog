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
	ISysOperLog interface {
		List(ctx context.Context, req *v1.GetOperLogListReq) (result *v1.GetOperLogListRes, err error)
	}
)

var (
	localSysOperLog ISysOperLog
)

func SysOperLog() ISysOperLog {
	if localSysOperLog == nil {
		panic("implement not found for interface ISysOperLog, forgot register?")
	}
	return localSysOperLog
}

func RegisterSysOperLog(i ISysOperLog) {
	localSysOperLog = i
}
