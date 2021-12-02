package version

import (
	"github.com/raspibuddy/rpi"
)

// Service represents all Version application services.
type Service interface {
	List() (rpi.Version, error)
}

// Version represents an Version application service.
type Version struct {
	vsys VSYS
	i    Infos
}

// VSYS represents a Version repository service.
type VSYS interface {
	List(string) (rpi.Version, error)
}

// Infos represents the infos interface
type Infos interface {
	ApiVersion(string) string
}

// New creates a Version application service instance.
func New(vsys VSYS, i Infos) *Version {
	return &Version{vsys: vsys, i: i}
}
