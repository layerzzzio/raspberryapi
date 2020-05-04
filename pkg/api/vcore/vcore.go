package vcore

import (
	"errors"
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
	percent, time, err := v.vsys.List()

	if err != nil || len(percent) != len(time) {
		return nil, echo.NewHTTPError(http.StatusAccepted, "results were not returned as they could not be guaranteed")
	}

	var result []rpi.VCore

	var vCoreID int
	for i, s := range time {
		vCoreID, err = concatID(extractNum(s.CPU, 0, 9))
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "parsing id was unsuccessful")
		}

		spec := rpi.VCore{
			ID:     vCoreID,
			Used:   percent[i],
			User:   time[i].User,
			System: time[i].System,
			Idle:   time[i].Idle,
		}
		result = append(result, spec)
	}
	return result, err
}

// View populates and returns one single CPU model.
func (v *VCore) View(id int) (*rpi.VCore, error) {
	percentTot, timeTot, err := v.vsys.List()

	if err != nil || len(percentTot) != len(timeTot) {
		return nil, echo.NewHTTPError(http.StatusAccepted, "results were not returned as they could not be guaranteed")
	}

	if id > len(percentTot) || id < 0 {
		return nil, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("id out of range : this system has  %v vcores; count starts from 0", len(percentTot)))
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
	for _, s := range timeTot {
		vCoreID, err = concatID(extractNum(s.CPU, 0, 9))
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "parsing id was unsuccessful")
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
		return -1, errors.New("invalid syntax")
	}
	return res, nil
}
