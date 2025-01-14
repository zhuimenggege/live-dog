// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOperLog is the golang structure for table sys_oper_log.
type SysOperLog struct {
	OperId        int64       `json:"operId"        orm:"oper_id"        description:"日志主键"`
	Title         string      `json:"title"         orm:"title"          description:"模块标题"`
	BusinessType  int         `json:"businessType"  orm:"business_type"  description:"业务类型（0其它 1新增 2修改 3删除）"`
	Method        string      `json:"method"        orm:"method"         description:"方法名称"`
	RequestMethod string      `json:"requestMethod" orm:"request_method" description:"请求方式"`
	OperatorType  int         `json:"operatorType"  orm:"operator_type"  description:"操作类别（0其它 1后台用户 2手机端用户）"`
	OperName      string      `json:"operName"      orm:"oper_name"      description:"操作人员"`
	OperUrl       string      `json:"operUrl"       orm:"oper_url"       description:"请求URL"`
	OperIp        string      `json:"operIp"        orm:"oper_ip"        description:"主机地址"`
	OperLocation  string      `json:"operLocation"  orm:"oper_location"  description:"操作地点"`
	OperParam     string      `json:"operParam"     orm:"oper_param"     description:"请求参数"`
	JsonResult    string      `json:"jsonResult"    orm:"json_result"    description:"返回参数"`
	Status        int         `json:"status"        orm:"status"         description:"操作状态（0正常 1异常）"`
	ErrorMsg      string      `json:"errorMsg"      orm:"error_msg"      description:"错误消息"`
	OperTime      *gtime.Time `json:"operTime"      orm:"oper_time"      description:"操作时间"`
	CostTime      int64       `json:"costTime"      orm:"cost_time"      description:"消耗时间"`
}
