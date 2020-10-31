package largestfile

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
)

// Service represents all LargestFile application services.
type Service interface {
	View(path string) ([]rpi.LargestFile, error)
}

// LargestFile represents a LargestFile application service.
type LargestFile struct {
	lfsys LFSYS
	mt    Metrics
}

// LFSYS represents a LargestFile repository service.
type LFSYS interface {
	View([]metrics.PathSize) ([]rpi.LargestFile, error)
}

// Metrics represents the system metrics interface
type Metrics interface {
	Top100Files(path string) ([]metrics.PathSize, string, error)
}

// New creates a LargestFile application service instance.
func New(lfsys LFSYS, mt Metrics) *LargestFile {
	return &LargestFile{lfsys: lfsys, mt: mt}
}