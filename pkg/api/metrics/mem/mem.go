package mem

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of MEM model.
func (m *Mem) List() (rpi.Mem, error) {
	smem, errS := m.mt.SwapMemory()
	vmem, errV := m.mt.VirtualMemory()

	if errS != nil || errV != nil {
		return rpi.Mem{}, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the mem metrics")
	}

	return m.msys.List(smem, vmem)
}
