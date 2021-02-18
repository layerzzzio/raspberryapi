package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/actions/configure"
)

// HTTP is a struct implementing a core application service.
type HTTP struct {
	svc configure.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc configure.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/configure")
	cr.POST("/changehostname", h.changehostname)
	cr.POST("/changepassword", h.changepassword)

}

func (h *HTTP) changehostname(ctx echo.Context) error {
	hostname := ctx.QueryParam("hostname")
	if hostname == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - hostname is null")
	}

	result, err := h.svc.ExecuteCH(hostname)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) changepassword(ctx echo.Context) error {
	password := ctx.QueryParam("password")
	if password == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - password is null")
	}

	username := ctx.QueryParam("username")
	if username == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - username is null")
	}

	result, err := h.svc.ExecuteCP(password, username)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
