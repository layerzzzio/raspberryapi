package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/infos/display"
)

// HTTP is a struct implementing a display application service.
type HTTP struct {
	svc display.Service
}

// NewHTTP creates new display http service
func NewHTTP(svc display.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/displays")
	cr.GET("", h.list)
}

func (h *HTTP) list(ctx echo.Context) error {
	result, err := h.svc.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
