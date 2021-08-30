package software

import (
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of Software model.
func (so *Software) List() (rpi.Software, error) {
	isVNCInstalled := so.i.IsDPKGInstalled("realvnc-vnc-server")
	isOpenVPNInstalled := so.i.IsDPKGInstalled("openvpn")
	isUnzipInstalled := so.i.IsDPKGInstalled("unzip")

	openVpnEtcDir := "/etc/openvpn/wov_"
	isNordVPNInstalled := so.i.IsFileExists(openVpnEtcDir + "nordvpn/")
	isSurfSharkVPNInstalled := so.i.IsFileExists(openVpnEtcDir + "surfshark/")
	isIpVanishVPNInstalled := so.i.IsFileExists(openVpnEtcDir + "ipvanish/")
	isVyprVpnVPNInstalled := so.i.IsFileExists(openVpnEtcDir + "vyprvpn/")

	return so.sofsys.List(
		isVNCInstalled,
		isOpenVPNInstalled,
		isUnzipInstalled,
		isNordVPNInstalled,
		isSurfSharkVPNInstalled,
		isIpVanishVPNInstalled,
		isVyprVpnVPNInstalled,
	)
}
