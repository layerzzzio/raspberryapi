package cpu

import (
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/cpu"
)

// Service represents all CPU application services.
type Service interface {
	List() ([]rpi.CPU, error)
	View(int) (rpi.CPU, error)
}

// CPU represents a CPU application service.
type CPU struct {
	csys CSYS
	m    Metrics
}

// CSYS represents a CPU repository service.
type CSYS interface {
	List([]cpu.InfoStat, []float64, []cpu.TimesStat) ([]rpi.CPU, error)
	View(int, []cpu.InfoStat, []float64, []cpu.TimesStat) (rpi.CPU, error)
}

// Metrics represents the system metrics interface
type Metrics interface {
	CPUInfo() ([]cpu.InfoStat, error)
	CPUPercent(time.Duration, bool) ([]float64, error)
	CPUTimes(bool) ([]cpu.TimesStat, error)
}

// New creates a CPU application service instance.
func New(csys CSYS, m Metrics) *CPU {
	return &CPU{csys: csys, m: m}
}
