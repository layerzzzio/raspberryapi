package cpu

import (
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/cpu"
)

// Service is a cpu service interface (controller)
type Service interface {
	List() ([]rpi.CPU, error)
}

// CPU represents a cpu application service (service)
type CPU struct {
	csys CSYS
}

// CSYS represents cpu data layer interface
type CSYS interface {
	List() ([]cpu.InfoStat, []float64, []cpu.TimesStat, error)
}

// New creates a cpu service
func New(csys CSYS) *CPU {
	return &CPU{csys: csys}
}
