package appconfig

import (
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of AppConfig model.
func (apc *AppConfig) List() (rpi.AppConfig, error) {
	vpnCountries := apc.i.VPNCountries("/etc/openvpn")
	return apc.apcfsys.List(
		vpnCountries,
	)
}
