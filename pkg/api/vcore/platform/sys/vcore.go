package sys

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/cpu"
)

// VCore is a vCore entity on the current system
type VCore struct{}

// List returns a list of vCore stats per CPU including percent & time usage per workload (in USER_HZ or Jiffies)
// TODO: to find a way to know which unit is used between USER_HZ and Jiffies
func (v VCore) List(percent []float64, times []cpu.TimesStat) ([]rpi.VCore, error) {
	if len(percent) != len(times) {
		return nil, echo.NewHTTPError(http.StatusAccepted, "results were not returned as they could not be guaranteed")
	}

	var result []rpi.VCore
	for i, s := range times {
		vCoreID, err := concatID(extractNum(s.CPU, 0, 9))
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "parsing id was unsuccessful")
		}

		spec := rpi.VCore{
			ID:     vCoreID + 1,
			Used:   percent[i],
			User:   times[i].User,
			System: times[i].System,
			Idle:   times[i].Idle,
		}
		result = append(result, spec)
	}

	return result, nil
}

// View returns a vCore stats including percent & time usage per workload (in USER_HZ or Jiffies)
// TODO: i.e. than List() here above
func (v VCore) View(id int, percentTot []float64, timesTot []cpu.TimesStat) (rpi.VCore, error) {
	if len(percentTot) != len(timesTot) {
		return rpi.VCore{}, echo.NewHTTPError(http.StatusAccepted, "results were not returned as they could not be guaranteed")
	}

	if id > len(percentTot) || id < 1 {
		return rpi.VCore{}, echo.NewHTTPError(http.StatusInternalServerError, "id out of range")
	}

	var percent float64
	for i, s := range percentTot {
		if id == i+1 {
			percent = s
			break
		}
	}

	var vCoreID int
	var err error
	var vCore cpu.TimesStat
	for _, s := range timesTot {
		vCoreID, err = concatID(extractNum(s.CPU, 0, 9))
		if err != nil {
			return rpi.VCore{}, echo.NewHTTPError(http.StatusInternalServerError, "parsing id was unsuccessful")
		}

		if id == vCoreID+1 {
			vCore = s
			break
		}
	}

	result := rpi.VCore{
		ID:     vCoreID + 1,
		Used:   percent,
		User:   vCore.User,
		System: vCore.System,
		Idle:   vCore.Idle,
	}

	return result, nil
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
