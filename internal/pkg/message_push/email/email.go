package email

import (
	"context"

	"github.com/shichen437/live-dog/internal/app/admin/model"
	mp "github.com/shichen437/live-dog/internal/pkg/message_push"
	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/v2/os/gctx"
	"gopkg.in/mail.v2"
)

const (
	channelType = "email"
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
	var model mp.MessageModel
	model.Title = "开播通知"
	model.Content = "你关注的主播["+lm.RoomInfo.Anchor+"]开播了！"
	err = p.CustomPush(ctx, channel, &model)
	return err
}

func (p *MessagePush) CustomPush(ctx context.Context, channel *model.PushChannel, model *mp.MessageModel) (err error) {
	m := mail.NewMessage()
	m.SetHeader("From", channel.Email.From)
	m.SetHeader("To", channel.Email.To)
	m.SetHeader("Subject", model.Title)
	m.SetBody("text/html", model.Content)
	d := mail.NewDialer(channel.Email.Server, channel.Email.Port, channel.Email.From, channel.Email.AuthCode)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	err = d.DialAndSend(m)
	return err
}