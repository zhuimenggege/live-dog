// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictType is the golang structure for table sys_dict_type.
type SysDictType struct {
	DictId     int64       `json:"dictId"     orm:"dict_id"     description:"字典主键"`
	DictName   string      `json:"dictName"   orm:"dict_name"   description:"字典名称"`
	DictType   string      `json:"dictType"   orm:"dict_type"   description:"字典类型"`
	Status     string      `json:"status"     orm:"status"      description:"状态（0正常 1停用）"`
	CreateBy   string      `json:"createBy"   orm:"create_by"   description:"创建者"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`
	UpdateBy   string      `json:"updateBy"   orm:"update_by"   description:"更新者"`
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:"更新时间"`
	Remark     string      `json:"remark"     orm:"remark"      description:"备注"`
}
