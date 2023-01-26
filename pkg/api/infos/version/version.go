package version

import (
	"github.com/raspibuddy/rpi"
)

// ListAllApis populates and returns a Version model.
func (v *Version) ListAll() (rpi.Version, error) {
	raspibuddyVersion := v.i.ApiVersion("/usr/bin", "raspibuddy")
	raspibuddyDeployVersion := v.i.ApiVersion("/usr/bin", "raspibuddy_deploy")
	return v.vsys.ListAll(
		raspibuddyVersion, 
		raspibuddyDeployVersion,
	)
}

// ListAllApis populates and returns a Version model.
func (v *Version) ListAllApis() (rpi.Version, error) {
	raspibuddyVersion := v.i.ApiVersion("/usr/bin", "raspibuddy")
	raspibuddyDeployVersion := v.i.ApiVersion("/usr/bin", "raspibuddy_deploy")
	return v.vsys.ListAllApis(
		raspibuddyVersion, 
		raspibuddyDeployVersion,
	)
}
