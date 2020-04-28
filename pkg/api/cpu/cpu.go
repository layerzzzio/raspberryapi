package cpu

import "github.com/raspibuddy/rpi"

// List returns a list of cpus
func (c *CPU) List() ([]rpi.CPU, error) {
	info, percent, vCore, err := c.csys.List()

	if len(info) != len(percent) && len(percent) != len(vCore) {
		panic(err)
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
