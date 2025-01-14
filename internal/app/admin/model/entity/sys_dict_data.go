// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictData is the golang structure for table sys_dict_data.
type SysDictData struct {
	DictCode   int64       `json:"dictCode"   orm:"dict_code"   description:"字典编码"`
	DictSort   int         `json:"dictSort"   orm:"dict_sort"   description:"字典排序"`
	DictLabel  string      `json:"dictLabel"  orm:"dict_label"  description:"字典标签"`
	DictValue  string      `json:"dictValue"  orm:"dict_value"  description:"字典键值"`
	DictType   string      `json:"dictType"   orm:"dict_type"   description:"字典类型"`
	CssClass   string      `json:"cssClass"   orm:"css_class"   description:"样式属性（其他样式扩展）"`
	ListClass  string      `json:"listClass"  orm:"list_class"  description:"表格回显样式"`
	IsDefault  string      `json:"isDefault"  orm:"is_default"  description:"是否默认（Y是 N否）"`
	Status     string      `json:"status"     orm:"status"      description:"状态（0正常 1停用）"`
	CreateBy   string      `json:"createBy"   orm:"create_by"   description:"创建者"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`
	UpdateBy   string      `json:"updateBy"   orm:"update_by"   description:"更新者"`
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:"更新时间"`
	Remark     string      `json:"remark"     orm:"remark"      description:"备注"`
}
