package load

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of Load model.
func (l *Load) List() (rpi.Load, error) {
	temp, errT := l.m.LoadAvg()
	procs, errP := l.m.LoadProcs()

	if errT != nil || errP != nil {
		return rpi.Load{}, echo.NewHTTPError(http.StatusInternalServerError, "could not list the load metrics")
	}

	return l.lsys.List(temp, procs)
}
