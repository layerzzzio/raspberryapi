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
	List([]string, bool, bool) (rpi.Display, error)
}

// Infos represents the infos interface
type Infos interface {
	IsFileExists(string) bool
	ReadFile(string) ([]string, error)
	GetConfigFiles() map[string]rpi.ConfigFileDetails
	IsXscreenSaverInstalled() (bool, error)
}

// New creates a Display application service instance.
func New(dissys DISSYS, i Infos) *Display {
	return &Display{dissys: dissys, i: i}
}
