package host

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of Host model.
func (h *Host) List() (rpi.Host, error) {
	info, errI := h.mt.HostInfo()
	users, errU := h.mt.Users()
	cpus, errC := h.mt.CPUInfo()
	vcores, errVC := h.mt.CPUPercent(300*time.Second, true)
	vMemPer, errV := h.mt.VirtualMemory()
	sMemPer, errS := h.mt.SwapMemory()
	temp, stdErrT, errT := h.mt.Temperature()
	rpiv, stdErrR, errR := h.mt.RaspModel()

	if errI != nil || errU != nil || errC != nil || errVC != nil || errV != nil || errS != nil || (errT != nil && stdErrT != "") || (errR != nil && stdErrR != "") {
		return rpi.Host{}, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the host metrics")
	}

	return h.hsys.List(info, users, cpus, vcores, vMemPer, sMemPer, temp, rpiv)
}
