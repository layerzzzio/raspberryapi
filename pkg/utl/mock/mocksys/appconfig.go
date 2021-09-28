package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// AppConfig mock
type AppConfig struct {
	ListFn func(
		map[string](map[string]string),
	) (rpi.AppConfig, error)
}

// List mock
func (in AppConfig) List(
	vpnCountries map[string](map[string]string),
) (rpi.AppConfig, error) {
	return in.ListFn(
		vpnCountries,
	)
}
