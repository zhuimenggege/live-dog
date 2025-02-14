package v1

import (
	"github.com/shichen437/live-dog/api/v1/common"
	"github.com/shichen437/live-dog/internal/app/monitor/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetJobListReq struct {
	g.Meta `path:"/monitor/job/list" method:"get" tags:"定时任务" summary:"定时任务列表"`
	common.PageReq
	JobName string `p:"jobName"`
	Status  *int   `p:"status"`
}

type GetJobListRes struct {
	g.Meta `mime:"application/json"`
	Rows   []*entity.SysJob `json:"rows"`
	Total  int              `json:"total"`
}

type GetJobLogListReq struct {
	g.Meta `path:"/monitor/jobLog/list" method:"get" tags:"定时任务" summary:"定时任务日志列表"`
	common.PageReq
	JobName   string `p:"jobName"`
	Status    *int   `p:"status"`
	JobId     int64  `p:"jobId"`
	BeginTime string `p:"params[beginTime]"`
	EndTime   string `p:"params[endTime]"`
}

type GetJobLogListRes struct {
	g.Meta `mime:"application/json"`
	Rows   []*entity.SysJobLog `json:"rows"`
	Total  int                 `json:"total"`
}

type PostJobReq struct {
	g.Meta         `path:"/monitor/job" method:"post" tags:"定时任务" summary:"新增定时任务"`
	JobName        string `p:"jobName"  v:"required"`
	CronExpression string `p:"cron"  v:"required"`
	InvokeTarget   string `p:"invokeTarget"  v:"required"`
	MisfirePolicy  string `p:"misfirePolicy"  v:"required"`
	Concurrent     string `p:"concurrent"  v:"required"`
	Status         string `p:"status"  v:"required"`
	Remark         string `p:"remark"`
	JobParams      string `p:"jobParams"`
}

type PostJobRes struct {
	g.Meta `mime:"application/json"`
}

type PutJobReq struct {
	g.Meta         `path:"/monitor/job" method:"put" tags:"定时任务" summary:"修改定时任务"`
	JobId          int64  `p:"jobId"  v:"required"`
	JobName        string `p:"jobName"  v:"required"`
	CronExpression string `p:"cron"  v:"required"`
	InvokeTarget   string `p:"invokeTarget"  v:"required"`
	MisfirePolicy  string `p:"misfirePolicy"  v:"required"`
	Concurrent     string `p:"concurrent"  v:"required"`
	Status         string `p:"status"  v:"required"`
	Remark         string `p:"remark"`
	JobParams      string `p:"jobParams"`
}

type PutJobRes struct {
	g.Meta `mime:"application/json"`
}

type GetJobDetailReq struct {
	g.Meta `path:"/monitor/job/{jobId}" method:"get" tags:"定时任务" summary:"定时任务详情"`
	JobId  int64 `p:"jobId"  v:"required"`
}

type GetJobDetailRes struct {
	g.Meta `mime:"application/json"`
	*entity.SysJob
}

type DeleteJobReq struct {
	g.Meta `path:"/monitor/job" method:"delete" tags:"定时任务" summary:"删除定时任务"`
	JobId  string `p:"jobId"  v:"required"`
}

type DeleteJobRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteJobLogReq struct {
	g.Meta `path:"/monitor/jobLog/{jobLogId}" method:"delete" tags:"定时任务" summary:"删除定时任务日志"`
	JobLogId  string `p:"jobLogId"  v:"required"`
}

type DeleteJobLogRes struct {
	g.Meta `mime:"application/json"`
}

type PutJobStatusReq struct {
	g.Meta `path:"/monitor/job/changeStatus" method:"put" tags:"定时任务" summary:"修改定时任务状态"`
	JobId  int64  `p:"jobId"  v:"required"`
	Status string `p:"status"  v:"required"`
}

type PutJobStatusRes struct {
	g.Meta `mime:"application/json"`
}

type PutJobRunReq struct {
	g.Meta `path:"/monitor/job/run" method:"put" tags:"定时任务" summary:"手动触发任务"`
	JobId  int64 `p:"jobId"  v:"required"`
}

type PutJobRunRes struct {
	g.Meta `mime:"application/json"`
}
