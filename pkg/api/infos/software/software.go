package software

import (
	"github.com/raspibuddy/rpi"
)

type NordVPN struct {
	IsTCP bool
	IsUDP bool
}

// List populates and returns an array of Software model.
func (so *Software) List() (rpi.Software, error) {
	isVNC := so.i.IsDPKGInstalled("realvnc-vnc-server")
	isOpenVPN := so.i.IsDPKGInstalled("openvpn")
	isUnzip := so.i.IsDPKGInstalled("unzip")

	tcp := "/etc/openvpn/nordvpn/ovpn_tcp"
	udp := "/etc/openvpn/nordvpn/ovpn_udp"

	nordVPN := NordVPN{
		IsTCP: so.i.IsFileExists(tcp),
		IsUDP: so.i.IsFileExists(udp),
	}

	return so.sofsys.List(
		isVNC,
		isOpenVPN,
		isUnzip,
		nordVPN,
	)
}
