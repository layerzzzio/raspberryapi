package sys

import (
	"github.com/raspibuddy/rpi"
)

// SoftwareConfig represents a SoftwareConfig entity on the current system.
type SoftwareConfig struct{}

// List returns a list of SoftwareConfig info
func (int SoftwareConfig) List(
	VPNCountries map[string][]string,
) (rpi.SoftwareConfig, error) {
	return rpi.SoftwareConfig{
		VPNCountries: VPNCountries,
	}, nil
}
