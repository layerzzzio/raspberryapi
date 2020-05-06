package cpu

import (
	"fmt"
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
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the cpu metrics")
	}

	return c.csys.List(info, percent, times)
}

// View populates and returns one single CPU model.
func (c *CPU) View(id int) (rpi.CPU, error) {
	info, errI := c.m.Info()
	percent, errP := c.m.Percent(1, false)
	times, errT := c.m.Times(false)

	fmt.Println(info)
	fmt.Println(percent)
	fmt.Println(times)

	if errP != nil || errT != nil || errI != nil {
		return rpi.CPU{}, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the cpu metrics")
	}

	return c.csys.View(id, info, percent, times)
}
