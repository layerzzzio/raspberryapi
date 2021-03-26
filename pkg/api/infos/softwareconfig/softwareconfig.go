package softwareconfig

import (
	"github.com/raspibuddy/rpi"
)

type NordVPN struct {
	TCPCountries []string
	UDPCountries []string
}

// List populates and returns an array of SoftwareConfig model.
func (soc *SoftwareConfig) List() (rpi.SoftwareConfig, error) {
	tcp := "/etc/openvpn/nordvpn/ovpn_tcp"
	udp := "/etc/openvpn/nordvpn/ovpn_udp"
	nordVPN := NordVPN{
		TCPCountries: soc.i.VPNCountries(tcp),
		UDPCountries: soc.i.VPNCountries(udp),
	}

	return soc.socfsys.List(
		nordVPN,
	)
}
