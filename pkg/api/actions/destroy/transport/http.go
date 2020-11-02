package transport

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi/pkg/api/actions/destroy"
)

// HTTP is a struct implementing a core application service.
type HTTP struct {
	svc destroy.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc destroy.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/destroy")
	cr.GET("/deletefile/:path", h.deletefile)
}

func (h *HTTP) deletefile(ctx echo.Context) error {
	path := strings.Replace(ctx.Param("path"), "_", "/", -1)

	result, err := h.svc.ExecuteDF(path)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
