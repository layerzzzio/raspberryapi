package boot

import (
	"github.com/raspibuddy/rpi"
)

// Service represents all boot application services.
type Service interface {
	List() (rpi.Boot, error)
}

// Boot represents a Boot application service.
type Boot struct {
	boosys BOOSYS
	i      Infos
}

// BOOSYS represents a Boot repository service.
type BOOSYS interface {
	List(bool) (rpi.Boot, error)
}

// Infos represents the infos interface
type Infos interface {
	IsFileExists(string) bool
}

// New creates a Boot application service instance.
func New(boosys BOOSYS, i Infos) *Boot {
	return &Boot{boosys: boosys, i: i}
}
