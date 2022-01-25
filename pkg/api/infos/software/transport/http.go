package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/infos/software"
)

// HTTP is a struct implementing a software application service.
type HTTP struct {
	svc software.Service
}

// NewHTTP creates new software http service
func NewHTTP(svc software.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/softwares")
	cr.GET("", h.list)
	cr.GET("/specifics", h.view)
}

func (h *HTTP) list(ctx echo.Context) error {
	result, err := h.svc.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) view(ctx echo.Context) error {
	pkg := ctx.QueryParam("pkg")
	if pkg == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - pkg is null")
	}

	result, err := h.svc.View(pkg)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
