// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysJob is the golang structure for table sys_job.
type SysJob struct {
	JobId          int64       `json:"jobId"          orm:"job_id"          description:"任务ID"`
	JobName        string      `json:"jobName"        orm:"job_name"        description:"任务名称"`
	JobGroup       string      `json:"jobGroup"       orm:"job_group"       description:"任务组名"`
	InvokeTarget   string      `json:"invokeTarget"   orm:"invoke_target"   description:"调用目标字符串"`
	CronExpression string      `json:"cronExpression" orm:"cron_expression" description:"cron执行表达式"`
	MisfirePolicy  string      `json:"misfirePolicy"  orm:"misfire_policy"  description:"计划执行错误策略（1立即执行 2执行一次 3放弃执行）"`
	Concurrent     string      `json:"concurrent"     orm:"concurrent"      description:"是否并发执行（0允许 1禁止）"`
	Status         string      `json:"status"         orm:"status"          description:"状态（0正常 1暂停）"`
	CreateBy       string      `json:"createBy"       orm:"create_by"       description:"创建者"`
	CreateTime     *gtime.Time `json:"createTime"     orm:"create_time"     description:"创建时间"`
	UpdateBy       string      `json:"updateBy"       orm:"update_by"       description:"更新者"`
	UpdateTime     *gtime.Time `json:"updateTime"     orm:"update_time"     description:"更新时间"`
	Remark         string      `json:"remark"         orm:"remark"          description:"备注信息"`
}
