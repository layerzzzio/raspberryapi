package process

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
)

// Service represents all process application services.
type Service interface {
	List() ([]rpi.ProcessSummary, error)
	View(int32) (rpi.ProcessDetails, error)
}

// Process represents a process application service.
type Process struct {
	psys PSYS
	m    Metrics
}

// PSYS represents a process repository service.
type PSYS interface {
	List([]metrics.PInfo) ([]rpi.ProcessSummary, error)
	View(int32, []metrics.PInfo) (rpi.ProcessDetails, error)
}

// Metrics represents the system metrics interface
type Metrics interface {
	Processes(id ...int32) ([]metrics.PInfo, error)
}

// New creates a Process application service instance.
func New(psys PSYS, m Metrics) *Process {
	return &Process{psys: psys, m: m}
}
