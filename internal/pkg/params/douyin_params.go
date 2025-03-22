package params

import (
	"net/url"
	"os"
	"path/filepath"

	"github.com/dop251/goja"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/shichen437/live-dog/internal/pkg/utils"
)

var (
	randomCookieChars = "1234567890abcdef"
	ttwid_url         = "https://ttwid.bytedance.com/ttwid/union/register/"
	ttwid_data        = `{"region":"cn","aid":1768,"needFid":false,"service":"www.ixigua.com","migrate_info":{"ticket":"","source":"node"},"cbUrlProtocol":"https","union":true}`
)

func GetABogus(params, ua string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		g.Log().Info(gctx.New(), wd)
	}
	jsFilePath := filepath.Join("internal/pkg/params/js/douyin", "a_bogus.js")
	jsContent, err := os.ReadFile(jsFilePath)
	if err != nil {
		return "", gerror.New("read js file failed!")
	}

	vm := vmPool.Get().(*goja.Runtime)
	defer vmPool.Put(vm)

	// 执行 JavaScript 文件内容
	_, err = vm.RunString(string(jsContent))
	if err != nil {
		return "", gerror.New("exec js file failed!")
	}

	// 获取 generate_a_bogus 方法
	generateABogus, ok := goja.AssertFunction(vm.Get("sign_detail"))
	if !ok {
		return "", gerror.New("read js method failed!")
	}
	result, err := generateABogus(goja.Undefined(), vm.ToValue(params), vm.ToValue(ua))
	if err != nil {
		return "", gerror.New("get js param failed!")
	}
	return url.QueryEscape(result.String()), nil
}

func GetOdintt() string {
	return utils.GenRandomString(160, randomCookieChars)
}

func GetMsToken() string {
	return utils.GenRandomString(107, randomCookieChars)
}

func GetTtwid() string {
	c := g.Client()
	resp, err := c.Post(gctx.New(), ttwid_url, ttwid_data)
	if err != nil {
		return ""
	}
	return resp.GetCookie("ttwid")
}
