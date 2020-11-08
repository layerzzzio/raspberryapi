package transport

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi/pkg/api/metrics/largestfile"
)

// HTTP is a struct implementing a largest files application service.
type HTTP struct {
	svc largestfile.Service
}

// NewHTTP creates new largestfile http service
func NewHTTP(svc largestfile.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/largestfiles")
	cr.GET("", h.view)
}

func (h *HTTP) view(ctx echo.Context) error {
	directorypath := ctx.QueryParam("directorypath")
	if directorypath == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Not found - directory path is null")
	}

	result, err := h.svc.View(directorypath)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}
