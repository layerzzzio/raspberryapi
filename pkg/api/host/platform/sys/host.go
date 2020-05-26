package sys

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// Host represents a host entity.
type Host struct{}

// List returns a list of Host info
func (h Host) List(info host.InfoStat,
	users []host.UserStat,
	cpus []cpu.InfoStat,
	vcores []float64,
	vMemPer mem.VirtualMemoryStat,
	sMemPer mem.SwapMemoryStat,
	temp string,
	rpiv string) (rpi.Host, error) {
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

	vmemPer := vMemPer.UsedPercent
	smenPer := sMemPer.UsedPercent

	result := rpi.Host{
		ID:                 info.HostID,
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
		CPUUsedPercent:     cpuPer,
		VUsedPercent:       vmemPer,
		SUsedPercent:       smenPer,
		Processes:          info.Procs,
		ActiveVirtualUsers: virtualUsers,
		Temperature:        extractTemp(temp),
	}
	return result, nil
}

func extractTemp(s string) float32 {
	r := regexp.MustCompile("[" + strconv.Itoa(0) + "-" + strconv.Itoa(9) + "]+")
	num := r.FindAllString(s, -1)
	temp := strings.Join(num[:], ".")

	res, err := strconv.ParseFloat(temp, 16)
	if err != nil {
		return -1
	}

	return float32(res)
}
