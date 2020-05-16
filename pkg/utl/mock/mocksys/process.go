package mocksys

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
)

// Process mock
type Process struct {
	ListFn func([]metrics.PInfo) ([]rpi.ProcessSummary, error)
	ViewFn func(int32, []metrics.PInfo) (rpi.ProcessDetails, error)
}

// List mock
func (p Process) List(pinfo []metrics.PInfo) ([]rpi.ProcessSummary, error) {
	return p.ListFn(pinfo)
}

// View mock
func (p Process) View(id int32, pinfo []metrics.PInfo) (rpi.ProcessDetails, error) {
	return p.ViewFn(id, pinfo)
}
