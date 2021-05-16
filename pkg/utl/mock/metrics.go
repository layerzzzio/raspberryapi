package mock

import (
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

// Metrics mock
type Metrics struct {
	CPUInfoFn        func() ([]cpu.InfoStat, error)
	CPUPercentFn     func(time.Duration, bool) ([]float64, error)
	CPUTimesFn       func(bool) ([]cpu.TimesStat, error)
	SwapMemFn        func() (mem.SwapMemoryStat, error)
	VirtualMemFn     func() (mem.VirtualMemoryStat, error)
	DiskStatsFn      func(bool) (map[string][]metrics.DStats, error)
	LoadAvgFn        func() (load.AvgStat, error)
	LoadProcsFn      func() (load.MiscStat, error)
	ProcessesFn      func(id ...int32) ([]metrics.PInfo, error)
	PsPIDFn          func(p *process.Process, c chan (int32))
	PsNameFn         func(p *process.Process, c chan (string))
	PsCPUPerFn       func(p *process.Process, c chan (float64))
	PsMemPerFn       func(p *process.Process, c chan (float32))
	PsUsernameFn     func(p *process.Process, c chan (string))
	PsCmdLineFn      func(p *process.Process, c chan (string))
	PsStatusFn       func(p *process.Process, c chan (string))
	PsCreationTimeFn func(p *process.Process, c chan (int64))
	PsBackgroundFn   func(p *process.Process, c chan (bool))
	PsForegroundFn   func(p *process.Process, c chan (bool))
	PsIsRunningFn    func(p *process.Process, c chan (bool))
	PsParentFn       func(p *process.Process, c chan (int32))
	HostInfoFn       func() (host.InfoStat, error)
	UsersFn          func() ([]host.UserStat, error)
	TemperatureFn    func() (string, string, error)
	RaspModelFn      func() (string, string, error)
	NetInfoFn        func() ([]net.InterfaceStat, error)
	NetStatsFn       func() ([]net.IOCountersStat, error)
	WalkFolderFn     func(
		string,
		metrics.ReadDir,
		uint64,
		float32,
		metrics.ShouldIgnoreFolder,
		chan int,
	) (*rpi.File, map[string]int64)
}

// CPUInfo mock
func (m Metrics) CPUInfo() ([]cpu.InfoStat, error) {
	return m.CPUInfoFn()
}

// CPUPercent mock
func (m Metrics) CPUPercent(interval time.Duration, perVCore bool) ([]float64, error) {
	return m.CPUPercentFn(interval, perVCore)
}

// CPUTimes mock
func (m Metrics) CPUTimes(perVCore bool) ([]cpu.TimesStat, error) {
	return m.CPUTimesFn(perVCore)
}

// SwapMemory mock
func (m Metrics) SwapMemory() (mem.SwapMemoryStat, error) {
	return m.SwapMemFn()
}

// VirtualMemory mock
func (m Metrics) VirtualMemory() (mem.VirtualMemoryStat, error) {
	return m.VirtualMemFn()
}

// DiskStats mock
func (m Metrics) DiskStats(all bool) (map[string][]metrics.DStats, error) {
	return m.DiskStatsFn(all)
}

// LoadAvg mock
func (m Metrics) LoadAvg() (load.AvgStat, error) {
	return m.LoadAvgFn()
}

// LoadProcs mock
func (m Metrics) LoadProcs() (load.MiscStat, error) {
	return m.LoadProcsFn()
}

// Processes mock
func (m Metrics) Processes(id ...int32) ([]metrics.PInfo, error) {
	if len(id) > 1 && id[0] > 0 {
		return m.ProcessesFn(id[0])
	}
	return m.ProcessesFn()
}

// PsPID mock
func (m Metrics) PsPID(p *process.Process, c chan (int32)) {
	m.PsPIDFn(p, c)
}

// PsName mock
func (m Metrics) PsName(p *process.Process, c chan (string)) {
	m.PsNameFn(p, c)
}

// PsCPUPer mock
func (m Metrics) PsCPUPer(p *process.Process, c chan (float64)) {
	m.PsCPUPerFn(p, c)
}

// PsMemPer mock
func (m Metrics) PsMemPer(p *process.Process, c chan (float32)) {
	m.PsMemPerFn(p, c)
}

// PsUsername mock
func (m Metrics) PsUsername(p *process.Process, c chan (string)) {
	m.PsUsernameFn(p, c)
}

// PsCmdLine mock
func (m Metrics) PsCmdLine(p *process.Process, c chan (string)) {
	m.PsCmdLineFn(p, c)
}

// PsStatus mock
func (m Metrics) PsStatus(p *process.Process, c chan (string)) {
	m.PsStatusFn(p, c)
}

// PsCreationTime mock
func (m Metrics) PsCreationTime(p *process.Process, c chan (int64)) {
	m.PsCreationTimeFn(p, c)
}

// PsForeground mock
func (m Metrics) PsForeground(p *process.Process, c chan (bool)) {
	m.PsForegroundFn(p, c)
}

// PsBackground mock
func (m Metrics) PsBackground(p *process.Process, c chan (bool)) {
	m.PsBackgroundFn(p, c)
}

// PsIsRunning mock
func (m Metrics) PsIsRunning(p *process.Process, c chan (bool)) {
	m.PsIsRunningFn(p, c)
}

// PsParent mock
func (m Metrics) PsParent(p *process.Process, c chan (int32)) {
	m.PsParentFn(p, c)
}

// HostInfo mock
func (m Metrics) HostInfo() (host.InfoStat, error) {
	return m.HostInfoFn()
}

// Users mock
func (m Metrics) Users() ([]host.UserStat, error) {
	return m.UsersFn()
}

// Temperature mock
func (m Metrics) Temperature() (string, string, error) {
	return m.TemperatureFn()
}

// RaspModel mock
func (m Metrics) RaspModel() (string, string, error) {
	return m.RaspModelFn()
}

// NetInfo mock
func (m Metrics) NetInfo() ([]net.InterfaceStat, error) {
	return m.NetInfoFn()
}

// NetStats mock
func (m Metrics) NetStats() ([]net.IOCountersStat, error) {
	return m.NetStatsFn()
}

// WalkFolder mock
func (m Metrics) WalkFolder(
	path string,
	readDir metrics.ReadDir,
	pathSize uint64,
	fileLimit float32,
	ignoreFunction metrics.ShouldIgnoreFolder,
	progress chan int,
) (*rpi.File, map[string]int64) {
	return m.WalkFolderFn(path, readDir, pathSize, fileLimit, ignoreFunction, progress)
}
