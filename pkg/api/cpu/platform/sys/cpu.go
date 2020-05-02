package sys

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/shirou/gopsutil/cpu"
)

// CPU represents an empty CPU entity on the current system.
type CPU struct{}

// List retrieves the current system CPU statistics.
// TODO: to test this method by simulating different OS scenarios in a Docker container (raspbian/strech)
func (c CPU) List() ([]cpu.InfoStat, []float64, []cpu.TimesStat, error) {
	info, errI := cpu.Info()
	percent, errP := cpu.Percent(1, false)
	time, errT := cpu.Times(false)

	if errI != nil || errP != nil || errT != nil {
		return nil, nil, nil, echo.NewHTTPError(http.StatusInternalServerError, "Could not retrieve the CPU metrics")
	}

	return info, percent, time, nil
}
