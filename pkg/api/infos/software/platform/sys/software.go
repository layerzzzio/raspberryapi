package sys

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/software"
)

// Software represents a Software entity on the current system.
type Software struct{}

// List returns a list of Software info
func (int Software) List(
	isVNC bool,
	isOpenVPN bool,
	isUnzip bool,
	nordVPN software.NordVPN,
) (rpi.Software, error) {
	isNordVPN := false
	if isOpenVPN && nordVPN.TCP && nordVPN.UDP {
		isNordVPN = true
	}

	return rpi.Software{
		IsVNC:     isVNC,
		IsOpenVPN: isOpenVPN,
		IsUnzip:   isUnzip,
		IsNordVpn: isNordVPN,
	}, nil
}
