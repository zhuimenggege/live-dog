// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PushChannelEmail is the golang structure of table push_channel_email for DAO operations like Where/Data.
type PushChannelEmail struct {
	g.Meta     `orm:"table:push_channel_email, do:true"`
	Id         interface{} // 主键 ID
	ChannelId  interface{} // 渠道 ID
	From       interface{} // 发送人
	To         interface{} // 接收人
	Server     interface{} // 发送服务器地址
	Port       interface{} // 发送端口
	AuthCode   interface{} // 授权码
	CreateTime *gtime.Time // 创建时间
	ActionTime *gtime.Time // 修改时间
}
