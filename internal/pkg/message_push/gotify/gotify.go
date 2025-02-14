package gotify

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/shichen437/live-dog/internal/app/admin/model"
	mp "github.com/shichen437/live-dog/internal/pkg/message_push"
	"github.com/shichen437/live-dog/internal/pkg/utils"
)

const (
	channelType = "gotify"
)

func init() {
	mp.Register(channelType, &builder{})
}

type builder struct{}

func (b *builder) Build(channelType string, liveId int) (mp.MessagePush, error) {
	return &MessagePush{
		LiveId: liveId,
	}, nil
}

type MessagePush struct {
	LiveId int
}

func (p *MessagePush) Push(ctx context.Context, channel *model.PushChannel) (err error) {
	global := utils.GetGlobal(gctx.GetInitCtx())
	lm := global.ModelsMap[p.LiveId]
	if lm.LiveManage.EnableNotice != 1 {
		return
	}
	gotify(channel.Web.Url, "开播通知", "你关注的主播["+lm.RoomInfo.Anchor+"]开播了！")
	return nil
}

func (p *MessagePush) CustomPush(ctx context.Context, channel *model.PushChannel, model *mp.MessageModel) (err error){
	gotify(channel.Web.Url, model.Title, model.Content)
	return
}

func gotify(url string, title, message string) {
	c := g.Client()
	data := g.Map{
		"title":   title,
		"message": message,
	}
	c.Post(context.Background(), url, data)
}
