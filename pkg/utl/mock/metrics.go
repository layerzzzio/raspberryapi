package mock

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// Metrics mock
type Metrics struct {
	InfoFn    func() ([]cpu.InfoStat, error)
	PercentFn func(time.Duration, bool) ([]float64, error)
	TimesFn   func(bool) ([]cpu.TimesStat, error)
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
