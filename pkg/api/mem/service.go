package mem

import (
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/mem"
)

// Service represents all MEM application services.
type Service interface {
	List() (rpi.MEM, error)
}

// MEM represents a MEM application service.
type MEM struct {
	msys MSYS
	mt   Metrics
}

// MSYS represents a MEM repository service.
type MSYS interface {
	List(mem.SwapMemoryStat, mem.VirtualMemoryStat) (rpi.MEM, error)
}

// Metrics represents the system metrics interface
type Metrics interface {
	SwapMemory() (mem.SwapMemoryStat, error)
	VirtualMemory() (mem.VirtualMemoryStat, error)
}

// New creates a MEM application service instance.
func New(msys MSYS, mt Metrics) *MEM {
	return &MEM{msys: msys, mt: mt}
}
