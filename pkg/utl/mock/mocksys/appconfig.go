package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// AppConfig mock
type AppConfigVPNWithOvpn struct {
	ListVPNFn func(
		map[string](map[string]string),
	) (rpi.AppConfigVPNWithOvpn, error)
}

// List mock
func (in AppConfigVPNWithOvpn) ListVPN(
	vpnCountries map[string](map[string]string),
) (rpi.AppConfigVPNWithOvpn, error) {
	return in.ListVPNFn(
		vpnCountries,
	)
}
