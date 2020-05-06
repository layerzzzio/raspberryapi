package cpu

import (
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/cpu"
)

// Service represents all CPU application services.
type Service interface {
	List() ([]rpi.CPU, error)
}

// CPU represents a CPU application service.
type CPU struct {
	csys CSYS
	m    Metrics
}

// CSYS represents a CPU repository service.
type CSYS interface {
	List([]cpu.InfoStat, []float64, []cpu.TimesStat) ([]rpi.CPU, error)
}

// Metrics represents the system metrics interface
type Metrics interface {
	Info() ([]cpu.InfoStat, error)
	Percent(time.Duration, bool) ([]float64, error)
	Times(bool) ([]cpu.TimesStat, error)
}

// New creates a CPU application service instance.
func New(csys CSYS, m Metrics) *CPU {
	return &CPU{csys: csys, m: m}
}
