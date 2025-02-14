package message_push

import (
	"context"
	"sync"

	"github.com/shichen437/live-dog/internal/app/admin/model"
	"github.com/shichen437/live-dog/internal/app/admin/service"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	builders sync.Map
)

type MessageModel struct {
	Title   string
	Content string
}

type MessagePush interface {
	Push(ctx context.Context, channel *model.PushChannel) (err error)
	CustomPush(ctx context.Context, channel *model.PushChannel, model *MessageModel) (err error)
}

func Register(channelType string, b Builder) {
	builders.Store(channelType, b)
}

type Builder interface {
	Build(string, int) (MessagePush, error)
}

func LivePush(ctx context.Context, liveId int) {
	channels, err := service.PushChannel().ListAll(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	for _, v := range channels {
		ChannelPush(ctx, v, liveId)
	}
}

func ChannelPush(ctx context.Context, v *model.PushChannel, liveId int) error {
	b, err := getBuilder(v.Type)
	if err != nil {
		return err
	}
	builder, err := b.Build(v.Type, liveId)
	if err != nil {
		return gerror.New("不支持的渠道类型！")
	}
	err = builder.Push(ctx, v)
	if err != nil {
		return gerror.New("消息推送失败")
	}
	return nil
}

func getBuilder(channelType string) (Builder, error) {
	builder, ok := builders.Load(channelType)
	if !ok {
		return nil, gerror.New("不支持的渠道类型！")
	}
	return builder.(Builder), nil
}

func CustomPush(ctx context.Context, model *MessageModel) (err error) {
	channels, err := service.PushChannel().ListAll(ctx)
	if err != nil {
		return err
	}
	for _, channel := range channels {
		b, err := getBuilder(channel.Type)
		if err != nil {
			return err
		}
		builder, err := b.Build(channel.Type, 0)
		if err != nil {
			return gerror.New("不支持的渠道类型！")
		}
		err = builder.CustomPush(ctx, channel, model)
		if err != nil {
			return gerror.New("消息推送失败")
		}
	}
	return nil
}