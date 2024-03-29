package mocksys

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// Host mock
type Host struct {
	ListFn func(
		host.InfoStat,
		[]host.UserStat,
		[]cpu.InfoStat,
		[]float64,
		mem.VirtualMemoryStat,
		mem.SwapMemoryStat,
		load.AvgStat,
		string,
		string,
		string,
		map[string][]metrics.DStats,
		[]net.InterfaceStat) (rpi.Host, error)
}

// List mock
func (h Host) List(infos host.InfoStat,
	users []host.UserStat,
	cpus []cpu.InfoStat,
	vcores []float64,
	vmem mem.VirtualMemoryStat,
	smem mem.SwapMemoryStat,
	load load.AvgStat,
	temp string,
	serialNumber string,
	rpiv string,
	listDev map[string][]metrics.DStats,
	netInfo []net.InterfaceStat) (rpi.Host, error) {
	return h.ListFn(infos, users, cpus, vcores, vmem, smem, load, temp, serialNumber, rpiv, listDev, netInfo)
}
