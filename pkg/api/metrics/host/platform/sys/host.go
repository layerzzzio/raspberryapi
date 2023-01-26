package sys

import (
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	disksys "github.com/raspibuddy/rpi/pkg/api/metrics/disk/platform/sys"
	netsys "github.com/raspibuddy/rpi/pkg/api/metrics/net/platform/sys"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// Host represents a host entity.
type Host struct{}

// List returns a list of Host info
func (h Host) List(
	info host.InfoStat,
	users []host.UserStat,
	cpus []cpu.InfoStat,
	vcores []float64,
	vMem mem.VirtualMemoryStat,
	sMemPer mem.SwapMemoryStat,
	load load.AvgStat,
	temp string,
	serialNumber string,
	rpiv string,
	listDev map[string][]metrics.DStats,
	netInfo []net.InterfaceStat) (rpi.Host, error) {
	hyperThreading := false
	virtualUsers := uint16(len(users))
	cpuCount := uint8(len(cpus))
	if cpuCount > 1 {
		hyperThreading = true
	}

	vCoresCount := uint8(len(vcores))
	var cpuPer float64
	if vcores != nil {
		for i := range vcores {
			cpuPer += vcores[i]
		}
		cpuPer = cpuPer / float64(vCoresCount)
	} else {
		cpuPer = 0
	}

	vTot := vMem.Total
	vmemPer := vMem.UsedPercent
	smenPer := sMemPer.UsedPercent

	var disks []rpi.Disk
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

		id := disksys.ExtractDeviceID(dev)
		if len(id) != 1 {
			return rpi.Host{}, echo.NewHTTPError(http.StatusNotFound, "parsing id was unsuccessful")
		}

		disks = append(
			disks,
			rpi.Disk{
				ID:          id[0],
				Filesystem:  dev,
				Fstype:      disksys.FsType(devMP),
				Mountpoints: devMP,
			})

		devMP = nil
	}

	sort.Slice(disks[:], func(i, j int) bool {
		return disks[i].ID < disks[j].ID
	})

	var nets []rpi.Net
	for i := range netInfo {
		data := rpi.Net{
			ID:    netInfo[i].Index,
			Name:  netInfo[i].Name,
			Flags: netInfo[i].Flags,
			IPv4:  netsys.ExtractIPv4(netInfo[i].Addrs),
		}
		nets = append(nets, data)
	}

	var allUsers []rpi.User
	for i := range users {
		data := rpi.User{
			User:     users[i].User,
			Terminal: users[i].Terminal,
			Started:  users[i].Started,
		}
		allUsers = append(allUsers, data)
	}

	result := rpi.Host{
		ID:                 serialNumber,
		RaspModel:          rpiv,
		Hostname:           info.Hostname,
		UpTime:             info.Uptime,
		BootTime:           info.BootTime,
		OS:                 info.OS,
		Platform:           info.Platform,
		PlatformFamily:     info.PlatformFamily,
		PlatformVersion:    info.PlatformVersion,
		KernelArch:         info.KernelArch,
		KernelVersion:      info.KernelVersion,
		CPU:                cpuCount,
		HyperThreading:     hyperThreading,
		VCore:              vCoresCount,
		VTotal:             vTot,
		CPUUsedPercent:     cpuPer,
		VUsedPercent:       vmemPer,
		SUsedPercent:       smenPer,
		Load1:              load.Load1,
		Load5:              load.Load5,
		Load15:             load.Load15,
		Processes:          info.Procs,
		ActiveVirtualUsers: virtualUsers,
		Users:              allUsers,
		Temperature:        extractTemp(temp),
		Disks:              disks,
		Nets:               nets,
	}
	return result, nil
}

func extractTemp(s string) float32 {
	r := regexp.MustCompile("[" + strconv.Itoa(0) + "-" + strconv.Itoa(9) + "]+")
	num := r.FindAllString(s, -1)
	temp := strings.Join(num[:], ".")

	res, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		return -100
	}
	return float32(res)
}
