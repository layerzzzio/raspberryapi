package sys

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/cpu"
)

// CPU represents an empty CPU entity on the current system.
type CPU struct{}

// List returns a list of CPU stats including some info about it and the percent & time usage per workload (in USER_HZ or Jiffies)
func (c CPU) List(info []cpu.InfoStat, percent []float64, times []cpu.TimesStat) ([]rpi.CPU, error) {
	if len(percent) != len(times) || len(times) != len(info) {
		return nil, echo.NewHTTPError(http.StatusAccepted, "results were not returned as they could not be guaranteed")
	}

	var result []rpi.CPU
	for i := range info {
		data := rpi.CPU{
			ID:        int(info[i].CPU) + 1,
			Cores:     info[i].Cores,
			ModelName: info[i].ModelName,
			Mhz:       info[i].Mhz,
			Stats: rpi.CPUStats{
				Used:   percent[i],
				User:   times[i].User,
				System: times[i].System,
				Idle:   times[i].Idle,
			},
		}
		result = append(result, data)
	}
	return result, nil
}
