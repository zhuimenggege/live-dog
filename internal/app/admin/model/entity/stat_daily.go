// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// StatDaily is the golang structure for table stat_daily.
type StatDaily struct {
	Id          int64       `json:"id"          orm:"id"           description:"记录ID"`
	Anchor      string      `json:"anchor"      orm:"anchor"       description:"主播"`
	DisplayName string      `json:"displayName" orm:"display_name" description:"展示名称"`
	DisplayType int         `json:"displayType" orm:"display_type" description:"展示类型（1 歌曲 2吉他）"`
	DisplayDate string      `json:"displayDate" orm:"display_date" description:"展示时间"`
	Count       uint        `json:"count"       orm:"count"        description:"次数"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
	CreateBy    string      `json:"createBy"    orm:"create_by"    description:"创建者"`
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"  description:"创建时间"`
	UpdateBy    string      `json:"updateBy"    orm:"update_by"    description:"更新者"`
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time"  description:"更新时间"`
	Action      uint        `json:"action"      orm:"action"       description:"标识：0 新增 1 修改 2 删除"`
}
