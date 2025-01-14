// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "github.com/shichen437/live-dog/api/v1/live"
	"github.com/shichen437/live-dog/internal/app/live/model/entity"
)

type (
	ILiveCookie interface {
		List(ctx context.Context, req *v1.GetCookieListReq) (res *v1.GetCookieListRes, err error)
		Add(ctx context.Context, req *v1.PostLiveCookieReq) (res *v1.PostLiveCookieRes, err error)
		Get(ctx context.Context, req *v1.GetLiveCookieReq) (res *v1.GetLiveCookieRes, err error)
		Update(ctx context.Context, req *v1.PutLiveCookieReq) (res *v1.PutLiveCookieRes, err error)
		Delete(ctx context.Context, req *v1.DeleteLiveCookieReq) (res *v1.DeleteLiveCookieRes, err error)
		GetAllCookie4Init(ctx context.Context) (res []*entity.LiveCookie, err error)
	}
)

var (
	localLiveCookie ILiveCookie
)

func LiveCookie() ILiveCookie {
	if localLiveCookie == nil {
		panic("implement not found for interface ILiveCookie, forgot register?")
	}
	return localLiveCookie
}

func RegisterLiveCookie(i ILiveCookie) {
	localLiveCookie = i
}
