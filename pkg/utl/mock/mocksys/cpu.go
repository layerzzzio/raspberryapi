package mocksys

import (
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/cpu"
)

// CPU mock
type CPU struct {
	ListFn func([]cpu.InfoStat, []float64, []cpu.TimesStat) ([]rpi.CPU, error)
}

// List mock
func (c *CPU) List(info []cpu.InfoStat, percent []float64, times []cpu.TimesStat) ([]rpi.CPU, error) {
	return c.ListFn(info, percent, times)
}
