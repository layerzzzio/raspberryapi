package host

import (
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// Service represents all Host application services.
type Service interface {
	List() (rpi.Host, error)
}

// Host represents a Host application service.
type Host struct {
	hsys HSYS
	mt   Metrics
}

// HSYS represents a Host repository service.
type HSYS interface {
	List(host.InfoStat,
		[]host.UserStat,
		[]cpu.InfoStat,
		[]float64,
		mem.VirtualMemoryStat,
		mem.SwapMemoryStat) (rpi.Host, error)
}

// Metrics represents the system metrics interface
type Metrics interface {
	HostInfo() (host.InfoStat, error)
	Users() ([]host.UserStat, error)
	CPUInfo() ([]cpu.InfoStat, error)
	CPUPercent(interval time.Duration, perVCore bool) ([]float64, error)
	VirtualMemory() (mem.VirtualMemoryStat, error)
	SwapMemory() (mem.SwapMemoryStat, error)
}

// New creates a Host application service instance.
func New(hsys HSYS, mt Metrics) *Host {
	return &Host{hsys: hsys, mt: mt}
}
