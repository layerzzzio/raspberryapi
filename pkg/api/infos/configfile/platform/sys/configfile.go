package sys

import (
	"github.com/raspibuddy/rpi"
)

// ConfigFile represents a ConfigFile entity on the current system.
type ConfigFile struct{}

// List returns a list of ConfigFile info
func (co ConfigFile) List(config map[string]rpi.ConfigFileDetails) (rpi.ConfigFile, error) {
	isFileMissing := false
	var configFiles []rpi.ConfigFileDetails

	for _, v := range config {
		if !v.IsExist {
			isFileMissing = true
		}

		configFiles = append(configFiles, v)
	}

	return rpi.ConfigFile{
		IsFileMissing: isFileMissing,
		ConfigFiles:   configFiles,
	}, nil
}
