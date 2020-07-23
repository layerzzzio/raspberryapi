package transport

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi/pkg/api/load"
)

// HTTP is a struct implementing a core application service.
type HTTP struct {
	svc load.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc load.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/loads")
	cr.GET("", h.list)
}

func (h *HTTP) list(ctx echo.Context) error {
	result, err := h.svc.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
