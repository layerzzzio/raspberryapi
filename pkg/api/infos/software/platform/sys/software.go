package sys

import (
	"github.com/raspibuddy/rpi"
)

// Software represents a Software entity on the current system.
type Software struct{}

// List returns a list of Software info
func (int Software) List(
	isVNC bool,
	isOpenVPN bool,
	isUnzip bool,
	isNordVPN bool,
	isSurfSharkVPN bool,
	isIpVanishVPN bool,
	isVyprVpnVPN bool,
) (rpi.Software, error) {
	nordVPN := false
	surfSharkVPN := false
	ipVanishVPN := false
	vyprVpnVPN := false

	if isOpenVPN {
		if isNordVPN {
			nordVPN = true
		}

		if isIpVanishVPN {
			ipVanishVPN = true
		}

		if isSurfSharkVPN {
			surfSharkVPN = true
		}

		if isVyprVpnVPN {
			vyprVpnVPN = true
		}
	}

	return rpi.Software{
		IsVNC:          isVNC,
		IsOpenVPN:      isOpenVPN,
		IsUnzip:        isUnzip,
		IsNordVPN:      nordVPN,
		IsSurfSharkVPN: surfSharkVPN,
		IsIpVanishVPN:  ipVanishVPN,
		IsVyprVpnVPN:   vyprVpnVPN,
	}, nil
}
