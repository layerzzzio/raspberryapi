package sys

import (
	"github.com/raspibuddy/rpi"
)

// Boot represents a Boot entity on the current system.
type Boot struct{}

// List returns a list of Boot info
func (b Boot) List(isWaitForNetwork bool) (rpi.Boot, error) {
	return rpi.Boot{
		IsWaitForNetwork: isWaitForNetwork,
	}, nil
}
