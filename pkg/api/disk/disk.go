package disk

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of Disk models.
func (d *Disk) List() ([]rpi.Disk, error) {
	dstats, err := d.m.DiskStats(false)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "could not list the disk metrics")
	}

	return d.dsys.List(dstats)
}

//View populates and returns a Disk model.
func (d *Disk) View(dev string) (rpi.Disk, error) {
	dstats, err := d.m.DiskStats(false)

	if err != nil {
		return rpi.Disk{}, echo.NewHTTPError(http.StatusInternalServerError, "could not view the disk metrics")
	}

	return d.dsys.View(dev, dstats)
}
