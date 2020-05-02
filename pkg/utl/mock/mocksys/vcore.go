package mocksys

import (
	"github.com/shirou/gopsutil/cpu"
)

// VCore is
type VCore struct {
	ListFn func() ([]float64, []cpu.TimesStat, error)
}

// List mock
func (v *VCore) List() ([]float64, []cpu.TimesStat, error) {
	return v.ListFn()
}
