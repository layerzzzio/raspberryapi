package sys

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/disk"
)

// Disk represents an empty disk entity on the current system.
type Disk struct{}

// List returns a list of disk stats.
func (d Disk) List(listDisks map[string]disk.PartitionStat, stats map[string]*disk.UsageStat) ([]rpi.Disk, error) {
	var result []rpi.Disk
	deviceMap := make(map[string][]rpi.MountPoint)
	for _, v := range listDisks {
		s := stats[v.Device]
		deviceMap[v.Device] = append(
			deviceMap[v.Device],
			rpi.MountPoint{
				Mountpoint:        v.Mountpoint,
				Fstype:            v.Fstype,
				Opts:              v.Opts,
				Total:             s.Total,
				Free:              s.Free,
				Used:              s.Used,
				UsedPercent:       s.UsedPercent,
				InodesTotal:       s.InodesTotal,
				InodesFree:        s.InodesFree,
				InodesUsed:        s.InodesUsed,
				InodesUsedPercent: s.InodesUsedPercent,
			},
		)
	}

	for k := range deviceMap {
		device := extractDeviceID(k)
		if len(device) != 1 {
			return nil, echo.NewHTTPError(http.StatusNotFound, "parsing id was unsuccessful")
		}
		result = append(
			result,
			rpi.Disk{
				ID:          device[0],
				Filesystem:  listDisks[k].Device,
				Fstype:      listDisks[k].Fstype,
				Opts:        listDisks[k].Opts,
				Mountpoints: deviceMap[k],
			})

	}
	return result, nil
}

// View returns a disk stats.
func (d Disk) View(dev string, listDisks map[string]disk.PartitionStat, stats map[string]*disk.UsageStat) (rpi.Disk, error) {
	var result rpi.Disk
	deviceMap := make(map[string][]rpi.MountPoint)

	for _, v := range listDisks {
		device := extractDeviceID(v.Device)

		if len(device) != 1 {
			return rpi.Disk{}, echo.NewHTTPError(http.StatusNotFound, "parsing id was unsuccessful")
		}

		if device[0] == dev {
			s := stats[v.Device]
			deviceMap[v.Device] = append(
				deviceMap[v.Device],
				rpi.MountPoint{
					Mountpoint:        v.Mountpoint,
					Fstype:            v.Fstype,
					Opts:              v.Opts,
					Total:             s.Total,
					Free:              s.Free,
					Used:              s.Used,
					UsedPercent:       s.UsedPercent,
					InodesTotal:       s.InodesTotal,
					InodesFree:        s.InodesFree,
					InodesUsed:        s.InodesUsed,
					InodesUsedPercent: s.InodesUsedPercent,
				},
			)
			break
		}
	}

	for k := range deviceMap {
		device := extractDeviceID(k)
		if len(device) != 1 {
			return rpi.Disk{}, echo.NewHTTPError(http.StatusNotFound, "parsing id was unsuccessful")
		}
		result = rpi.Disk{
			ID:          device[0],
			Filesystem:  listDisks[k].Device,
			Fstype:      listDisks[k].Fstype,
			Opts:        listDisks[k].Opts,
			Mountpoints: deviceMap[k],
		}
	}
	return result, nil
}

func extractDeviceID(s string) []string {
	r := regexp.MustCompile("[^\"/]+$")
	res := r.FindAllString(s, -1)
	return res
}
