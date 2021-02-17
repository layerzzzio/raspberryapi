package humanuser

import (
	"github.com/raspibuddy/rpi"
)

// Service represents all HumanUser application services.
type Service interface {
	List() ([]rpi.HumanUser, error)
}

// HumanUser represents an HumanUser application service.
type HumanUser struct {
	humsys HUMSYS
	i      Infos
}

// HUMSYS represents an HumanUser repository service.
type HUMSYS interface {
	List([]string) ([]rpi.HumanUser, error)
}

// Infos represents the infos interface
type Infos interface {
	ReadFile(string) ([]string, error)
}

// New creates a HumanUser application service instance.
func New(humsys HUMSYS, i Infos) *HumanUser {
	return &HumanUser{humsys: humsys, i: i}
}
