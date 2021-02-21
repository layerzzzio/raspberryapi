package transport

import (
	"net/http"
	"regexp"

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
	cr.POST("/waitfornetworkatboot", h.waitfornetworkatboot)
	cr.POST("/overscan", h.overscan)
}

func (h *HTTP) changehostname(ctx echo.Context) error {
	hostname := ctx.QueryParam("hostname")

	re := regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9\-]+[a-zA-Z0-9]$`)
	if !re.MatchString(hostname) || hostname == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - hostname badly formatted or null")
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

func (h *HTTP) waitfornetworkatboot(ctx echo.Context) error {
	action := ctx.QueryParam("action")
	if err := ActionCheck(action); err != nil {
		return err
	}

	result, err := h.svc.ExecuteWNB(action)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) overscan(ctx echo.Context) error {
	action := ctx.QueryParam("action")
	if err := ActionCheck(action); err != nil {
		return err
	}

	result, err := h.svc.ExecuteOV(action)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func ActionCheck(action string) error {
	re := regexp.MustCompile(`enable|disable`)
	if !re.MatchString(action) || action == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type or action type is null")
	} else {
		return nil
	}
}
