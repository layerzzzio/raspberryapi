package mocksys

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
)

// LargestFile mock
type LargestFile struct {
	ViewFn func([]metrics.PathSize) ([]rpi.LargestFile, error)
}

// View mock
func (lf LargestFile) View(top100files []metrics.PathSize) ([]rpi.LargestFile, error) {
	return lf.ViewFn(top100files)
}
