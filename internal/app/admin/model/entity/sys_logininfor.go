// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLogininfor is the golang structure for table sys_logininfor.
type SysLogininfor struct {
	InfoId        int64       `json:"infoId"        orm:"info_id"        description:"访问ID"`
	UserName      string      `json:"userName"      orm:"user_name"      description:"用户账号"`
	Ipaddr        string      `json:"ipaddr"        orm:"ipaddr"         description:"登录IP地址"`
	LoginLocation string      `json:"loginLocation" orm:"login_location" description:"登录地点"`
	Browser       string      `json:"browser"       orm:"browser"        description:"浏览器类型"`
	Os            string      `json:"os"            orm:"os"             description:"操作系统"`
	Status        string      `json:"status"        orm:"status"         description:"登录状态（0成功 1失败）"`
	Msg           string      `json:"msg"           orm:"msg"            description:"提示消息"`
	LoginTime     *gtime.Time `json:"loginTime"     orm:"login_time"     description:"访问时间"`
}
