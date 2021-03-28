package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/infos/softwareconfig"
)

// HTTP is a struct implementing a softwareconfig application service.
type HTTP struct {
	svc softwareconfig.Service
}

// NewHTTP creates new softwareconfig http service
func NewHTTP(svc softwareconfig.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/softwareconfigs")
	cr.GET("/vpnwithovpn", h.list)
}

func (h *HTTP) list(ctx echo.Context) error {
	result, err := h.svc.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
