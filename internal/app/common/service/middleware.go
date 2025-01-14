// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// Ctx injects custom business context variable into context of current request.
		Ctx(r *ghttp.Request)
		// Auth validates the request to allow only signed-in users visit.
		Auth(r *ghttp.Request)
		// CORS allows Cross-origin resource sharing.
		CORS(r *ghttp.Request)
		HandlerResponse(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
