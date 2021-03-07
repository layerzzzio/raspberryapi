package configfile

import (
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of Display model.
func (co *ConfigFile) List() (rpi.ConfigFile, error) {
	configFiles := co.i.GetConfigFiles()
	enrichedConfigFiles := co.i.GetEnrichedConfigFiles(configFiles)
	return co.cofsys.List(enrichedConfigFiles)
}
