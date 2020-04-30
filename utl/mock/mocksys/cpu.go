package mocksys

import (
	"github.com/shirou/gopsutil/cpu"
)

// CPU is
type CPU struct {
	ListFn    func() ([]cpu.InfoStat, []float64, []cpu.TimesStat, error)
	CPUInfoFn func() ([]cpu.InfoStat, error)
}

// List mock
func (c *CPU) List() ([]cpu.InfoStat, []float64, []cpu.TimesStat, error) {
	return c.ListFn()
}

func (c *CPU) CPUInfo() ([]cpu.InfoStat, error) {
	return c.CPUInfo()
}
