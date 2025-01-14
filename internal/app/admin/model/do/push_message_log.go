// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PushMessageLog is the golang structure of table push_message_log for DAO operations like Where/Data.
type PushMessageLog struct {
	g.Meta     `orm:"table:push_message_log, do:true"`
	Id         interface{} // 主键 ID
	ChannelId  interface{} // 渠道 ID
	Status     interface{} // 0：失败 1 成功
	Message    interface{} // 消息内容
	PushType   interface{} // 推送类型
	CreateTime *gtime.Time // 推送时间
}
