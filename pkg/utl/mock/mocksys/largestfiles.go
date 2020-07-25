package mocksys

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
)

// LargestFiles mock
type LargestFiles struct {
	ListFn func([]metrics.PathSize) ([]rpi.LargestFiles, error)
}

// List mock
func (lf LargestFiles) List(top100files []metrics.PathSize) ([]rpi.LargestFiles, error) {
	return lf.ListFn(top100files)
}
