package mocksys

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
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
		map[string][]metrics.DStats) (rpi.Host, error)
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
	rpiv string,
	listDev map[string][]metrics.DStats) (rpi.Host, error) {
	return h.ListFn(infos, users, cpus, vcores, vmem, smem, load, temp, rpiv, listDev)
}
