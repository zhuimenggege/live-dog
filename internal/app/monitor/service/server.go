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
	IServerInfo interface {
		GetServerInfo(ctx context.Context, req *v1.GetServerInfoReq) (res *v1.GetServerInfoRes, err error)
	}
)

var (
	localServerInfo IServerInfo
)

func ServerInfo() IServerInfo {
	if localServerInfo == nil {
		panic("implement not found for interface IServerInfo, forgot register?")
	}
	return localServerInfo
}

func RegisterServerInfo(i IServerInfo) {
	localServerInfo = i
}
