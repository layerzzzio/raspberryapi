package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/infos/humanuser"
)

// HTTP is a struct implementing a humanuser application service.
type HTTP struct {
	svc humanuser.Service
}

// NewHTTP creates new humanuser http service
func NewHTTP(svc humanuser.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/humanusers")
	cr.GET("", h.list)
}

func (h *HTTP) list(ctx echo.Context) error {
	result, err := h.svc.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
