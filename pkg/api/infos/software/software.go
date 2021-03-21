package software

import (
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of Software model.
func (so *Software) List() (rpi.Software, error) {
	isVNC := so.i.IsDPKGInstalled("realvnc-vnc-server")
	isOpenVPN := so.i.IsDPKGInstalled("openvpn")
	
	return so.sofsys.List(
		isVNC,
		isOpenVPN,
	)
}
