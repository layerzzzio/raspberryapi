package vcore

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of vCore models.
func (v *VCore) List() ([]rpi.VCore, error) {
	percent, errP := v.m.Percent(1, true)
	times, errT := v.m.Times(true)

	if errP != nil || errT != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the VCore metrics")
	}

	return v.vsys.List(percent, times)
}

// View populates and returns one single CPU model.
func (v *VCore) View(id int) (rpi.VCore, error) {
	percent, errP := v.m.Percent(1, true)
	times, errT := v.m.Times(true)

	if errP != nil || errT != nil {
		return rpi.VCore{}, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the VCore metrics")
	}

	return v.vsys.View(id, percent, times)
}
