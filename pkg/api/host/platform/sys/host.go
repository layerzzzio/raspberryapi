package sys

import (
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// Host represents an empty Host entity on the current system.
type Host struct{}

// List returns a list of Host info
func (h Host) List(info host.InfoStat,
	users []host.UserStat,
	cpus []cpu.InfoStat,
	vcores []float64,
	vMemPer mem.VirtualMemoryStat,
	sMemPer mem.SwapMemoryStat) (rpi.Host, error) {

	virtualUsers := uint16(len(users))
	cpuCount := uint8(len(cpus))
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
		Hostname:           info.Hostname,
		Uptime:             info.Uptime,
		BootTime:           info.BootTime,
		OS:                 info.OS,
		Platform:           info.Platform,
		PlatformFamily:     info.PlatformFamily,
		PlatformVersion:    info.PlatformVersion,
		KernelArch:         info.KernelArch,
		KernelVersion:      info.KernelVersion,
		CPU:                cpuCount,
		VCore:              vCoresCount,
		CPUUsedPercent:     cpuPer,
		VUsedPercent:       vmemPer,
		SUsedPercent:       smenPer,
		Processes:          info.Procs,
		ActiveVirtualUsers: virtualUsers,
	}
	return result, nil
}
