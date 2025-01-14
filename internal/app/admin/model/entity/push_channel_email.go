// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PushChannelEmail is the golang structure for table push_channel_email.
type PushChannelEmail struct {
	Id         int         `json:"id"         orm:"id"          description:"主键 ID"`
	ChannelId  int         `json:"channelId"  orm:"channel_id"  description:"渠道 ID"`
	From       string      `json:"from"       orm:"from"        description:"发送人"`
	To         string      `json:"to"         orm:"to"          description:"接收人"`
	Server     string      `json:"server"     orm:"server"      description:"发送服务器地址"`
	Port       int         `json:"port"       orm:"port"        description:"发送端口"`
	AuthCode   string      `json:"authCode"   orm:"auth_code"   description:"授权码"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`
	ActionTime *gtime.Time `json:"actionTime" orm:"action_time" description:"修改时间"`
}
