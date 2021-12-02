package sys

import (
	"github.com/raspibuddy/rpi"
)

// Version represents a Version entity on the current system.
type Version struct{}

// List returns a list of Version info
func (v Version) List(
	ApiVersion string,
) (rpi.Version, error) {
	return rpi.Version{ApiVersion: ApiVersion}, nil
}
