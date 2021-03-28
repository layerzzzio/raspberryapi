package softwareconfig

import (
	"github.com/raspibuddy/rpi"
)

// Service represents all SoftwareConfig application services.
type Service interface {
	List() (rpi.SoftwareConfig, error)
}

// SoftwareConfig represents an SoftwareConfig application service.
type SoftwareConfig struct {
	socfsys SOCFSYS
	i       Infos
}

// INTSYS represents an SoftwareConfig repository service.
type SOCFSYS interface {
	List(
		map[string][]string,
	) (rpi.SoftwareConfig, error)
}

// Infos represents the infos interface
type Infos interface {
	VPNCountries(string) map[string][]string
}

// New creates a SoftwareConfig application service instance.
func New(socfsys SOCFSYS, i Infos) *SoftwareConfig {
	return &SoftwareConfig{socfsys: socfsys, i: i}
}
