package v1

import (
	"github.com/shichen437/live-dog/api/v1/common"
	"github.com/shichen437/live-dog/internal/app/admin/model"
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetPushChannelListReq struct {
	g.Meta `path:"/system/push/channel/list" method:"get" tags:"消息推送" summary:"渠道列表"`
	common.PageReq
	Name string `p:"name"`
	Type string `p:"type"`
}
type GetPushChannelListRes struct {
	g.Meta `mime:"application/json"`
	Rows   []*entity.PushChannel `json:"rows"`
	Total  int                   `json:"total"`
}

type PostPushChannelReq struct {
	g.Meta `path:"/system/push/channel" method:"post" tags:"消息推送" summary:"添加渠道"`
	*model.PushChannel
}
type PostPushChannelRes struct {
	g.Meta `mime:"application/json"`
}

type PutPushChannelReq struct {
	g.Meta `path:"/system/push/channel" method:"put" tags:"消息推送" summary:"修改渠道"`
	*model.PushChannel
}
type PutPushChannelRes struct {
	g.Meta `mime:"application/json"`
}

type GetPushChannelReq struct {
	g.Meta `path:"/system/push/channel/{id}" method:"get" tags:"消息推送" summary:"获取渠道信息"`
	Id     int `p:"id"  v:"required"`
}
type GetPushChannelRes struct {
	g.Meta `mime:"application/json"`
	*model.PushChannel
}

type DeletePushChannelReq struct {
	g.Meta `path:"/system/push/channel/{id}" method:"delete" tags:"消息推送" summary:"删除渠道"`
	Id     string `p:"id"  v:"required"`
}
type DeletePushChannelRes struct {
	g.Meta `mime:"application/json"`
}
