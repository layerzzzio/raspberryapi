package appconfig

import (
	"github.com/raspibuddy/rpi"
)

// Service represents all AppConfig application services.
type Service interface {
	ListVPN() (rpi.AppConfigVPNWithOvpn, error)
}

// AppConfig represents an VPN AppConfig application service.
type AppConfigVPNWithOvpn struct {
	apcfsys APCFVPNSYS
	i       Infos
}

// APCFSYS represents a VPN AppConfig repository service.
type APCFVPNSYS interface {
	ListVPN(
		map[string](map[string]string),
	) (rpi.AppConfigVPNWithOvpn, error)
}

// Infos represents the infos interface
type Infos interface {
	VPNCountries(string) map[string](map[string]string)
}

// New creates a VPN AppConfig application service instance.
func New(apcfsys APCFVPNSYS, i Infos) *AppConfigVPNWithOvpn {
	return &AppConfigVPNWithOvpn{apcfsys: apcfsys, i: i}
}
