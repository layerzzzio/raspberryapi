package sys

import (
	"github.com/raspibuddy/rpi"
)

// Version represents a Version entity on the current system.
type Version struct{}

// ListAll returns a list of all versions on the system
func (v Version) ListAll(
	RaspiBuddyVersion string,
	RaspiBuddyDeployVersion string,
) (rpi.Version, error) {
	return rpi.Version{
		RaspiBuddyVersion:       RaspiBuddyVersion,
		RaspiBuddyDeployVersion: RaspiBuddyDeployVersion,
	}, nil
}

// ListAllApis returns a list of all api versions on the system
func (v Version) ListAllApis(
	RaspiBuddyVersion string,
	RaspiBuddyDeployVersion string,
) (rpi.Version, error) {
	return rpi.Version{
		RaspiBuddyVersion:       RaspiBuddyVersion,
		RaspiBuddyDeployVersion: RaspiBuddyDeployVersion,
	}, nil
}
