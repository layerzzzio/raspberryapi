package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/infos/software"
)

// HTTP is a struct implementing a software application service.
type HTTP struct {
	svc software.Service
}

// NewHTTP creates new software http service
func NewHTTP(svc software.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/softwares")
	cr.GET("", h.list)
}

func (h *HTTP) list(ctx echo.Context) error {
	result, err := h.svc.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
