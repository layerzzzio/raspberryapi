package mocksys

import (
	"github.com/shirou/gopsutil/cpu"
)

// CPU is a mocked cpu
type CPU struct {
	ListFn func() ([]cpu.InfoStat, []float64, []cpu.TimesStat, error)
}

// List mock
func (c *CPU) List() ([]cpu.InfoStat, []float64, []cpu.TimesStat, error) {
	return c.ListFn()
}
