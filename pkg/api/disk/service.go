package disk

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
)

// Service represents all disk application services.
type Service interface {
	List() ([]rpi.Disk, error)
	View(string) (rpi.Disk, error)
}

// Disk represents a disk application service.
type Disk struct {
	dsys DSYS
	m    Metrics
}

// DSYS represents a disk repository service.
type DSYS interface {
	List(map[string][]metrics.DStats) ([]rpi.Disk, error)
	View(string, map[string][]metrics.DStats) (rpi.Disk, error)
}

// Metrics represents the system metrics interface
type Metrics interface {
	DiskStats(bool) (map[string][]metrics.DStats, error)
}

// New creates a Disk application service instance.
func New(dsys DSYS, m Metrics) *Disk {
	return &Disk{dsys: dsys, m: m}
}
