package software

import (
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of Software model.
func (so *Software) List() (rpi.Software, error) {
	isVNC := so.i.IsDPKGInstalled("realvnc-vnc-server")
	isOpenVPN := so.i.IsDPKGInstalled("openvpn")
	isUnzip := so.i.IsDPKGInstalled("unzip")

	openVpnEtcDir := "/etc/openvpn/wov_"
	isNordVPN := so.i.IsFileExists(openVpnEtcDir + "nordvpn/")
	isSurfSharkVPN := so.i.IsFileExists(openVpnEtcDir + "surfshark/")
	isIpVanishVPN := so.i.IsFileExists(openVpnEtcDir + "ipvanish/")
	isVyprVpnVPN := so.i.IsFileExists(openVpnEtcDir + "vyprvpn/")

	return so.sofsys.List(
		isVNC,
		isOpenVPN,
		isUnzip,
		isNordVPN,
		isSurfSharkVPN,
		isIpVanishVPN,
		isVyprVpnVPN,
	)
}
