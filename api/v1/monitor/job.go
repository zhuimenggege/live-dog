package v1

import (
	"github.com/shichen437/live-dog/api/v1/common"
	"github.com/shichen437/live-dog/internal/app/monitor/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetJobListReq struct {
	g.Meta `path:"/monitor/job/list" method:"get" tags:"定时任务" summary:"定时任务列表"`
	common.PageReq
	JobName  string `p:"jobName"`
	JobGroup string `p:"jobGroup"`
	Status   int    `p:"status"`
}

type GetJobListRes struct {
	g.Meta `mime:"application/json"`
	Rows   []*entity.SysJob `json:"rows"`
	Total  int              `json:"total"`
}

type GetJobLogListReq struct {
	g.Meta `path:"/monitor/jobLog/list" method:"get" tags:"定时任务" summary:"定时任务日志列表"`
	common.PageReq
	JobName  string `p:"jobName"`
	JobGroup string `p:"jobGroup"`
	Status   int    `p:"status"`
}

type GetJobLogListRes struct {
	g.Meta `mime:"application/json"`
	Rows   []*entity.SysJobLog `json:"rows"`
	Total  int                 `json:"total"`
}
