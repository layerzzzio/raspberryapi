package mocksys

import (
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/mem"
)

// MEM mock
type MEM struct {
	ListFn func(mem.SwapMemoryStat, mem.VirtualMemoryStat) (rpi.MEM, error)
}

// List mock
func (m MEM) List(smem mem.SwapMemoryStat, vmem mem.VirtualMemoryStat) (rpi.MEM, error) {
	return m.ListFn(smem, vmem)
}
