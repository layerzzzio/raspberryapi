package sys

import (
	"github.com/raspibuddy/rpi"
)

// AppStatus represents a AppStatus entity on the current system.
type AppStatus struct{}

// List returns a list of AppStatus info
func (int AppStatus) List(
	statusVPNWithOpenVPN map[string]bool,
) (rpi.AppStatus, error) {
	return rpi.AppStatus{
		VPNwithOpenVPN: statusVPNWithOpenVPN,
	}, nil
}
