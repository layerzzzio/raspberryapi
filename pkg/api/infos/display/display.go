package display

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/infos"
)

// List populates and returns an array of Display model.
func (d *Display) List() (rpi.Display, error) {
	bootConfig, err := d.i.ReadFile(infos.BootConfig)

	if err != nil {
		return rpi.Display{}, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the display details")
	}

	return d.dissys.List(bootConfig)
}
