package transport

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/infos/port"
)

// HTTP is a struct implementing a port application service.
type HTTP struct {
	svc port.Service
}

// NewHTTP creates new port http service
func NewHTTP(svc port.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/ports")
	cr.GET("/:port", h.view)
}

func (h *HTTP) view(ctx echo.Context) error {
	port, err := strconv.Atoi(ctx.Param("port"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request due to an invalid port - should be an integer")
	}

	result, err := h.svc.View(int32(port))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}
