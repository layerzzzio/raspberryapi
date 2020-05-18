package sys

import (
	"fmt"
	"net/http"
	"regexp"
	"sort"

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
				Fstype:      fsType(devMP),
				Mountpoints: devMP,
			})

		devMP = nil
	}

	sort.Slice(result[:], func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result, nil
}

// View returns a disk stats
func (d Disk) View(device string, listDev map[string][]metrics.DStats) (rpi.Disk, error) {
	var result rpi.Disk
	var devMP []rpi.MountPoint
	isDiskFound := false

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

			result = rpi.Disk{
				ID:          id[0],
				Filesystem:  dev,
				Fstype:      fsType(devMP),
				Mountpoints: devMP,
			}
			isDiskFound = true
			break
		}
	}

	if isDiskFound == false {
		return rpi.Disk{}, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%v does not exist", device))
	}

	return result, nil
}

func extractDeviceID(dev string) []string {
	r := regexp.MustCompile("[^\"/]+$")
	res := r.FindAllString(dev, -1)
	return res
}

func fsType(mp []rpi.MountPoint) string {
	var fstypes []string

	for i := range mp {
		if len(fstypes) > 0 {
			if fstypes[i-1] != mp[i].Fstype {
				fstypes = append(fstypes, mp[i].Fstype)
			} else {
				fstypes = nil
			}
		} else {
			fstypes = append(fstypes, mp[i].Fstype)
		}
	}

	var fstype string
	if len(fstypes) == 1 {
		fstype = mp[0].Fstype
	} else {
		fstype = "multi_fstype"
	}
	return fstype
}
