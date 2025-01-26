// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PushChannelWeb is the golang structure for table push_channel_web.
type PushChannelWeb struct {
	Id           int         `json:"id"           orm:"id"            description:"记录 ID"`
	ChannelId    int         `json:"channelId"    orm:"channel_id"    description:"渠道 ID"`
	Url          string      `json:"url"          orm:"url"           description:"推送 URL"`
	HttpMethod   string      `json:"httpMethod"   orm:"http_method"   description:"请求方式"`
	Secret       string      `json:"secret"       orm:"secret"        description:"密钥/token/key"`
	AppId        string      `json:"appId"        orm:"app_id"        description:"应用 ID"`
	CorpId       string      `json:"corpId"       orm:"corp_id"       description:"企业 ID"`
	ReceiverId   string      `json:"receiverId"   orm:"receiver_id"   description:"接收人 ID"`
	ReceiverType string      `json:"receiverType" orm:"receiver_type" description:"接收人类型"`
	ExtraParams  string      `json:"extraParams"  orm:"extra_params"  description:"额外参数"`
	CreateTime   *gtime.Time `json:"createTime"   orm:"create_time"   description:"创建时间"`
	ActionTime   *gtime.Time `json:"actionTime"   orm:"action_time"   description:"修改时间"`
}
