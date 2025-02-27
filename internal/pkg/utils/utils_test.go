package utils_test

import (
	"testing"

	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/shichen437/live-dog/internal/pkg/utils"
)

func TestPush(t *testing.T) {
	ok := utils.IsTimeRange("21:00", "00:00")
	g.Log().Info(gctx.New(), ok)
}
