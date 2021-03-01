package configfile

import (
	"github.com/raspibuddy/rpi"
)

// Service represents all configfile application services.
type Service interface {
	List() (rpi.ConfigFile, error)
}

// ConfigFile represents a ConfigFile application service.
type ConfigFile struct {
	cofsys COFSYS
	i      Infos
}

// COFSYS represents a ConfigFile repository service.
type COFSYS interface {
	List(map[string]rpi.ConfigFileDetails) (rpi.ConfigFile, error)
}

// Infos represents the infos interface
type Infos interface {
	GetConfigFiles() map[string]rpi.ConfigFileDetails
	GetEnrichedConfigFiles(map[string]rpi.ConfigFileDetails) map[string]rpi.ConfigFileDetails
}

// New creates a ConfigFile application service instance.
func New(cofsys COFSYS, i Infos) *ConfigFile {
	return &ConfigFile{cofsys: cofsys, i: i}
}
