// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// StatDaily is the golang structure of table stat_daily for DAO operations like Where/Data.
type StatDaily struct {
	g.Meta      `orm:"table:stat_daily, do:true"`
	Id          interface{} // 记录ID
	Anchor      interface{} // 主播
	DisplayName interface{} // 展示名称
	DisplayType interface{} // 展示类型（1 歌曲 2吉他）
	DisplayDate interface{} // 展示时间
	Count       interface{} // 次数
	Remark      interface{} // 备注
	CreateBy    interface{} // 创建者
	CreateTime  *gtime.Time // 创建时间
	UpdateBy    interface{} // 更新者
	UpdateTime  *gtime.Time // 更新时间
	Action      interface{} // 标识：0 新增 1 修改 2 删除
}
