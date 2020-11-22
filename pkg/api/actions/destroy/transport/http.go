package transport

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/actions/destroy"
)

// HTTP is a struct implementing a core application service.
type HTTP struct {
	svc destroy.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc destroy.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/destroy")
	cr.POST("/deletefile", h.deletefile)
	cr.POST("/disconnectuser", h.disconnectuser)
	cr.POST("/killprocess/:pid", h.killprocess)
}

func (h *HTTP) deletefile(ctx echo.Context) error {
	path := ctx.QueryParam("filepath")
	if path == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - path is null")
	}

	result, err := h.svc.ExecuteDF(path)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) disconnectuser(ctx echo.Context) error {
	terminalname := ctx.QueryParam("processname")
	if terminalname == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - processname is null")
	}

	result, err := h.svc.ExecuteDU(terminalname, "terminal")
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) killprocess(ctx echo.Context) error {
	pid, err := strconv.Atoi(ctx.Param("pid"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request due to an invalid pid - should be an integer")
	}

	result, err := h.svc.ExecuteKP(pid)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
