package mocksys

import (
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
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
		string) (rpi.Host, error)
}

// List mock
func (h Host) List(infos host.InfoStat,
	users []host.UserStat,
	cpus []cpu.InfoStat,
	vcores []float64,
	vmem mem.VirtualMemoryStat,
	smem mem.SwapMemoryStat,
	temp string) (rpi.Host, error) {
	return h.ListFn(infos, users, cpus, vcores, vmem, smem, temp)
}
