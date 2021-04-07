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
	cr.POST("/blanking", h.blanking)
	cr.POST("/adduser", h.adduser)
	cr.POST("/deleteuser", h.deleteuser)
	cr.POST("/camera", h.camera)
	cr.POST("/ssh", h.ssh)
	cr.POST("/vnc", h.vnc)
	cr.POST("/spi", h.spi)
	cr.POST("/i2c", h.i2c)
	cr.POST("/onewire", h.onewire)
	cr.POST("/rgpio", h.rgpio)
	cr.POST("/update", h.update)
	cr.POST("/upgrade", h.upgrade)
	cr.POST("/updateupgrade", h.updateupgrade)
	cr.POST("/wificountry", h.wificountry)
}

func ActionCheck(action string, regex string) error {
	re := regexp.MustCompile(`^(` + regex + `)$`)
	if !re.MatchString(action) || action == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type or action type is null")
	} else {
		return nil
	}
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
	if err := ActionCheck(action, `enable|disable`); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type")
	}

	result, err := h.svc.ExecuteWNB(action)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) overscan(ctx echo.Context) error {
	action := ctx.QueryParam("action")
	if err := ActionCheck(action, `enable|disable`); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type")
	}

	result, err := h.svc.ExecuteOV(action)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) blanking(ctx echo.Context) error {
	action := ctx.QueryParam("action")
	if err := ActionCheck(action, `enable|disable`); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type")
	}

	result, err := h.svc.ExecuteBL(action)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) adduser(ctx echo.Context) error {
	username := ctx.QueryParam("username")
	if username == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - username is null")
	}

	password := ctx.QueryParam("password")
	if password == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - password is null")
	}

	result, err := h.svc.ExecuteAUS(username, password)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) deleteuser(ctx echo.Context) error {
	username := ctx.QueryParam("username")
	if username == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - username is null")
	}

	result, err := h.svc.ExecuteDUS(username)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) camera(ctx echo.Context) error {
	action := ctx.QueryParam("action")
	if err := ActionCheck(action, `enable|disable`); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type")
	}

	result, err := h.svc.ExecuteCA(action)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) ssh(ctx echo.Context) error {
	action := ctx.QueryParam("action")
	if err := ActionCheck(action, `enable|disable`); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type")
	}

	result, err := h.svc.ExecuteSSH(action)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) vnc(ctx echo.Context) error {
	action := ctx.QueryParam("action")
	if err := ActionCheck(action, `enable|disable`); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type")
	}

	result, err := h.svc.ExecuteVNC(action)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) spi(ctx echo.Context) error {
	action := ctx.QueryParam("action")
	if err := ActionCheck(action, `enable|disable`); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type")
	}

	result, err := h.svc.ExecuteSPI(action)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) i2c(ctx echo.Context) error {
	action := ctx.QueryParam("action")
	if err := ActionCheck(action, `enable|disable`); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type")
	}

	result, err := h.svc.ExecuteI2C(action)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) onewire(ctx echo.Context) error {
	action := ctx.QueryParam("action")
	if err := ActionCheck(action, `enable|disable`); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type")
	}

	result, err := h.svc.ExecuteONW(action)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) rgpio(ctx echo.Context) error {
	action := ctx.QueryParam("action")
	if err := ActionCheck(action, `enable|disable`); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type")
	}

	result, err := h.svc.ExecuteRG(action)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) update(ctx echo.Context) error {
	result, err := h.svc.ExecuteUPD()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) upgrade(ctx echo.Context) error {
	result, err := h.svc.ExecuteUPG()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) updateupgrade(ctx echo.Context) error {
	result, err := h.svc.ExecuteUPDG()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) wificountry(ctx echo.Context) error {
	iface := ctx.QueryParam("iface")
	if iface == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - iface is null")
	}

	country := ctx.QueryParam("country")
	if country == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - country is null")
	}

	result, err := h.svc.ExecuteWC(iface, country)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}
