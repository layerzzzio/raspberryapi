package vcore

import (
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/cpu"
)

// Service is a core service interface (controller)
type Service interface {
	List() ([]rpi.VCore, error)
	View(int) (*rpi.VCore, error)
}

// VCore represents a core application service (service)
type VCore struct {
	vsys VSYS
}

// VSYS represents core data layer interface
type VSYS interface {
	List() ([]float64, []cpu.TimesStat, error)
}

// New creates a core service
func New(vsys VSYS) *VCore {
	return &VCore{vsys: vsys}
}
