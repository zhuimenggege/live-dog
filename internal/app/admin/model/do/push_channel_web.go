// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PushChannelWeb is the golang structure of table push_channel_web for DAO operations like Where/Data.
type PushChannelWeb struct {
	g.Meta       `orm:"table:push_channel_web, do:true"`
	Id           interface{} // 记录 ID
	ChannelId    interface{} // 渠道 ID
	Url          interface{} // 推送 URL
	HttpMethod   interface{} // 请求方式
	Secret       interface{} // 密钥/token/key
	AppId        interface{} // 应用 ID
	CorpId       interface{} // 企业 ID
	ReceiverId   interface{} // 接收人 ID
	ReceiverType interface{} // 接收人类型
	ExtraParams  interface{} // 额外参数
	CreateTime   *gtime.Time // 创建时间
	ActionTime   *gtime.Time // 修改时间
}
