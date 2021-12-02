package appconfig

import (
	"github.com/raspibuddy/rpi"
)

// ListVPN populates and returns an array of VPN AppConfig model.
func (apc *AppConfigVPNWithOvpn) ListVPN() (rpi.AppConfigVPNWithOvpn, error) {
	vpnCountries := apc.i.VPNCountries("/etc/openvpn")
	return apc.apcfsys.ListVPN(
		vpnCountries,
	)
}
