package rpinterface

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of RpInterface model.
func (in *RpInterface) List() (rpi.RpInterface, error) {
	bootConfigPath := in.i.GetConfigFiles()["bootconfig"].Path
	bootConfig, errB := in.i.ReadFile(bootConfigPath)

	startXElf := in.i.GetConfigFiles()["start_x_elf"].Path
	isStartXElf := in.i.IsFileExists(startXElf)

	// keep inactive as a keywork here
	isSSH := in.i.IsQuietGrep("service ssh status", "inactive", "quiet")
	isSSHKeyGenerating := in.i.IsSSHKeyGenerating("/var/log/regen_ssh_keys.log")

	isVNC := in.i.IsQuietGrep("systemctl status vncserver-x11-serviced.service", "active", "word-regexp")
	isVNCInstalledCheck := in.i.IsDPKGInstalled("realvnc-vnc-server")

	isSPI := in.i.IsSPI(bootConfigPath)
	isI2C := in.i.IsI2C(bootConfigPath)
	isOneWire := in.i.IsVariableSet(bootConfig, "dtoverlay", "w1-gpio")

	if errB != nil {
		return rpi.RpInterface{}, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the rpinterface details")
	}

	return in.intsys.List(
		bootConfig,
		isStartXElf,
		!isSSH,
		isSSHKeyGenerating,
		isVNC,
		isVNCInstalledCheck,
		isSPI,
		isI2C,
		isOneWire,
	)
}
