package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// AppStatus mock
type AppStatus struct {
	ListFn func(
		map[string]bool,
	) (rpi.AppStatus, error)
}

// List mock
func (aps AppStatus) List(
	statusVPNWithOpenVPNFn map[string]bool,
) (rpi.AppStatus, error) {
	return aps.ListFn(
		statusVPNWithOpenVPNFn,
	)
}
