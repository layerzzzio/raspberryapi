package display

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of Display model.
func (d *Display) List() (rpi.Display, error) {
	bootConfigPath := d.i.GetConfigFiles()["bootconfig"].Path
	blankingPath := d.i.GetConfigFiles()["blanking"].Path

	bootConfig, errB := d.i.ReadFile(bootConfigPath)
	isXscreenSaverInstalled, errI := d.i.IsXscreenSaverInstalled()
	isBlanking := d.i.IsFileExists(blankingPath)

	if errI != nil || errB != nil {
		return rpi.Display{}, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the display details")
	}

	return d.dissys.List(bootConfig, isXscreenSaverInstalled, isBlanking)
}
