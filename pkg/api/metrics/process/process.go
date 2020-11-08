package process

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of Process models.
func (p *Process) List() ([]rpi.Process, error) {
	pinfo, err := p.m.Processes()

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "could not list the process metrics")
	}

	return p.psys.List(pinfo)
}

//View populates and returns a Process model.
func (p *Process) View(id int32) (rpi.Process, error) {
	pinfo, err := p.m.Processes(id)

	if err != nil {
		if err.Error() == "process not found" {
			return rpi.Process{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("process %v does not exist", id))
		}
		return rpi.Process{}, echo.NewHTTPError(http.StatusInternalServerError, "could not view the process metrics")
	}

	return p.psys.View(id, pinfo)
}
