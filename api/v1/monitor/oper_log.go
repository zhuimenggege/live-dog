package v1

import (
	"github.com/shichen437/live-dog/api/v1/common"
	"github.com/shichen437/live-dog/internal/app/monitor/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetOperLogListReq struct {
	g.Meta `path:"/monitor/operlog/list" method:"get" tags:"操作日志" summary:"current data list"`
	common.PageReq
	Title        string `p:"title"`
	OperName     string `p:"operName"`
	BusinessType int    `p:"businessType"`
	Status       int    `p:"status"`
}
type GetOperLogListRes struct {
	g.Meta `mime:"application/json"`
	Rows   []*entity.SysOperLog `json:"rows"`
	Total  int                  `json:"total"`
}
