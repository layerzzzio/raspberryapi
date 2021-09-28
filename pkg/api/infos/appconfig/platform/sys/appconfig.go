package sys

import (
	"github.com/raspibuddy/rpi"
)

// AppConfigVPNWithOvpn represents a VPN AppConfig entity on the current system.
type AppConfigVPNWithOvpn struct{}

// ListVPN returns a list of VPN AppConfig info
func (int AppConfigVPNWithOvpn) ListVPN(
	VPNCountries map[string]map[string]string,
) (rpi.AppConfigVPNWithOvpn, error) {
	return rpi.AppConfigVPNWithOvpn{
		VPNCountries: VPNCountries,
	}, nil
}
