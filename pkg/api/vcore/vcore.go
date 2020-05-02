package vcore

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/cpu"
)

// List populates and returns an array of vCore models.
func (v *VCore) List() ([]rpi.VCore, error) {
	percent, vCore, err := v.vsys.List()

	if err != nil || len(percent) != len(vCore) {
		return nil, echo.NewHTTPError(http.StatusAccepted, "Results were not returned as they could not be guaranteed")
	}

	var result []rpi.VCore

	var vCoreID int
	for i, s := range vCore {
		vCoreID, err = concatID(extractNum(s.CPU, 0, 9))
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "Parsing vCoreID was unsuccessful")
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

// View populates and returns one single CPU model.
func (v *VCore) View(id int) (*rpi.VCore, error) {
	percentTot, vCoreTot, err := v.vsys.List()

	if len(percentTot) != len(vCoreTot) {
		return nil, echo.NewHTTPError(http.StatusAccepted, "Results were not returned as they could not be guaranteed")
	}

	if id > len(percentTot) || id < 0 {
		return nil, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("There are only %v vCores; count starts from 0", len(percentTot)))
	}

	var percent float64
	for i, s := range percentTot {
		if id == i {
			percent = s
			break
		}
	}

	var vCoreID int
	var vCore cpu.TimesStat
	for _, s := range vCoreTot {
		vCoreID, err = concatID(extractNum(s.CPU, 0, 9))
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "Parsing vCoreID was unsuccessful")
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

	return &result, err
}
func extractNum(s string, min int, max int) []string {
	r := regexp.MustCompile("[" + strconv.Itoa(min) + "-" + strconv.Itoa(max) + "]+")
	res := r.FindAllString(s, -1)
	return res
}

func concatID(input []string) (int, error) {
	str := strings.Join(input[:], "")
	res, err := strconv.Atoi(str)
	if err != nil {
		return -1, err
	}
	return res, nil
}
