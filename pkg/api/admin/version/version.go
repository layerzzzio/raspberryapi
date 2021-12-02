package version

import (
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of Version model.
func (v *Version) List() (rpi.Version, error) {
	apiVersion := v.i.ApiVersion("/usr/bin")
	return v.vsys.List(apiVersion)
}
