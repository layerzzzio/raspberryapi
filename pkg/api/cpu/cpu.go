package cpu

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of CPU models.
func (c *CPU) List() ([]rpi.CPU, error) {
	info, errI := c.m.Info()
	percent, errP := c.m.Percent(1, false)
	times, errT := c.m.Times(false)

	if errP != nil || errT != nil || errI != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the CPU metrics")
	}

	return c.csys.List(info, percent, times)
}
