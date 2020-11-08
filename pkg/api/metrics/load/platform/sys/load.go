package sys

import (
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/load"
)

// Load represents an empty MEM entity on the current system.
type Load struct{}

// List returns a list of Load stats
func (l Load) List(temp load.AvgStat, procs load.MiscStat) (rpi.Load, error) {
	result := rpi.Load{
		Load1:        temp.Load1,
		Load5:        temp.Load5,
		Load15:       temp.Load15,
		ProcsTotal:   procs.ProcsTotal,
		ProcsRunning: procs.ProcsRunning,
		ProcsBlocked: procs.ProcsBlocked,
	}

	return result, nil
}
