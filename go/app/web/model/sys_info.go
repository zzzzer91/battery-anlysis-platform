package model

import "github.com/shirou/gopsutil/mem"

type Memory struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
}

type SysInfo struct {
	Memory Memory `json:"memory"`
}

func NewSysInfo() *SysInfo {
	vm, _ := mem.VirtualMemory()
	return &SysInfo{
		Memory: Memory{
			Total:       vm.Total,
			Free:        vm.Free,
			UsedPercent: vm.UsedPercent,
		},
	}
}
