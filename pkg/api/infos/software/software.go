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
	isNordVPNInstalled := so.i.HasDirectoryAtLeastOneFile(openVpnEtcDir+"nordvpn/vpnconfigs", true)
	isSurfSharkVPNInstalled := so.i.HasDirectoryAtLeastOneFile(openVpnEtcDir+"surfshark/vpnconfigs", true)
	isIpVanishVPNInstalled := so.i.HasDirectoryAtLeastOneFile(openVpnEtcDir+"ipvanish/vpnconfigs", true)
	isVyprVpnVPNInstalled := so.i.HasDirectoryAtLeastOneFile(openVpnEtcDir+"vyprvpn/vpnconfigs", true)

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

// List populates and returns an array of Software model.
func (so *Software) View(pkg string) (rpi.Software, error) {
	isSpecificSoftwareInstalled := so.i.IsDPKGInstalled(pkg)
	return so.sofsys.View(isSpecificSoftwareInstalled)
}
