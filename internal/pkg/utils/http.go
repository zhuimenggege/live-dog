package utils

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

// GetClientIp 获取客户端IP
func GetClientIp(ctx context.Context) string {
	return g.RequestFromCtx(ctx).GetClientIp()
}

// GetUserAgent 获取user-agent
func GetUserAgent(ctx context.Context) string {
	return ghttp.RequestFromCtx(ctx).Header.Get("User-Agent")
}

func Text(r *http.Response) (string, error) {
	if r.Body == nil {
		return "", io.EOF
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			g.Log().Error(gctx.GetInitCtx(), "Error closing response body:", err)
		}
	}()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func GetCookieList(platform string) []*http.Cookie {
	cookiesList := make([]*http.Cookie, 0)
	cookie := GetGlobal(gctx.GetInitCtx()).CookieMap[platform]
	if cookie == "" {
		return cookiesList
	}
	for _, cStr := range strings.Split(cookie, ";") {
		cArr := strings.SplitN(cStr, "=", 2)
		if len(cArr) != 2 {
			continue
		}
		cookiesList = append(cookiesList, &http.Cookie{
			Name:  strings.TrimSpace(cArr[0]),
			Value: strings.TrimSpace(cArr[1]),
		})
	}
	return cookiesList
}
