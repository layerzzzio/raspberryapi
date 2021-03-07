package transport

import (
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/metrics/filestructure"
)

// HTTP is a struct implementing a file structure application service.
type HTTP struct {
	svc filestructure.Service
}

// NewHTTP creates new filestructure http service
func NewHTTP(svc filestructure.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/filestructure")
	cr.GET("/largestfiles", h.ViewLF)
}

func (h *HTTP) ViewLF(ctx echo.Context) error {
	directoryPath := ctx.QueryParam("directorypath")
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - directorypath is not a directory")
	}

	fileLimit, err := strconv.ParseFloat(ctx.QueryParam("filelimit"), 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - filelimit is not an integer")
	}

	pathSize, err := strconv.ParseUint(ctx.QueryParam("pathsize"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - pathsize is not an integer")
	}

	result, err := h.svc.ViewLF(directoryPath, pathSize, float32(fileLimit))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}
