package sys

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/softwareconfig"
)

// SoftwareConfig represents a SoftwareConfig entity on the current system.
type SoftwareConfig struct{}

// List returns a list of SoftwareConfig info
func (int SoftwareConfig) List(
	nordVPN softwareconfig.NordVPN,
) (rpi.SoftwareConfig, error) {
	return rpi.SoftwareConfig{
		NordVPN: rpi.NordVPN{
			TCPFiles: nordVPN.TCPFiles,
			UDPFiles: nordVPN.UDPFiles,
		},
	}, nil
}
