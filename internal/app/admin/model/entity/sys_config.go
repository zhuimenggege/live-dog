// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure for table sys_config.
type SysConfig struct {
	ConfigId    int         `json:"configId"    orm:"config_id"    description:"参数主键"`
	ConfigName  string      `json:"configName"  orm:"config_name"  description:"参数名称"`
	ConfigKey   string      `json:"configKey"   orm:"config_key"   description:"参数键名"`
	ConfigValue string      `json:"configValue" orm:"config_value" description:"参数键值"`
	ConfigType  string      `json:"configType"  orm:"config_type"  description:"系统内置（Y是 N否）"`
	CreateBy    string      `json:"createBy"    orm:"create_by"    description:"创建者"`
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"  description:"创建时间"`
	UpdateBy    string      `json:"updateBy"    orm:"update_by"    description:"更新者"`
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time"  description:"更新时间"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
}
