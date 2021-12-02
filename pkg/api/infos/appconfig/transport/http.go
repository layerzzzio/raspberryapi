package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/infos/appconfig"
)

// HTTP is a struct implementing a appconfig application service.
type HTTP struct {
	svc appconfig.Service
}

// NewHTTP creates new appconfig http service
func NewHTTP(svc appconfig.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/appconfigs")
	cr.GET("/vpnwithovpn", h.list)
}

func (h *HTTP) list(ctx echo.Context) error {
	result, err := h.svc.ListVPN()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
