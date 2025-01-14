// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoomInfo is the golang structure for table room_info.
type RoomInfo struct {
	Id         int         `json:"id"         orm:"id"          description:"房间信息 ID"`
	LiveId     int         `json:"liveId"     orm:"live_id"     description:"房间 ID"`
	RoomName   string      `json:"roomName"   orm:"room_name"   description:"房间名称"`
	Anchor     string      `json:"anchor"     orm:"anchor"      description:"主播"`
	Platform   string      `json:"platform"   orm:"platform"    description:"直播平台"`
	Status     int         `json:"status"     orm:"status"      description:"状态"`
	CreateBy   string      `json:"createBy"   orm:"create_by"   description:"创建者"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`
	ActionTime *gtime.Time `json:"actionTime" orm:"action_time" description:"修改时间"`
}
