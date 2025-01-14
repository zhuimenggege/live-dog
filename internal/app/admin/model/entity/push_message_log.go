// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PushMessageLog is the golang structure for table push_message_log.
type PushMessageLog struct {
	Id         int         `json:"id"         orm:"id"          description:"主键 ID"`
	ChannelId  int         `json:"channelId"  orm:"channel_id"  description:"渠道 ID"`
	Status     int         `json:"status"     orm:"status"      description:"0：失败 1 成功"`
	Message    string      `json:"message"    orm:"message"     description:"消息内容"`
	PushType   string      `json:"pushType"   orm:"push_type"   description:"推送类型"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"推送时间"`
}
