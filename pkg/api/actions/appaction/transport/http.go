package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/actions/appaction"
	"github.com/raspibuddy/rpi/pkg/api/actions/configure/transport"
)

// HTTP is a struct implementing a core application service.
type HTTP struct {
	svc appaction.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc appaction.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/appaction")
	cr.POST("/vpnwithovpn", h.vpnwithopenvpn)
}

func (h *HTTP) vpnwithopenvpn(ctx echo.Context) error {
	username := ""
	password := ""
	country := ""

	action := ctx.QueryParam("action")
	if err := transport.ActionCheck(action, `connect|disconnect`); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type")
	}

	vpnName := ctx.QueryParam("vpnName")
	if vpnName == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - vpnName is nil")
	}

	if action == "connect" {
		country = ctx.QueryParam("country")
		// if country == "" {
		// 	return echo.NewHTTPError(http.StatusNotFound, "Not found - country is nil")
		// }

		username = ctx.QueryParam("username")
		// if username == "" {
		// 	return echo.NewHTTPError(http.StatusNotFound, "Not found - username is nil")
		// }

		password = ctx.QueryParam("password")
		// if password == "" {
		// 	return echo.NewHTTPError(http.StatusNotFound, "Not found - password is nil")
		// }
	}

	result, err := h.svc.ExecuteWOVA(action, vpnName, country, username, password)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
