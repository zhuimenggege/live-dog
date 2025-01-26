// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PushChannel is the golang structure of table push_channel for DAO operations like Where/Data.
type PushChannel struct {
	g.Meta     `orm:"table:push_channel, do:true"`
	Id         interface{} // 主键 ID
	Name       interface{} // 渠道名称
	Type       interface{} // 渠道类型
	Status     interface{} // 状态：0 禁用 1 启用
	Remark     interface{} // 备注
	CreateBy   interface{} // 创建人
	CreateTime *gtime.Time // 创建时间
	ActionTime *gtime.Time // 修改时间
}
