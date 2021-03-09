package mock

import (
	"github.com/raspibuddy/rpi"
)

// Infos mock
type Infos struct {
	ReadFileFn                func(path string) ([]string, error)
	IsFileExistsFn            func(path string) bool
	GetConfigFilesFn          func() map[string]rpi.ConfigFileDetails
	GetEnrichedConfigFilesFn  func(configFiles map[string]rpi.ConfigFileDetails) map[string]rpi.ConfigFileDetails
	IsXscreenSaverInstalledFn func() (bool, error)
	IsQuietGrepFn             func(string, string, string) bool
	IsSSHKeyGeneratingFn      func(string) bool
	IsDPKGInstalledFn         func(string) bool
	IsSPIFn                   func(string) bool
	IsI2CFn                   func(string) bool
	IsVariableSetFn           func([]string, string, string) bool
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

// GetEnrichedConfigFiles mock
func (i Infos) GetEnrichedConfigFiles(configFiles map[string]rpi.ConfigFileDetails) map[string]rpi.ConfigFileDetails {
	return i.GetEnrichedConfigFilesFn(configFiles)
}

// IsXscreenSaverInstalled mock
func (i Infos) IsXscreenSaverInstalled() (bool, error) {
	return i.IsXscreenSaverInstalledFn()
}

// IsQuietGrep mock
func (i Infos) IsQuietGrep(command string, quietGrep string, grepType string) bool {
	return i.IsQuietGrepFn(command, quietGrep, grepType)
}

// IsSSHKeyGenerating mock
func (i Infos) IsSSHKeyGenerating(path string) bool {
	return i.IsSSHKeyGeneratingFn(path)
}

// IsDPKGInstalled mock
func (i Infos) IsDPKGInstalled(path string) bool {
	return i.IsDPKGInstalledFn(path)
}

// IsSPI mock
func (i Infos) IsSPI(path string) bool {
	return i.IsSPIFn(path)
}

// IsI2C mock
func (i Infos) IsI2C(path string) bool {
	return i.IsI2CFn(path)
}

// IsVariableSetFn mock
func (i Infos) IsVariableSet(rawLines []string, key string, value string) bool {
	return i.IsVariableSetFn(rawLines, key, value)
}
