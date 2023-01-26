package humanuser

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of HumanUser model.
func (hu *HumanUser) List() ([]rpi.HumanUser, error) {
	etcPasswdPath := hu.i.GetConfigFiles()["etcpasswd"].Path
	humanUsers, err := hu.i.ReadFile(etcPasswdPath)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the human users")
	}

	return hu.humsys.List(humanUsers)
}
