package load

import (
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/load"
)

// Service represents all Load application services.
type Service interface {
	List() (rpi.Load, error)
}

// Load represents a Load application service.
type Load struct {
	lsys LSYS
	m    Metrics
}

// LSYS represents a Load repository service.
type LSYS interface {
	List(load.AvgStat, load.MiscStat) (rpi.Load, error)
}

// Metrics represents the system metrics interface
type Metrics interface {
	LoadAvg() (load.AvgStat, error)
	LoadProcs() (load.MiscStat, error)
}

// New creates a Load application service instance.
func New(lsys LSYS, m Metrics) *Load {
	return &Load{lsys: lsys, m: m}
}
