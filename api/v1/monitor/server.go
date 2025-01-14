package v1

import (
	"github.com/shichen437/live-dog/internal/app/monitor/model"

	"github.com/gogf/gf/v2/frame/g"
)

type GetServerInfoReq struct {
	g.Meta `path:"/monitor/server" method:"get" tags:"服务监控" summary:"服务器详情"`
}
type GetServerInfoRes struct {
	g.Meta     `mime:"application/json"`
	CpuInfo    *model.CpuInfo    `json:"cpu"`
	MemoryInfo *model.MemoryInfo `json:"mem"`
	SystemInfo *model.SystemInfo `json:"sys"`
	DiskInfo   *[]model.DiskInfo `json:"disks"`
}
