package system

import (
	"github.com/shirou/gopsutil/cpu"
)

// VCore is a cpu entity on the current system
type VCore struct{}

// List returns two lists of cpu stats: percent and time usage per workload (in USER_HZ or Jiffies)
func (c VCore) List() ([]float64, []cpu.TimesStat, error) {
	percent, err := cpu.Percent(1, true)
	if err != nil {
		error.Error(err)
	}

	vcore, err := cpu.Times(true)
	if err != nil {
		error.Error(err)
	}

	return percent, vcore, err
}
