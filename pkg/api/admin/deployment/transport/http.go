package transport

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/admin/deployment"
	"github.com/raspibuddy/rpi/pkg/utl/infos"
)

// HTTP is a struct implementing a deployment application service.
type HTTP struct {
	svc deployment.Service
}

// NewHTTP creates new deployment http service
func NewHTTP(svc deployment.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/deploy")
	cr.GET("/:version", h.deployVersion)
}

func (h *HTTP) deployVersion(ctx echo.Context) error {
	version := ctx.Param("version")
	matched, _ := regexp.MatchString(infos.ApiVersionRegex, version)
	if !matched {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request due to an invalid process version format")
	}

	result, err := h.svc.ExecuteDP(version)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
