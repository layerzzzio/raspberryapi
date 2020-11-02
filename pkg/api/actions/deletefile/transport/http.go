package transport

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi/pkg/api/actions/deletefile"
)

// HTTP is a struct implementing a core application service.
type HTTP struct {
	svc deletefile.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc deletefile.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/deletefile")
	cr.GET("/:path", h.execute)
}

func (h *HTTP) execute(ctx echo.Context) error {
	path := strings.Replace(ctx.Param("path"), "_", "/", -1)

	result, err := h.svc.Execute(path)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
