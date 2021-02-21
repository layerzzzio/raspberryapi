package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/infos/boot"
)

// HTTP is a struct implementing a boot application service.
type HTTP struct {
	svc boot.Service
}

// NewHTTP creates new boot http service
func NewHTTP(svc boot.Service, r *echo.Group) {
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
