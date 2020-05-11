package mocksys

import (
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/load"
)

// Load mock
type Load struct {
	ListFn func(load.AvgStat, load.MiscStat) (rpi.Load, error)
}

// List mock
func (l Load) List(temp load.AvgStat, procs load.MiscStat) (rpi.Load, error) {
	return l.ListFn(temp, procs)
}
