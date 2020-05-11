package mock

import (
	"time"

	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

// Metrics mock
type Metrics struct {
	CPUInfoFn    func() ([]cpu.InfoStat, error)
	CPUPercentFn func(time.Duration, bool) ([]float64, error)
	CPUTimesFn   func(bool) ([]cpu.TimesStat, error)
	SwapMemFn    func() (mem.SwapMemoryStat, error)
	VirtualMemFn func() (mem.VirtualMemoryStat, error)
	DiskStatsFn  func(bool) (map[string][]metrics.DStats, error)
	LoadAvgFn    func() (load.AvgStat, error)
	LoadProcsFn  func() (load.MiscStat, error)
}

// CPUInfo mock
func (m Metrics) CPUInfo() ([]cpu.InfoStat, error) {
	return m.CPUInfoFn()
}

// CPUPercent mock
func (m Metrics) CPUPercent(interval time.Duration, perVCore bool) ([]float64, error) {
	return m.CPUPercentFn(interval, perVCore)
}

// CPUTimes mock
func (m Metrics) CPUTimes(perVCore bool) ([]cpu.TimesStat, error) {
	return m.CPUTimesFn(perVCore)
}

// SwapMemory mock
func (m Metrics) SwapMemory() (mem.SwapMemoryStat, error) {
	return m.SwapMemFn()
}

// VirtualMemory mock
func (m Metrics) VirtualMemory() (mem.VirtualMemoryStat, error) {
	return m.VirtualMemFn()
}

// DiskStats mock
func (m Metrics) DiskStats(all bool) (map[string][]metrics.DStats, error) {
	return m.DiskStatsFn(all)
}

// LoadAvg mock
func (m Metrics) LoadAvg() (load.AvgStat, error) {
	return m.LoadAvgFn()
}

// LoadProcs mock
func (m Metrics) LoadProcs() (load.MiscStat, error) {
	return m.LoadProcsFn()
}
