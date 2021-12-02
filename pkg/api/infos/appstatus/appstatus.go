package appstatus

import (
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of AppStatus model.
func (aps *AppStatus) List() (rpi.AppStatus, error) {
	regexVPNPs := `openvpn --config\s*.*--auth-user-pass`
	regexVPNName := `wov_[a-zA-Z]+`
	statusVPNWithOpenVPN := aps.i.StatusVPNWithOpenVPN(regexVPNPs, regexVPNName)

	return aps.apsfsys.List(
		statusVPNWithOpenVPN,
	)
}
