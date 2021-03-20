package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// RpInterface mock
type RpInterface struct {
	ListFn func([]string, bool, bool, bool, bool, bool, bool, bool, bool, bool, []string, map[string]string) (rpi.RpInterface, error)
}

// List mock
func (in RpInterface) List(
	readLines []string,
	isStartXElf bool,
	isSSH bool,
	isSSHKeyGenerating bool,
	isVNC bool,
	isVNCInstalledCheck bool,
	isSPI bool,
	isI2C bool,
	isVariableSet bool,
	isRemoteGpio bool,
	wifiInterfaces []string,
	zoneInfo map[string]string,
) (rpi.RpInterface, error) {
	return in.ListFn(
		readLines,
		isStartXElf,
		isSSH,
		isSSHKeyGenerating,
		isVNC,
		isVNCInstalledCheck,
		isSPI,
		isI2C,
		isVariableSet,
		isRemoteGpio,
		wifiInterfaces,
		zoneInfo,
	)
}
