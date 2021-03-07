package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// ConfigFile mock
type ConfigFile struct {
	ListFn func(map[string]rpi.ConfigFileDetails) (rpi.ConfigFile, error)
}

// List mock
func (cof *ConfigFile) List(configFiles map[string]rpi.ConfigFileDetails) (rpi.ConfigFile, error) {
	return cof.ListFn(configFiles)
}
