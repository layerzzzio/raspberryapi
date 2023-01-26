package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/infos/version"
)

// HTTP is a struct implementing a version application service.
type HTTP struct {
	svc version.Service
}

// NewHTTP creates new version http service
func NewHTTP(svc version.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/versions")
	cr.GET("", h.listAllVersions)
	cr.GET("/apis", h.listAllApisVersions)
}

func (h *HTTP) listAllVersions(ctx echo.Context) error {
	result, err := h.svc.ListAll()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) listAllApisVersions(ctx echo.Context) error {
	result, err := h.svc.ListAllApis()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
