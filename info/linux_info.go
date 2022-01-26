package info

import (
	"fmt"
	"github.com/cloudapp3/host/model"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	linux "github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"time"
)

func GetLinuxInfo() model.HostInfo {
	i := model.HostInfo{}
	u, uErr := linux.Info()
	if uErr != nil {
		fmt.Println("get linux info error " + uErr.Error())
		return i
	}
	i.UpTime = u.Uptime
	i.HostName = u.Hostname
	i.Platform = u.Platform + u.PlatformVersion
	d, dErr := disk.Usage("/")
	if dErr != nil {
		fmt.Println("get disk error " + dErr.Error())
		return i
	}
	i.TotalDisk = d.Total / 1024 / 1024
	i.FreeDisk = d.Free / 1024 / 1024
	c, cErr := cpu.Percent(time.Second*2, false)
	if cErr != nil {
		fmt.Println("get cpu percent error " + cErr.Error())
		return i
	}
	if len(c) != 0 {
		i.CPUUsage = float32(c[0])
	}
	m, mErr := mem.VirtualMemory()
	if mErr != nil {
		fmt.Println("get mem error " + mErr.Error())
		return i
	}
	i.FreeRam = m.Free / 1024 / 1024
	i.TotalRam = m.Total / 1024 / 1024
	return i
}
