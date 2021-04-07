package sys

import (
	"github.com/raspibuddy/rpi"
)

// AppConfig represents a AppConfig entity on the current system.
type AppConfig struct{}

// List returns a list of AppConfig info
func (int AppConfig) List(
	VPNCountries map[string][]string,
) (rpi.AppConfig, error) {
	return rpi.AppConfig{
		VPNCountries: VPNCountries,
	}, nil
}
