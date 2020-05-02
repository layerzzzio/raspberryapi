package transport

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi/pkg/api/cpu"
)

// HTTP is a struct implementing a core application service.
type HTTP struct {
	svc cpu.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc cpu.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/cpus")
	cr.GET("", h.list)
}

func (h *HTTP) list(ctx echo.Context) error {
	result, err := h.svc.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
