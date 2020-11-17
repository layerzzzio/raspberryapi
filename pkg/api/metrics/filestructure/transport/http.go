package transport

import (
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
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
	directoryPath := ctx.QueryParam("directoryPath")
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - directoryPath is not a directory")
	}

	fileLimit, err := strconv.Atoi(ctx.QueryParam("fileLimit"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Not found - fileLimit is not an integer")
	}

	pathSize, err := strconv.ParseUint(ctx.QueryParam("pathSize"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Not found - pathSize is not an integer")
	}

	result, err := h.svc.ViewLF(directoryPath, pathSize, int8(fileLimit))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}
