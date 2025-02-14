// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysJobLog is the golang structure for table sys_job_log.
type SysJobLog struct {
	JobLogId      int64       `json:"jobLogId"      orm:"job_log_id"     description:"任务日志ID"`
	JobId         int64       `json:"jobId"         orm:"job_id"         description:"对应任务 ID"`
	JobName       string      `json:"jobName"       orm:"job_name"       description:"任务名称"`
	InvokeTarget  string      `json:"invokeTarget"  orm:"invoke_target"  description:"调用目标字符串"`
	JobMessage    string      `json:"jobMessage"    orm:"job_message"    description:"日志信息"`
	Status        string      `json:"status"        orm:"status"         description:"执行状态（0正常 1失败）"`
	ExceptionInfo string      `json:"exceptionInfo" orm:"exception_info" description:"异常信息"`
	CreateTime    *gtime.Time `json:"createTime"    orm:"create_time"    description:"创建时间"`
}
