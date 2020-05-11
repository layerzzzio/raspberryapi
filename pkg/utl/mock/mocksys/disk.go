package mocksys

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
)

// Disk mock
type Disk struct {
	ListFn func(map[string][]metrics.DStats) ([]rpi.Disk, error)
	ViewFn func(string, map[string][]metrics.DStats) (rpi.Disk, error)
}

// List mock
func (d *Disk) List(dstats map[string][]metrics.DStats) ([]rpi.Disk, error) {
	return d.ListFn(dstats)
}

// View mock
func (d *Disk) View(id string, dstats map[string][]metrics.DStats) (rpi.Disk, error) {
	return d.ViewFn(id, dstats)
}
