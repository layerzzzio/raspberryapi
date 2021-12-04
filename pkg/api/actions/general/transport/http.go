package transport

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/actions/general"
)

// HTTP is a struct implementing a core application service.
type HTTP struct {
	svc general.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc general.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/general")
	cr.POST("/boot", h.rebootshutdown)
	cr.POST("/startstop", h.startStop)
}

func (h *HTTP) rebootshutdown(ctx echo.Context) error {
	option := ctx.QueryParam("option")

	if err := BootOptionCheck(option, `reboot|shutdown`); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - bad option type")
	}

	result, err := h.svc.ExecuteRBS(option)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) startStop(ctx echo.Context) error {
	action := ctx.QueryParam("action")
	if err := ActionCheck(action, `start|stop`); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad action type")
	}

	service := ctx.QueryParam("service")
	if service == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - please enter a service")
	}

	result, err := h.svc.ExecuteSASO(action, service)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

func BootOptionCheck(action string, regex string) error {
	re := regexp.MustCompile(`^(` + regex + `)$`)
	if !re.MatchString(action) || action == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - bad option type")
	} else {
		return nil
	}
}

func ActionCheck(action string, regex string) error {
	re := regexp.MustCompile(`^(` + regex + `)$`)
	if !re.MatchString(action) || action == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type or action type is null")
	} else {
		return nil
	}
}
