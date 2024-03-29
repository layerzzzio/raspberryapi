package transport

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/actions/general/transport"
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
	cr.POST("/deploy-api", h.deployVersion)
	// cr.GET("/purge", h.purge)

}

func (h *HTTP) deployVersion(ctx echo.Context) error {
	// deployType
	deployType := ctx.QueryParam("deployType")
	// only full_deploy & partial_deploy are authorized for RPI
	// because we don't want RPI to deploy itself
	if err := transport.ActionCheck(deployType, `full_deploy|partial_deploy`); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Not found - bad action deployType: only full_deploy & partial_deploy authorized")
	}

	// URL
	url := ctx.QueryParam("url")
	if url == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Not found - url is null")
	}

	// VERSION
	version := ctx.QueryParam("version")
	if version == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Not found - version is null")
	}

	matched, _ := regexp.MatchString(infos.ApiVersionRegex, version)
	if !matched {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request due to an invalid version format")
	}

	result, err := h.svc.ExecuteDPTOOL(deployType, url, version)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
