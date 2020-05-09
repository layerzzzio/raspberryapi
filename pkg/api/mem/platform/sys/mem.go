package sys

import (
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/mem"
)

// Mem represents an empty MEM entity on the current system.
type Mem struct{}

// List returns a list of MEM stats including some swap and virtual memory
func (c Mem) List(swap mem.SwapMemoryStat, vmem mem.VirtualMemoryStat) (rpi.Mem, error) {
	result := rpi.Mem{
		STotal:       swap.Total,
		SUsed:        swap.Used,
		SFree:        swap.Free,
		SUsedPercent: swap.UsedPercent,
		VTotal:       vmem.Total,
		VUsed:        vmem.Used,
		VAvailable:   vmem.Available,
		VUsedPercent: vmem.UsedPercent,
	}

	return result, nil
}
