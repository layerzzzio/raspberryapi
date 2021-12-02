package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// Version mock
type Version struct {
	ListFn func(
		string,
	) (rpi.Version, error)
}

// List mock
func (v Version) List(
	version string,
) (rpi.Version, error) {
	return v.ListFn(
		version,
	)
}
