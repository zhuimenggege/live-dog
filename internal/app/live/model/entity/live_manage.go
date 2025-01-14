// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LiveManage is the golang structure for table live_manage.
type LiveManage struct {
	Id           int         `json:"id"           orm:"id"            description:"房间 id"`
	RoomUrl      string      `json:"roomUrl"      orm:"room_url"      description:"房间 url"`
	Interval     uint        `json:"interval"     orm:"interval"      description:"轮询间隔"`
	Format       string      `json:"format"       orm:"format"        description:"导出视频格式"`
	EnableNotice int         `json:"enableNotice" orm:"enable_notice" description:"启用通知"`
	MonitorType  int         `json:"monitorType"  orm:"monitor_type"  description:"监控类型"`
	MonitorStart string      `json:"monitorStart" orm:"monitor_start" description:"监控开始时间"`
	MonitorStop  string      `json:"monitorStop"  orm:"monitor_stop"  description:"监控结束时间"`
	Remark       string      `json:"remark"       orm:"remark"        description:"房间备注"`
	CreateBy     string      `json:"createBy"     orm:"create_by"     description:"创建人"`
	CreateTime   *gtime.Time `json:"createTime"   orm:"create_time"   description:"创建时间"`
	ActionBy     string      `json:"actionBy"     orm:"action_by"     description:"修改人"`
	ActionTime   *gtime.Time `json:"actionTime"   orm:"action_time"   description:"修改时间"`
}
