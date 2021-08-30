package sys

import (
	"github.com/raspibuddy/rpi"
)

// Software represents a Software entity on the current system.
type Software struct{}

// List returns a list of Software info
func (int Software) List(
	isVNCInstalled bool,
	isOpenVPNInstalled bool,
	isUnzipInstalled bool,
	isNordVPNInstalled bool,
	isSurfSharkVPNInstalled bool,
	isIpVanishVPNInstalled bool,
	isVyprVpnVPNInstalled bool,
) (rpi.Software, error) {
	nordVPNInstalled := false
	surfSharkVPNInstalled := false
	ipVanishVPNInstalled := false
	vyprVpnVPNInstalled := false

	if isOpenVPNInstalled {
		if isNordVPNInstalled {
			nordVPNInstalled = true
		}

		if isIpVanishVPNInstalled {
			ipVanishVPNInstalled = true
		}

		if isSurfSharkVPNInstalled {
			surfSharkVPNInstalled = true
		}

		if isVyprVpnVPNInstalled {
			vyprVpnVPNInstalled = true
		}
	}

	return rpi.Software{
		IsVNCInstalled:          isVNCInstalled,
		IsOpenVPNInstalled:      isOpenVPNInstalled,
		IsUnzipInstalled:        isUnzipInstalled,
		IsNordVPNInstalled:      nordVPNInstalled,
		IsSurfSharkVPNInstalled: surfSharkVPNInstalled,
		IsIpVanishVPNInstalled:  ipVanishVPNInstalled,
		IsVyprVpnVPNInstalled:   vyprVpnVPNInstalled,
	}, nil
}
