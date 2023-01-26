package vcore

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of vCore models.
func (v *VCore) List() ([]rpi.VCore, error) {
	percent, errP := v.m.CPUPercent(1, true)
	times, errT := v.m.CPUTimes(true)

	if errP != nil || errT != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the vcore metrics")
	}

	return v.vsys.List(percent, times)
}

// View populates and returns one single CPU model.
func (v *VCore) View(id int) (rpi.VCore, error) {
	percent, errP := v.m.CPUPercent(1, true)
	times, errT := v.m.CPUTimes(true)

	if errP != nil || errT != nil {
		return rpi.VCore{}, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the vcore metrics")
	}

	return v.vsys.View(id, percent, times)
}
