package appconfig

import (
	"github.com/raspibuddy/rpi"
)

// Service represents all AppConfig application services.
type Service interface {
	List() (rpi.AppConfig, error)
}

// AppConfig represents an AppConfig application service.
type AppConfig struct {
	apcfsys APCFSYS
	i       Infos
}

// APCFSYS represents an AppConfig repository service.
type APCFSYS interface {
	List(
		map[string][]string,
	) (rpi.AppConfig, error)
}

// Infos represents the infos interface
type Infos interface {
	VPNCountries(string) map[string][]string
}

// New creates a AppConfig application service instance.
func New(apcfsys APCFSYS, i Infos) *AppConfig {
	return &AppConfig{apcfsys: apcfsys, i: i}
}
