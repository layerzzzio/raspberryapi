package system

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/shirou/gopsutil/cpu"
)

// CPU is a cpu entity on the current system
type CPU struct{}

// List is a function
func (c CPU) List() ([]cpu.InfoStat, []float64, []cpu.TimesStat, error) {
	perVCore := false
	errMetrics := echo.NewHTTPError(http.StatusInternalServerError, "Could not retrieve the CPU metrics")

	info, err := cpu.Info()
	if err != nil {
		return nil, nil, nil, errMetrics
	}

	percent, err := cpu.Percent(1, perVCore)
	if err != nil {
		return nil, nil, nil, errMetrics
	}

	vCore, err := cpu.Times(perVCore)
	if err != nil {
		return nil, nil, nil, errMetrics
	}

	return info, percent, vCore, nil
}
