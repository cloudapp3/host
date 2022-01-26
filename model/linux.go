package model

type HostInfo struct {
	HostName       string
	UpTime         uint64
	Platform       string
	TotalRam       uint64
	FreeRam        uint64
	TotalDisk      uint64
	FreeDisk       uint64
	CPUUsage       float32
}
