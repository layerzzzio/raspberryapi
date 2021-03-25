package softwareconfig

import (
	"github.com/raspibuddy/rpi"
)

type NordVPN struct {
	TCPFiles []string
	UDPFiles []string
}

// List populates and returns an array of SoftwareConfig model.
func (soc *SoftwareConfig) List() (rpi.SoftwareConfig, error) {
	tcp := "/etc/openvpn/nordvpn/ovpn_tcp"
	udp := "/etc/openvpn/nordvpn/ovpn_udp"
	nordVPN := NordVPN{
		TCPFiles: soc.i.ListNameFilesInDirectory(tcp),
		UDPFiles: soc.i.ListNameFilesInDirectory(udp),
	}

	return soc.socfsys.List(
		nordVPN,
	)
}
