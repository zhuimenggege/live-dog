package media_parser_test

import (
	"testing"

	"github.com/gogf/gf/os/gctx"
	"github.com/shichen437/live-dog/internal/pkg/media_parser"
	_ "github.com/shichen437/live-dog/internal/pkg/media_parser/douyin"
)

func TestDownload(t *testing.T) {
	loader, err := media_parser.NewParser("https://www.iesdouyin.com/share/video/7479687670092942627")
	if err != nil {
		t.Error(err)
	}
	loader.ParseURL(gctx.New())
}
