package transport

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/metrics/net"
)

// HTTP is a struct implementing a core application service.
type HTTP struct {
	svc net.Service
}

// NewHTTP creates new net http service
func NewHTTP(svc net.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/nets")
	cr.GET("", h.list)
	cr.GET("/:id", h.view)
}

func (h *HTTP) list(ctx echo.Context) error {
	result, err := h.svc.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) view(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request due to an invalid net id - should be an integer")
	}

	result, err := h.svc.View(id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}
