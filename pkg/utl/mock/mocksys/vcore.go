package mocksys

import (
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/cpu"
)

// VCore mock
type VCore struct {
	ListFn func([]float64, []cpu.TimesStat) ([]rpi.VCore, error)
	ViewFn func(int, []float64, []cpu.TimesStat) (rpi.VCore, error)
}

// List mock
func (v VCore) List(percent []float64, times []cpu.TimesStat) ([]rpi.VCore, error) {
	return v.ListFn(percent, times)
}

// View mock
func (v VCore) View(id int, percent []float64, times []cpu.TimesStat) (rpi.VCore, error) {
	return v.ViewFn(id, percent, times)
}
