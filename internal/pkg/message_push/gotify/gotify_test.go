package gotify_test

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
)

func TestPush(t *testing.T) {
	url := "http://192.168.3.16:36073/message?token=AqgRo24hjfzUqtp"
	c := g.Client()
	data := g.Map{
		"title":   "开播通知",
		"message": "你关注的主播已开播",
	}
	c.Post(context.Background(), url, data)
}
