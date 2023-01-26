package host

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of Host model.
func (h *Host) List() (rpi.Host, error) {
	info, errI := h.mt.HostInfo()
	users, errU := h.mt.Users()
	cpus, errC := h.mt.CPUInfo()
	vcores, errVC := h.mt.CPUPercent(5000000, true)
	vMem, errV := h.mt.VirtualMemory()
	sMemPer, errS := h.mt.SwapMemory()
	temp, stdErrT, errT := h.mt.Temperature()
	serialNumber, stdErrSN, errSN := h.mt.SerialNumber()
	rpiv, stdErrR, errR := h.mt.RaspModel()
	load, errL := h.mt.LoadAvg()
	listDev, errD := h.mt.DiskStats(false)
	netInfo, errNI := h.mt.NetInfo()

	if errNI != nil || errD != nil || errL != nil || errI != nil || errU != nil || errC != nil || errVC != nil || errV != nil || errS != nil || (errSN != nil && stdErrSN != "") || (errT != nil && stdErrT != "") || (errR != nil && stdErrR != "") {
		return rpi.Host{}, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the host metrics")
	}

	return h.hsys.List(info, users, cpus, vcores, vMem, sMemPer, load, temp, serialNumber, rpiv, listDev, netInfo)
}
