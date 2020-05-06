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
		return nil, echo.NewHTTPError(http.StatusNotFound, "results were not returned as they could not be guaranteed")
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

// View returns a CPU stats including some info about it and the percent & time usage per workload (in USER_HZ or Jiffies)
func (c CPU) View(id int, infoTot []cpu.InfoStat, percentTot []float64, timesTot []cpu.TimesStat) (rpi.CPU, error) {
	if len(percentTot) != len(timesTot) || len(timesTot) != len(infoTot) {
		return rpi.CPU{}, echo.NewHTTPError(http.StatusNotFound, "results were not returned as they could not be guaranteed")
	}

	if id > len(percentTot) || id < 1 {
		return rpi.CPU{}, echo.NewHTTPError(http.StatusNotFound, "id out of range")
	}

	var cpuID int
	var info cpu.InfoStat
	for _, t := range infoTot {
		cpuID = int(t.CPU) + 1
		if id == cpuID {
			info = t
			break
		}
	}

	percent := percentTot[id-1]
	times := timesTot[id-1]

	result := rpi.CPU{
		ID:        cpuID,
		Cores:     info.Cores,
		ModelName: info.ModelName,
		Mhz:       info.Mhz,
		Stats: rpi.CPUStats{
			Used:   percent,
			User:   times.User,
			System: times.System,
			Idle:   times.Idle,
		}}

	return result, nil
}
