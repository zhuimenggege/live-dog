// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LiveCookie is the golang structure for table live_cookie.
type LiveCookie struct {
	Id         int         `json:"id"         orm:"id"          description:"ID"`
	Platform   string      `json:"platform"   orm:"platform"    description:"平台"`
	Cookie     string      `json:"cookie"     orm:"cookie"      description:"cookie"`
	Remark     string      `json:"remark"     orm:"remark"      description:"备注"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`
	ActionTime *gtime.Time `json:"actionTime" orm:"action_time" description:"更新时间"`
}
