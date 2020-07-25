package largestfile

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of LargestFile model.
func (lf *LargestFile) List() ([]rpi.LargestFile, error) {
	top100files, errStd, err := lf.mt.Top100Files()

	if err != nil && (errStd != "" || errStd == "") {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the largest files")
	}

	return lf.lfsys.List(top100files)
}
