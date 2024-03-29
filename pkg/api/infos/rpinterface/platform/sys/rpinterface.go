package sys

import (
	"regexp"

	"github.com/raspibuddy/rpi"
)

// RpInterface represents a RpInterface entity on the current system.
type RpInterface struct{}

// List returns a list of RpInterface info
func (int RpInterface) List(
	readLines []string,
	isStartXElf bool,
	isSSH bool,
	isSSHKeyGenerating bool,
	isVNC bool,
	isVNCInstalledCheck bool,
	isSPI bool,
	isI2C bool,
	isOneWire bool,
	isRemoteGpio bool,
	wifiInterfaces []string,
	isWpaSupCom map[string]bool,
	zoneInfo map[string]string,
) (rpi.RpInterface, error) {
	isCamera := false
	// I use a regex here to cover the below cases:
	// start_x=0 (regular)
	// start_x   = 0 (whitespace)
	// start_x=0 #random bash comment (comment)
	re := regexp.MustCompile(`^\s*start_x\s*=\s*1\s*\.*`)

	for _, v := range readLines {
		if re.MatchString(v) {
			isCamera = true
		}
	}

	isWifiInterfaces := false
	if len(wifiInterfaces) > 0 {
		isWifiInterfaces = true
	}

	return rpi.RpInterface{
		IsStartXElf:        isStartXElf,
		IsCamera:           isCamera,
		IsSSH:              isSSH,
		IsSSHKeyGenerating: isSSHKeyGenerating,
		IsVNC:              isVNC,
		IsVNCInstalled:     isVNCInstalledCheck,
		IsSPI:              isSPI,
		IsI2C:              isI2C,
		IsOneWire:          isOneWire,
		IsRemoteGpio:       isRemoteGpio,
		IsWifiInterfaces:   isWifiInterfaces,
		IsWpaSupCom:        isWpaSupCom,
		ZoneInfo:           zoneInfo,
	}, nil
}
