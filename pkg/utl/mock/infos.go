package mock

import "github.com/raspibuddy/rpi"

// Infos mock
type Infos struct {
	ReadFileFn       func(path string) ([]string, error)
	IsFileExistsFn   func(path string) bool
	GetConfigFilesFn func() map[string]rpi.ConfigFileDetails
}

// ReadFile mock
func (i Infos) ReadFile(path string) ([]string, error) {
	return i.ReadFileFn(path)
}

// IsFileExists mock
func (i Infos) IsFileExists(path string) bool {
	return i.IsFileExistsFn(path)
}

// GetConfigFiles mock
func (i Infos) GetConfigFiles() map[string]rpi.ConfigFileDetails {
	return i.GetConfigFilesFn()
}
