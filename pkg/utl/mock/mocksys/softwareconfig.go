package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// SoftwareConfig mock
type SoftwareConfig struct {
	ListFn func(
		map[string][]string,
	) (rpi.SoftwareConfig, error)
}

// List mock
func (in SoftwareConfig) List(
	vpnCountries map[string][]string,
) (rpi.SoftwareConfig, error) {
	return in.ListFn(
		vpnCountries,
	)
}
