package sys

import (
	"github.com/raspibuddy/rpi"
)

// Port represents a Port entity on the current system.
type Port struct{}

// View returns a list of all api versions on the system
func (p Port) View(
	isListen bool,
) (rpi.Port, error) {
	return rpi.Port{IsSpecificPortListen: isListen}, nil
}
