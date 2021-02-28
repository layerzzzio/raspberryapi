package sys

import (
	"github.com/raspibuddy/rpi"
)

// ConfigFile represents a ConfigFile entity on the current system.
type ConfigFile struct{}

// List returns a list of ConfigFile info
func (co ConfigFile) List(config map[string]bool) ([]rpi.ConfigFile, error) {
	return nil, nil
}
