package sys

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/shirou/gopsutil/cpu"
)

// VCore is a vCore entity on the current system
type VCore struct{}

// List returns two lists of vCore stats: percent and time usage per workload (in USER_HZ or Jiffies)
// TODO: to find a way to know which unit is used between USER_HZ and Jiffies
func (c VCore) List() ([]float64, []cpu.TimesStat, error) {
	percent, errP := cpu.Percent(1, true)

	time, errT := cpu.Times(true)
	if errP != nil || errT != nil {
		return nil, nil, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the CPU metrics")
	}

	return percent, time, nil
}
