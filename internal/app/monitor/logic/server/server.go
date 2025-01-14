package server

import (
	"context"
	"time"

	v1 "github.com/shichen437/live-dog/api/v1/monitor"
	"github.com/shichen437/live-dog/internal/app/monitor/model"
	"github.com/shichen437/live-dog/internal/app/monitor/service"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func init() {
	service.RegisterServerInfo(New())
}

func New() *sServerInfo {
	return &sServerInfo{}
}

type sServerInfo struct{}

func (s *sServerInfo) GetServerInfo(ctx context.Context, req *v1.GetServerInfoReq) (res *v1.GetServerInfoRes, err error) {
	res = &v1.GetServerInfoRes{}
	cpuInfo, _ := cpu.Info()
	percents, _ := cpu.Percent(100*time.Millisecond, false)
	var cpu *model.CpuInfo
	if len(cpuInfo) > 0 {
		gconv.Struct(cpuInfo[0], &cpu)
		if len(percents) > 0 {
			cpu.Percent = percents[0]
		}
	}
	memInfo, _ := mem.VirtualMemory()
	var mem *model.MemoryInfo
	if memInfo != nil {
		gconv.Struct(memInfo, &mem)
	}
	hostInfo, _ := host.Info()
	var sys *model.SystemInfo
	if hostInfo != nil {
		gconv.Struct(hostInfo, &sys)
	}
	diskInfo, _ := disk.Usage("/")
	var disks []model.DiskInfo
	if diskInfo != nil {
		var disk model.DiskInfo
		gconv.Struct(diskInfo, &disk)
		disks = append(disks, disk)
	}
	res.CpuInfo = cpu
	res.MemoryInfo = mem
	res.SystemInfo = sys
	res.DiskInfo = &disks
	return
}
