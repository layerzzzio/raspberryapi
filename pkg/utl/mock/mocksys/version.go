package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// Version mock
type Version struct {
	ListAllFn     func(string, string) (rpi.Version, error)
	ListAllApisFn func(string, string) (rpi.Version, error)
}

// ListAll mock
func (v Version) ListAll(
	raspibuddyVersion string,
	raspibuddyDeployVersion string,
) (rpi.Version, error) {
	return v.ListAllFn(
		raspibuddyVersion,
		raspibuddyDeployVersion,
	)
}

// ListAll mock
func (v Version) ListAllApis(
	raspibuddyVersion string,
	raspibuddyDeployVersion string,
) (rpi.Version, error) {
	return v.ListAllApisFn(
		raspibuddyVersion,
		raspibuddyDeployVersion,
	)
}
