package transport

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi/pkg/api/mem"
)

// HTTP is a struct implementing a core application service.
type HTTP struct {
	svc mem.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc mem.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/mems")
	cr.GET("", h.list)
}

func (h *HTTP) list(ctx echo.Context) error {
	result, err := h.svc.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
