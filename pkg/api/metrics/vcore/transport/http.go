package transport

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi/pkg/api/metrics/vcore"
)

// HTTP is a struct implementing a core application service.
type HTTP struct {
	svc vcore.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc vcore.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/vcores")
	cr.GET("", h.list)
	cr.GET("/:id", h.view)
}

func (h *HTTP) list(ctx echo.Context) error {
	result, err := h.svc.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

func (h *HTTP) view(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request due to an invalid vcore id - should be an integer")
	}

	result, err := h.svc.View(id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}
