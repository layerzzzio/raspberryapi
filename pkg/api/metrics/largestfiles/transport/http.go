package transport

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi/pkg/api/metrics/largestfiles"
)

// HTTP is a struct implementing a largest files application service.
type HTTP struct {
	svc largestfiles.Service
}

// NewHTTP creates new largestfiles http service
func NewHTTP(svc largestfiles.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/largestfiles")
	cr.GET("", h.list)
}

func (h *HTTP) list(ctx echo.Context) error {
	result, err := h.svc.List()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}
