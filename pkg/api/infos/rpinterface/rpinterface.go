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

	isSSH := in.i.IsQuietGrep("service ssh status", "active")

	if errB != nil {
		return rpi.RpInterface{}, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the rpinterface details")

	}

	return in.intsys.List(bootConfig, isStartXElf, isSSH)
}
