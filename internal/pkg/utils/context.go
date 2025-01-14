package utils

import (
	"context"

	"github.com/shichen437/live-dog/internal/pkg/lives"

	"github.com/gogf/gf/v2/os/gctx"
)

type key int

const (
	Key key = 102723
)

func GetGlobal(context context.Context) *lives.GLiveModel {
	if glm, ok := context.Value(Key).(*lives.GLiveModel); ok {
		return glm
	}
	return nil
}

func GetGlobalDefault() *lives.GLiveModel {
	return GetGlobal(gctx.GetInitCtx())
}
