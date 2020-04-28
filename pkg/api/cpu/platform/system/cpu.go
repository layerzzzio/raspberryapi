package system

import (
	"github.com/shirou/gopsutil/cpu"
)

// CPU is a cpu entity on the current system
type CPU struct{}

// List is a function
func (c CPU) List() ([]cpu.InfoStat, []float64, []cpu.TimesStat, error) {
	info, err := cpu.Info()
	if err != nil {
		error.Error(err)
	}

	percent, err := cpu.Percent(1, false)
	if err != nil {
		error.Error(err)
	}

	vCore, err := cpu.Times(false)
	if err != nil {
		error.Error(err)
	}

	return info, percent, vCore, nil

}
