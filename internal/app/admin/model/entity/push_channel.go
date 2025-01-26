// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PushChannel is the golang structure for table push_channel.
type PushChannel struct {
	Id         int         `json:"id"         orm:"id"          description:"主键 ID"`
	Name       string      `json:"name"       orm:"name"        description:"渠道名称"`
	Type       string      `json:"type"       orm:"type"        description:"渠道类型"`
	Status     int         `json:"status"     orm:"status"      description:"状态：0 禁用 1 启用"`
	Remark     string      `json:"remark"     orm:"remark"      description:"备注"`
	CreateBy   string      `json:"createBy"   orm:"create_by"   description:"创建人"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`
	ActionTime *gtime.Time `json:"actionTime" orm:"action_time" description:"修改时间"`
}
