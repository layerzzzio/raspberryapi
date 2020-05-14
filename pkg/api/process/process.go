package process

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of Process models.
func (d *Process) List() ([]rpi.ProcessSummary, error) {
	pinfo, err := d.m.Processes()

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "could not list the process metrics")
	}

	return d.psys.List(pinfo)
}

//View populates and returns a Process model.
func (d *Process) View(id int32) (rpi.ProcessDetails, error) {
	pinfo, err := d.m.Processes(id)

	if err != nil {
		return rpi.ProcessDetails{}, echo.NewHTTPError(http.StatusInternalServerError, "could not view the process metrics")
	}

	return d.psys.View(id, pinfo)
}
