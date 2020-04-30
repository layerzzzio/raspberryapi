package system

import (
	"github.com/shirou/gopsutil/cpu"
)

// CPU is a cpu entity on the current system
type CPU struct{}

// List is a function
func (c CPU) List() ([]cpu.InfoStat, []float64, []cpu.TimesStat, error) {
	perVCore := false

	info, err := cpu.Info()
	if err != nil {
		return nil, nil, nil, err
	}

	percent, err := cpu.Percent(1, perVCore)
	if err != nil {
		return nil, nil, nil, err
	}

	vCore, err := cpu.Times(perVCore)
	if err != nil {
		return nil, nil, nil, err
	}

	return info, percent, vCore, nil
}
