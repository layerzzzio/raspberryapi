package display

import (
	"github.com/raspibuddy/rpi"
)

// Service represents all display application services.
type Service interface {
	List() (rpi.Display, error)
}

// Display represents a Display application service.
type Display struct {
	dissys DISSYS
	i      Infos
}

// DISSYS represents a Display repository service.
type DISSYS interface {
	List([]string) (rpi.Display, error)
}

// Infos represents the infos interface
type Infos interface {
	ReadFile(string) ([]string, error)
}

// New creates a Display application service instance.
func New(boosys DISSYS, i Infos) *Display {
	return &Display{dissys: boosys, i: i}
}
