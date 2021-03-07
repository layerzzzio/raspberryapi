package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/infos/rpinterface"
)

// HTTP is a struct implementing a rpinterface application service.
type HTTP struct {
	svc rpinterface.Service
}

// NewHTTP creates new rpinterface http service
func NewHTTP(svc rpinterface.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/rpinterfaces")
	cr.GET("", h.list)
}

func (h *HTTP) list(ctx echo.Context) error {
	result, err := h.svc.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
