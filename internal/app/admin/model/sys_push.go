package model

import (
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type PushChannel struct {
	g.Meta `orm:"table:push_channel" desc:"消息推送渠道"`
	*entity.PushChannel
	Email *entity.PushChannelEmail `json:"email"  desc:"邮箱信息"`
	Web   *entity.PushChannelWeb   `json:"web"    desc:"web 信息"`
}
