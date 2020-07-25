package largestfiles

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
)

// Service represents all LargestFiles application services.
type Service interface {
	List() ([]rpi.LargestFiles, error)
}

// LargestFiles represents a LargestFiles application service.
type LargestFiles struct {
	lfsys LFSYS
	mt    Metrics
}

// LFSYS represents a LargestFiles repository service.
type LFSYS interface {
	List([]metrics.PathSize) ([]rpi.LargestFiles, error)
}

// Metrics represents the system metrics interface
type Metrics interface {
	Top100Files() ([]metrics.PathSize, string, error)
}

// New creates a LargestFiles application service instance.
func New(lfsys LFSYS, mt Metrics) *LargestFiles {
	return &LargestFiles{lfsys: lfsys, mt: mt}
}
