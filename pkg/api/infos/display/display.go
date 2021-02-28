package display

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of Display model.
func (d *Display) List() (rpi.Display, error) {
	bootConfigPath := d.i.GetConfigFiles()["bootconfig"].Path
	bootConfig, err := d.i.ReadFile(bootConfigPath)

	if err != nil {
		return rpi.Display{}, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the display details")
	}

	return d.dissys.List(bootConfig)
}
