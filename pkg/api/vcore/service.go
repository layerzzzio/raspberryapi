package vcore

import (
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/cpu"
)

// Service interface
type Service interface {
	List() ([]rpi.VCore, error)
	View(int) (rpi.VCore, error)
}

// VCore represents an application service
type VCore struct {
	vsys VSYS
	m    Metrics
}

// VSYS represents the vcore repository interface
type VSYS interface {
	List([]float64, []cpu.TimesStat) ([]rpi.VCore, error)
	View(int, []float64, []cpu.TimesStat) (rpi.VCore, error)
}

// Metrics represents the system metrics interface
type Metrics interface {
	CPUPercent(time.Duration, bool) ([]float64, error)
	CPUTimes(bool) ([]cpu.TimesStat, error)
}

// New creates a core service
func New(vsys VSYS, m Metrics) *VCore {
	return &VCore{vsys: vsys, m: m}
}
