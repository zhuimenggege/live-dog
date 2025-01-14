// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// LiveManage is the golang structure of table live_manage for DAO operations like Where/Data.
type LiveManage struct {
	g.Meta       `orm:"table:live_manage, do:true"`
	Id           interface{} // 房间 id
	RoomUrl      interface{} // 房间 url
	Interval     interface{} // 轮询间隔
	Format       interface{} // 导出视频格式
	EnableNotice interface{} // 启用通知
	MonitorType  interface{} // 监控类型
	MonitorStart interface{} // 监控开始时间
	MonitorStop  interface{} // 监控结束时间
	Remark       interface{} // 房间备注
	CreateBy     interface{} // 创建人
	CreateTime   *gtime.Time // 创建时间
	ActionBy     interface{} // 修改人
	ActionTime   *gtime.Time // 修改时间
}
