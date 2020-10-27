package largestfile

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
)

// View populates and returns an array of LargestFile model.
func (lf *LargestFile) View(path string) ([]rpi.LargestFile, error) {
	top100files, errStd, err := lf.mt.Top100Files(path)

	if err != nil && (errStd != "" || errStd == "") {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the largest files")
	}

	return lf.lfsys.View(top100files)
}
