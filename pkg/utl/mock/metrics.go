package mock

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// Metrics mock
type Metrics struct {
	InfoFn       func() ([]cpu.InfoStat, error)
	PercentFn    func(time.Duration, bool) ([]float64, error)
	TimesFn      func(bool) ([]cpu.TimesStat, error)
	SwapMemFn    func() (mem.SwapMemoryStat, error)
	VirtualMemFn func() (mem.VirtualMemoryStat, error)
}

// Info mock
func (m Metrics) Info() ([]cpu.InfoStat, error) {
	return m.InfoFn()
}

// Percent mock
func (m Metrics) Percent(interval time.Duration, perVCore bool) ([]float64, error) {
	return m.PercentFn(interval, perVCore)
}

// Times mock
func (m Metrics) Times(perVCore bool) ([]cpu.TimesStat, error) {
	return m.TimesFn(perVCore)
}

// SwapMemory mock
func (m Metrics) SwapMemory() (mem.SwapMemoryStat, error) {
	return m.SwapMemFn()
}

// VirtualMemory mock
func (m Metrics) VirtualMemory() (mem.VirtualMemoryStat, error) {
	return m.VirtualMemFn()
}
