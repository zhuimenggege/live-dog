// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	RoleId            int64       `json:"roleId"            orm:"role_id"             description:"角色ID"`
	RoleName          string      `json:"roleName"          orm:"role_name"           description:"角色名称"`
	RoleKey           string      `json:"roleKey"           orm:"role_key"            description:"角色权限字符串"`
	RoleSort          int         `json:"roleSort"          orm:"role_sort"           description:"显示顺序"`
	DataScope         string      `json:"dataScope"         orm:"data_scope"          description:"数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）"`
	MenuCheckStrictly int         `json:"menuCheckStrictly" orm:"menu_check_strictly" description:"菜单树选择项是否关联显示"`
	Status            string      `json:"status"            orm:"status"              description:"角色状态（0正常 1停用）"`
	DelFlag           string      `json:"delFlag"           orm:"del_flag"            description:"删除标志（0代表存在 2代表删除）"`
	CreateBy          string      `json:"createBy"          orm:"create_by"           description:"创建者"`
	CreateTime        *gtime.Time `json:"createTime"        orm:"create_time"         description:"创建时间"`
	UpdateBy          string      `json:"updateBy"          orm:"update_by"           description:"更新者"`
	UpdateTime        *gtime.Time `json:"updateTime"        orm:"update_time"         description:"更新时间"`
	Remark            string      `json:"remark"            orm:"remark"              description:"备注"`
}
