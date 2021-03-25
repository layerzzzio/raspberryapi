package mocksys

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/softwareconfig"
)

// SoftwareConfig mock
type SoftwareConfig struct {
	ListFn func(
		softwareconfig.NordVPN,
	) (rpi.SoftwareConfig, error)
}

// List mock
func (in SoftwareConfig) List(
	nordVPN softwareconfig.NordVPN,
) (rpi.SoftwareConfig, error) {
	return in.ListFn(
		nordVPN,
	)
}
