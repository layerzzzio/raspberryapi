package mock

import (
	"github.com/raspibuddy/rpi"
)

// Infos mock
type Infos struct {
	ReadFileFn                   func(path string) ([]string, error)
	IsFileExistsFn               func(path string) bool
	GetConfigFilesFn             func() map[string]rpi.ConfigFileDetails
	GetEnrichedConfigFilesFn     func(configFiles map[string]rpi.ConfigFileDetails) map[string]rpi.ConfigFileDetails
	IsXscreenSaverInstalledFn    func() (bool, error)
	IsQuietGrepFn                func(string, string, string) bool
	IsSSHKeyGeneratingFn         func(string) bool
	IsDPKGInstalledFn            func(string) bool
	IsSPIFn                      func(string) bool
	IsI2CFn                      func(string) bool
	IsVariableSetFn              func([]string, string, string) bool
	ListWifiInterfacesFn         func(string) []string
	IsWpaSupComFn                func() map[string]bool
	ZoneInfoFn                   func(string) map[string]string
	ListNameFilesInDirectoryFn   func(string) []string
	VPNCountriesFn               func(string) map[string](map[string]string)
	VPNConfigFileFn              func(string, string, string) []string
	ProcessesPidsFn              func(string) []string
	StatusVPNWithOpenVPNFn       func(string, string) map[string]bool
	HasDirectoryAtLeastOneFileFn func(string, bool) bool
	IsFileContainsKey1OrKey2Fn   func(string, string, string) (string, error)
	IsFileContainsUntilFn        func(string, string, string, int) (string, error)
	ApiVersionFn                 func(string) string
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

// IsVariableSet mock
func (i Infos) IsVariableSet(rawLines []string, key string, value string) bool {
	return i.IsVariableSetFn(rawLines, key, value)
}

// ListWifiInterfaces mock
func (i Infos) ListWifiInterfaces(directoryPath string) []string {
	return i.ListWifiInterfacesFn(directoryPath)
}

// IsWpaSupCom mock
func (i Infos) IsWpaSupCom() map[string]bool {
	return i.IsWpaSupComFn()
}

// ZoneInfo mock
func (i Infos) ZoneInfo(filePath string) map[string]string {
	return i.ZoneInfoFn(filePath)
}

// ListNameFilesInDirectory mock
func (i Infos) ListNameFilesInDirectory(directoryPath string) []string {
	return i.ListNameFilesInDirectoryFn(directoryPath)
}

// VPNCountries mock
func (i Infos) VPNCountries(directoryPath string) map[string](map[string]string) {
	return i.VPNCountriesFn(directoryPath)
}

// VPNConfigFiles mock
func (i Infos) VPNConfigFiles(vpnName string, vpnPath string, country string) []string {
	return i.VPNConfigFileFn(vpnName, vpnPath, country)
}

// VPNConfigFiles mock
func (i Infos) ProcessesPids(regex string) []string {
	return i.ProcessesPidsFn(regex)
}

// StatusVPNWithOpenVPN mock
func (i Infos) StatusVPNWithOpenVPN(regexPs string, regexName string) map[string]bool {
	return i.StatusVPNWithOpenVPNFn(regexPs, regexName)
}

// HasDirectoryAtLeastOneFile mock
func (i Infos) HasDirectoryAtLeastOneFile(path string, isIgnoreZip bool) bool {
	return i.HasDirectoryAtLeastOneFileFn(path, isIgnoreZip)
}

// IsFileContainsKey1OrKey2 mock
func (i Infos) IsFileContainsKey1OrKey2(filepath string, keyword1 string, keyword2 string) (string, error) {
	return i.IsFileContainsKey1OrKey2Fn(filepath, keyword1, keyword2)
}

// IsFileContainsUntil mock
func (i Infos) IsFileContainsUntil(filepath string, keyword1 string, keyword2 string, timelimit int) (string, error) {
	return i.IsFileContainsUntilFn(filepath, keyword1, keyword2, timelimit)
}

// ApiVersion mock
func (i Infos) ApiVersion(apiPath string) string {
	return i.ApiVersionFn(apiPath)
}
