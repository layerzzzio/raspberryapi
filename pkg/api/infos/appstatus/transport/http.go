package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/infos/appstatus"
)

// HTTP is a struct implementing a appstatus application service.
type HTTP struct {
	svc appstatus.Service
}

// NewHTTP creates new appstatus http service
func NewHTTP(svc appstatus.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/appstatuses")
	cr.GET("/vpnwithovpn", h.list)
}

func (h *HTTP) list(ctx echo.Context) error {
	result, err := h.svc.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
