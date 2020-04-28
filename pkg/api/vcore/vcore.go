package vcore

import (
	"regexp"
	"strconv"

	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/cpu"
)

// List returns a list of cores (virtual cores)
func (c *VCore) List() ([]rpi.VCore, error) {
	percent, vCore, err := c.vsys.List()

	if len(percent) != len(vCore) {
		panic(err)
	}

	var result []rpi.VCore

	for i, s := range vCore {
		vCoreID, err := extractNum(s.CPU, 0, 9)
		if err != nil {
			error.Error(err)
		}

		spec := rpi.VCore{
			ID:     vCoreID,
			Used:   percent[i],
			User:   vCore[i].User,
			System: vCore[i].System,
			Idle:   vCore[i].Idle,
			Nice:   vCore[i].Nice,
			Iowait: vCore[i].Iowait,
			Irq:    vCore[i].Irq,
		}
		result = append(result, spec)
	}
	return result, err
}

// View returns a list of cores (virtual cores)
func (c *VCore) View(id int) (rpi.VCore, error) {
	percentTot, vCoreTot, err := c.vsys.List()

	if len(percentTot) != len(vCoreTot) {
		panic(err)
	}

	var percent float64
	for i, s := range percentTot {
		if id == i {
			percent = s
			break
		} else {
			percent = -1
		}
	}

	var vCoreID int
	var vCore cpu.TimesStat
	for _, s := range vCoreTot {
		vCoreID, err = extractNum(s.CPU, 0, 9)
		if err != nil {
			error.Error(err)
		}

		if id == vCoreID {
			vCore = s
			break
		}
	}

	result := rpi.VCore{
		ID:     vCoreID,
		Used:   percent,
		User:   vCore.User,
		System: vCore.System,
		Idle:   vCore.Idle,
		Nice:   vCore.Nice,
		Iowait: vCore.Iowait,
		Irq:    vCore.Irq,
	}

	return result, err
}

func extractNum(s string, min int, max int) (int, error) {
	r := regexp.MustCompile("[" + strconv.Itoa(min) + "-" + strconv.Itoa(max) + "]+")
	num := r.FindString(s)
	res, err := strconv.Atoi(num)
	if err != nil {
		return -1, err
	}
	return res, nil
}
