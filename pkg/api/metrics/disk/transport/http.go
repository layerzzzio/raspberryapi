package transport

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi/pkg/api/metrics/disk"
)

// HTTP is a struct implementing a core application service.
type HTTP struct {
	svc disk.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc disk.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/disks")
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
	result, err := h.svc.View(ctx.Param("id"))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}
