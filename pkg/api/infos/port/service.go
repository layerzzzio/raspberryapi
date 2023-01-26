package port

import (
	"github.com/raspibuddy/rpi"
)

// Service represents all Version application services.
type Service interface {
	View(int32) (rpi.Port, error)
}

// Port represents an Port application service.
type Port struct {
	psys PSYS
	i    Infos
}

// PSYS represents a Port repository service.
type PSYS interface {
	View(bool) (rpi.Port, error)
}

// Infos represents the infos interface
type Infos interface {
	IsPortListening(int32) bool
}

// New creates a Port application service instance.
func New(psys PSYS, i Infos) *Port {
	return &Port{psys: psys, i: i}
}
