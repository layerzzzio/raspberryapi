package configfile

import (
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of Display model.
func (co *ConfigFile) List() ([]rpi.ConfigFile, error) {
	configFiles := co.i.GetConfigFiles()
	return co.cofsys.List(configFiles)
}
