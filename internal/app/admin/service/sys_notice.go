// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "github.com/shichen437/live-dog/api/v1/admin"
)

type (
	ISysNotice interface {
		GetSysNoticeList(ctx context.Context, req *v1.GetSysNoticeListReq) (res *v1.GetSysNoticeListRes, err error)
		GetNoticeData(ctx context.Context, req *v1.GetSysNoticeDetailReq) (res *v1.GetSysNoticeDetailRes, err error)
		Add(ctx context.Context, req *v1.PostSysNoticeReq) (res *v1.PostSysNoticeRes, err error)
		Update(ctx context.Context, req *v1.PutSysNoticeReq) (res *v1.PutSysNoticeRes, err error)
		Delete(ctx context.Context, req *v1.DeleteSysNoticeReq) (res *v1.DeleteSysNoticeRes, err error)
	}
)

var (
	localSysNotice ISysNotice
)

func SysNotice() ISysNotice {
	if localSysNotice == nil {
		panic("implement not found for interface ISysNotice, forgot register?")
	}
	return localSysNotice
}

func RegisterSysNotice(i ISysNotice) {
	localSysNotice = i
}
