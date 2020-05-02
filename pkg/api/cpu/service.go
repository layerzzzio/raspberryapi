package cpu

import (
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
}

// CSYS represents a CPU repository service.
type CSYS interface {
	List() ([]cpu.InfoStat, []float64, []cpu.TimesStat, error)
}

// New creates a CPU application service instance.
func New(csys CSYS) *CPU {
	return &CPU{csys: csys}
}
