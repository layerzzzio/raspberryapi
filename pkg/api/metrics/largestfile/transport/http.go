package transport

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi/pkg/api/metrics/largestfile"
)

// HTTP is a struct implementing a largest files application service.
type HTTP struct {
	svc largestfile.Service
}

// NewHTTP creates new largestfile http service
func NewHTTP(svc largestfile.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/largestfiles")
	cr.GET("/:path", h.view)
}

func (h *HTTP) view(ctx echo.Context) error {
	path := strings.Replace(ctx.Param("path"), "_", "/", -1)
	result, err := h.svc.View(path)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}
