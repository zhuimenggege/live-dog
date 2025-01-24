package system

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
)

func SpaceMonitor() {
	diskInfo, _ := disk.Usage("/")
	if diskInfo.UsedPercent > 80 {
		
	}
	fmt.Println(diskInfo)
}
