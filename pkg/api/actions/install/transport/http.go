package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/actions/configure/transport"
	"github.com/raspibuddy/rpi/pkg/api/actions/install"
)

// HTTP is a struct implementing a core application service.
type HTTP struct {
	svc install.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc install.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/install")
	cr.POST("/aptget", h.aptget)
	cr.POST("/nordvpn", h.nordvpn)
}

func (h *HTTP) aptget(ctx echo.Context) error {
	action := ctx.QueryParam("action")
	if err := transport.ActionCheck(action, `install|purge`); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type")
	}

	pkg := ctx.QueryParam("pkg")
	if pkg == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - pkg is null")
	}

	result, err := h.svc.ExecuteAG(action, pkg)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) nordvpn(ctx echo.Context) error {
	action := ctx.QueryParam("action")
	if err := transport.ActionCheck(action, `install|purge`); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type")
	}

	result, err := h.svc.ExecuteNV(action)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
