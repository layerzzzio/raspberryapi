package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/infos/configfile"
)

// HTTP is a struct implementing a configfile application service.
type HTTP struct {
	svc configfile.Service
}

// NewHTTP creates new configfile http service
func NewHTTP(svc configfile.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/configfiles")
	cr.GET("", h.list)
}

func (h *HTTP) list(ctx echo.Context) error {
	result, err := h.svc.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
