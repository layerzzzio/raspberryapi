package softwareconfig

import (
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of SoftwareConfig model.
func (soc *SoftwareConfig) List() (rpi.SoftwareConfig, error) {
	vpnCountries := soc.i.VPNCountries("/etc/openvpn")
	return soc.socfsys.List(
		vpnCountries,
	)
}
