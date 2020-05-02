package cpu

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of CPU models.
func (c *CPU) List() ([]rpi.CPU, error) {
	info, percent, vCore, err := c.csys.List()

	if err != nil || len(percent) != len(vCore) || len(percent) != len(info) {
		return nil, echo.NewHTTPError(http.StatusAccepted, "Results were not returned as they could not be guaranteed")
	}

	var result []rpi.CPU

	for i := range info {
		fmt.Println(i)
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
			},
		}
		result = append(result, data)
	}
	return result, err
}
