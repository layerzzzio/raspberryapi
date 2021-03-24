package software

import (
	"github.com/raspibuddy/rpi"
)

type NordVPN struct {
	TCP bool
	UDP bool
}

// List populates and returns an array of Software model.
func (so *Software) List() (rpi.Software, error) {
	isVNC := so.i.IsDPKGInstalled("realvnc-vnc-server")
	isOpenVPN := so.i.IsDPKGInstalled("openvpn")
	isUnzip := so.i.IsDPKGInstalled("unzip")

	nordVPN := NordVPN{
		TCP: so.i.IsFileExists("/etc/openvpn/nordvpn/ovpn_tcp"),
		UDP: so.i.IsFileExists("/etc/openvpn/nordvpn/ovpn_udp"),
	}

	return so.sofsys.List(
		isVNC,
		isOpenVPN,
		isUnzip,
		nordVPN,
	)
}
