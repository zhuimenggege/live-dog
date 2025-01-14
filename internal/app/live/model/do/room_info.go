// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoomInfo is the golang structure of table room_info for DAO operations like Where/Data.
type RoomInfo struct {
	g.Meta     `orm:"table:room_info, do:true"`
	Id         interface{} // 房间信息 ID
	LiveId     interface{} // 房间 ID
	RoomName   interface{} // 房间名称
	Anchor     interface{} // 主播
	Platform   interface{} // 直播平台
	Status     interface{} // 状态
	CreateBy   interface{} // 创建者
	CreateTime *gtime.Time // 创建时间
	ActionTime *gtime.Time // 修改时间
}
