package mocksys

import (
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/mem"
)

// Mem mock
type Mem struct {
	ListFn func(mem.SwapMemoryStat, mem.VirtualMemoryStat) (rpi.Mem, error)
}

// List mock
func (m Mem) List(smem mem.SwapMemoryStat, vmem mem.VirtualMemoryStat) (rpi.Mem, error) {
	return m.ListFn(smem, vmem)
}
