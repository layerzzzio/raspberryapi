package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/admin/version"
)

// HTTP is a struct implementing a version application service.
type HTTP struct {
	svc version.Service
}

// NewHTTP creates new version http service
func NewHTTP(svc version.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/version")
	cr.GET("", h.list)
}

func (h *HTTP) list(ctx echo.Context) error {
	result, err := h.svc.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
