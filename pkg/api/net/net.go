package net

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of Net model.
func (n *Net) List() ([]rpi.Net, error) {
	netInfo, errI := n.mt.NetInfo()

	if errI != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "could not list the net metrics")
	}

	return n.nsys.List(netInfo)
}

// View populates and returns a Net model.
func (n *Net) View(id int) (rpi.Net, error) {
	netInfo, errI := n.mt.NetInfo()
	netStats, errS := n.mt.NetStats()

	if errI != nil || errS != nil {
		return rpi.Net{}, echo.NewHTTPError(http.StatusInternalServerError, "could not view the net metrics")
	}

	return n.nsys.View(id, netInfo, netStats)
}
