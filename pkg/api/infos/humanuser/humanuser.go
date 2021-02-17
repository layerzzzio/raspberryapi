package humanuser

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of HumanUser model.
func (hu *HumanUser) List() ([]rpi.HumanUser, error) {
	humanUsers, err := hu.i.ReadFile("/etc/passwd")

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the human users")
	}

	return hu.humsys.List(humanUsers)
}
