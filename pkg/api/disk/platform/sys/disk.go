package sys

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
)

// Disk represents an empty disk entity on the current system.
type Disk struct{}

// List returns a list of disk stats
func (d Disk) List(listDev map[string][]metrics.DStats) ([]rpi.Disk, error) {
	var result []rpi.Disk
	var devMP []rpi.MountPoint

	for dev, dstats := range listDev {
		for _, v := range dstats {
			devMP = append(
				devMP,
				rpi.MountPoint{
					Mountpoint:        v.Mountpoint.Path,
					Fstype:            v.Mountpoint.Fstype,
					Opts:              v.Partition.Opts,
					Total:             v.Mountpoint.Total,
					Free:              v.Mountpoint.Free,
					Used:              v.Mountpoint.Used,
					UsedPercent:       v.Mountpoint.UsedPercent,
					InodesTotal:       v.Mountpoint.InodesTotal,
					InodesFree:        v.Mountpoint.InodesFree,
					InodesUsed:        v.Mountpoint.InodesUsed,
					InodesUsedPercent: v.Mountpoint.InodesUsedPercent,
				},
			)
		}

		id := extractDeviceID(dev)
		if len(id) != 1 {
			return nil, echo.NewHTTPError(http.StatusNotFound, "parsing id was unsuccessful")
		}

		result = append(
			result,
			rpi.Disk{
				ID:          id[0],
				Filesystem:  dev,
				Mountpoints: devMP,
			})

		devMP = nil
	}
	return result, nil
}

// View returns a disk stats
func (d Disk) View(device string, listDev map[string][]metrics.DStats) (rpi.Disk, error) {
	var result rpi.Disk
	var devMP []rpi.MountPoint

	for dev, dstats := range listDev {
		id := extractDeviceID(dev)
		if len(id) != 1 {
			return rpi.Disk{}, echo.NewHTTPError(http.StatusNotFound, "parsing id was unsuccessful")
		}

		if id[0] == device {
			for _, v := range dstats {
				devMP = append(
					devMP,
					rpi.MountPoint{
						Mountpoint:        v.Mountpoint.Path,
						Fstype:            v.Mountpoint.Fstype,
						Opts:              v.Partition.Opts,
						Total:             v.Mountpoint.Total,
						Free:              v.Mountpoint.Free,
						Used:              v.Mountpoint.Used,
						UsedPercent:       v.Mountpoint.UsedPercent,
						InodesTotal:       v.Mountpoint.InodesTotal,
						InodesFree:        v.Mountpoint.InodesFree,
						InodesUsed:        v.Mountpoint.InodesUsed,
						InodesUsedPercent: v.Mountpoint.InodesUsedPercent,
					},
				)
			}
			break
		}

		result = rpi.Disk{
			ID:          id[0],
			Filesystem:  dev,
			Mountpoints: devMP,
		}
	}
	return result, nil
}

func extractDeviceID(s string) []string {
	r := regexp.MustCompile("[^\"/]+$")
	res := r.FindAllString(s, -1)
	return res
}
