package cpu

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
)

// List returns a list of cpus
func (c *CPU) List() ([]rpi.CPU, error) {
	info, percent, vCore, err := c.csys.List()

	if err != nil || len(percent) != len(vCore) {
		return nil, echo.NewHTTPError(http.StatusAccepted, "Results were not returned as they could not be guaranteed")
	}

	var result []rpi.CPU

	for i := range info {
		data := rpi.CPU{
			ID:        int(info[i].CPU),
			Cores:     info[i].Cores,
			ModelName: info[i].ModelName,
			Mhz:       info[i].Mhz,
			Stats: rpi.CPUStats{
				Used:   percent[i],
				User:   vCore[i].User,
				System: vCore[i].System,
				Idle:   vCore[i].Idle,
				Nice:   vCore[i].Nice,
				Iowait: vCore[i].Iowait,
				Irq:    vCore[i].Irq,
			},
		}
		result = append(result, data)
	}
	return result, err
}
