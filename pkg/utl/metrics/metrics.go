package metrics

import (
	"bytes"
	"errors"
	"os/exec"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

// TODO: to test this method by simulating different OS scenarios in a Docker container (raspbian/strech)

// Service is
type Service struct {
	m Metrics
}

// Metrics is
type Metrics interface {
	PsPID(p *process.Process, c chan (int32))
	PsName(p *process.Process, c chan (string))
	PsCPUPer(p *process.Process, c chan (float64))
	PsMemPer(p *process.Process, c chan (float32))
	PsUsername(p *process.Process, c chan (string))
	PsCmdLine(p *process.Process, c chan (string))
	PsStatus(p *process.Process, c chan (string))
	PsCreationTime(p *process.Process, c chan (int64))
	PsBackground(p *process.Process, c chan (bool))
	PsForeground(p *process.Process, c chan (bool))
	PsIsRunning(p *process.Process, c chan (bool))
	PsParent(p *process.Process, c chan (int32))
}

// New creates a Process application service instance.
func New(m Metrics) *Service {
	return &Service{m: m}
}

// DStats represents a pair of partition stats and mount point usage stats
type DStats struct {
	Partition  *disk.PartitionStat
	Mountpoint *disk.UsageStat
}

// PInfo represents multiple aspects of a process
type PInfo struct {
	ID           int32
	Name         string
	Username     string
	CommandLine  string
	Status       string
	CreationTime int64
	Foreground   bool
	Background   bool
	IsRunning    bool
	CPUPercent   float64
	MemPercent   float32
	ParentP      int32
}

// CPUInfo is
func (s Service) CPUInfo() ([]cpu.InfoStat, error) {
	info, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	return info, nil
}

// CPUPercent is
func (s Service) CPUPercent(interval time.Duration, perVCore bool) ([]float64, error) {
	percent, err := cpu.Percent(interval, perVCore)
	if err != nil {
		return nil, err
	}
	return percent, nil
}

// CPUTimes is
func (s Service) CPUTimes(perVCore bool) ([]cpu.TimesStat, error) {
	times, err := cpu.Times(perVCore)
	if err != nil {
		return nil, err
	}
	return times, nil
}

// SwapMemory is
func (s Service) SwapMemory() (mem.SwapMemoryStat, error) {
	smem, err := mem.SwapMemory()
	if err != nil {
		return mem.SwapMemoryStat{}, err
	}
	return *smem, nil
}

// VirtualMemory is
func (s Service) VirtualMemory() (mem.VirtualMemoryStat, error) {
	vmem, err := mem.VirtualMemory()
	if err != nil {
		return mem.VirtualMemoryStat{}, err
	}
	return *vmem, nil
}

// DiskStats is
func (s Service) DiskStats(all bool) (map[string][]DStats, error) {
	dstats := make(map[string][]DStats)

	disks, err := disk.Partitions(all)
	if err != nil {
		return nil, err
	}

	for _, d := range disks {
		usage, err := disk.Usage(d.Mountpoint)
		if err != nil {
			return nil, err
		}

		dstats[d.Device] = append(
			dstats[d.Device],
			DStats{
				&disk.PartitionStat{
					Fstype:     d.Fstype,
					Opts:       d.Opts,
					Mountpoint: d.Mountpoint,
				},
				&disk.UsageStat{
					Path:              usage.Path,
					Fstype:            usage.Fstype,
					Total:             usage.Total,
					Free:              usage.Free,
					Used:              usage.Used,
					UsedPercent:       usage.UsedPercent,
					InodesTotal:       usage.InodesTotal,
					InodesFree:        usage.InodesFree,
					InodesUsed:        usage.InodesUsed,
					InodesUsedPercent: usage.InodesUsedPercent,
				},
			},
		)
	}

	return dstats, nil
}

// LoadAvg is
func (s Service) LoadAvg() (load.AvgStat, error) {
	temp, err := load.Avg()
	if err != nil {
		return load.AvgStat{}, err
	}
	return *temp, nil
}

// LoadProcs is
func (s Service) LoadProcs() (load.MiscStat, error) {
	procs, err := load.Misc()
	if err != nil {
		return load.MiscStat{}, err
	}
	return *procs, nil
}

//Processes is
func (s Service) Processes(id ...int32) ([]PInfo, error) {
	var pinfo []PInfo
	cID := make(chan (int32))
	cName := make(chan (string))
	cCPUPer := make(chan (float64))
	cMemPer := make(chan (float32))
	var cU chan (string)
	var cCL chan (string)
	var cS chan (string)
	var cCT chan (int64)
	var cFG chan (bool)
	var cBG chan (bool)
	var cIR chan (bool)
	var cP chan (int32)

	ps, err := process.Processes()
	if err != nil {
		error.Error(err)
	}

	pid := int32(-1)
	if len(id) == 1 {
		pid = id[0]
		if id[0] > 0 {
			cU = make(chan (string))
			cCL = make(chan (string))
			cS = make(chan (string))
			cCT = make(chan (int64))
			cFG = make(chan (bool))
			cBG = make(chan (bool))
			cIR = make(chan (bool))
			cP = make(chan (int32))
		}
	} else if len(id) > 1 {
		panic("only one id is authorized")
	}

	if pid > 0 {
		for i := range ps {
			if ps[i].Pid == pid {
				go s.m.PsPID(ps[i], cID)
				go s.m.PsName(ps[i], cName)
				go s.m.PsCPUPer(ps[i], cCPUPer)
				go s.m.PsMemPer(ps[i], cMemPer)
				go s.m.PsUsername(ps[i], cU)
				go s.m.PsCmdLine(ps[i], cCL)
				go s.m.PsStatus(ps[i], cS)
				go s.m.PsCreationTime(ps[i], cCT)
				go s.m.PsForeground(ps[i], cFG)
				go s.m.PsBackground(ps[i], cBG)
				go s.m.PsIsRunning(ps[i], cIR)
				go s.m.PsParent(ps[i], cP)

				pinfo = []PInfo{
					{
						ID:           <-cID,
						Name:         <-cName,
						CPUPercent:   <-cCPUPer,
						MemPercent:   <-cMemPer,
						Username:     <-cU,
						CommandLine:  <-cCL,
						Status:       <-cS,
						CreationTime: <-cCT,
						Foreground:   <-cFG,
						Background:   <-cBG,
						IsRunning:    <-cIR,
						ParentP:      <-cP,
					},
				}

				close(cID)
				close(cName)
				close(cCPUPer)
				close(cMemPer)
				close(cU)
				close(cCL)
				close(cS)
				close(cCT)
				close(cFG)
				close(cBG)
				close(cIR)
				close(cP)

				return pinfo, nil
			}
		}
		return nil, errors.New("process not found")
	} else if pid == 0 {
		return nil, errors.New("process not found")
	}

	for i := range ps {
		go s.m.PsPID(ps[i], cID)
		go s.m.PsName(ps[i], cName)
		go s.m.PsCPUPer(ps[i], cCPUPer)
		go s.m.PsMemPer(ps[i], cMemPer)

		pinfo = append(pinfo,
			PInfo{
				ID:         <-cID,
				Name:       <-cName,
				CPUPercent: <-cCPUPer,
				MemPercent: <-cMemPer,
			})
	}

	close(cID)
	close(cName)
	close(cCPUPer)
	close(cMemPer)

	return pinfo, nil
}

// PsPID is
func (s Service) PsPID(p *process.Process, c chan (int32)) {
	c <- p.Pid
}

// PsName is
func (s Service) PsName(p *process.Process, c chan (string)) {
	name, err := p.Name()
	if err != nil {
		log.Error()
	}
	c <- name
}

// PsCPUPer is
func (s Service) PsCPUPer(p *process.Process, c chan (float64)) {
	cpuper, err := p.CPUPercent()
	if err != nil {
		log.Error()
	}
	c <- cpuper
}

// PsMemPer is
func (s Service) PsMemPer(p *process.Process, c chan (float32)) {
	memper, err := p.MemoryPercent()
	if err != nil {
		log.Error()
	}
	c <- memper
}

// PsUsername is
func (s Service) PsUsername(p *process.Process, c chan (string)) {
	u, err := p.Username()
	if err != nil {
		log.Error()
	}
	c <- u
}

// PsCmdLine is
func (s Service) PsCmdLine(p *process.Process, c chan (string)) {
	cl, err := p.Cmdline()
	if err != nil {
		log.Error()
	}
	c <- cl
}

// PsStatus is
func (s Service) PsStatus(p *process.Process, c chan (string)) {
	st, err := p.Status()
	if err != nil {
		log.Error()
	}
	c <- st
}

// PsCreationTime is
func (s Service) PsCreationTime(p *process.Process, c chan (int64)) {
	ct, err := p.CreateTime()
	if err != nil {
		log.Error()
	}

	c <- ct
}

// PsBackground is
func (s Service) PsBackground(p *process.Process, c chan (bool)) {
	bg, err := p.Background()
	if err != nil {
		log.Error()
	}
	c <- bg
}

// PsForeground is
func (s Service) PsForeground(p *process.Process, c chan (bool)) {
	fg, err := p.Foreground()
	if err != nil {
		log.Error()
	}
	c <- fg
}

// PsIsRunning is
func (s Service) PsIsRunning(p *process.Process, c chan (bool)) {
	ir, err := p.IsRunning()
	if err != nil {
		log.Error()
	}
	c <- ir
}

// PsParent is
func (s Service) PsParent(p *process.Process, c chan (int32)) {
	var ppid int32

	ps, err := p.Parent()

	if err != nil {
		log.Error()
	}

	if ps == nil {
		ppid = -1
	} else {
		ppid = ps.Pid
	}

	c <- ppid
}

// HostInfo is
func (s Service) HostInfo() (host.InfoStat, error) {
	info, err := host.Info()
	if err != nil {
		return host.InfoStat{}, err
	}
	return *info, nil
}

// Users is
func (s Service) Users() ([]host.UserStat, error) {
	users, err := host.Users()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Temperature is
func (s Service) Temperature() (string, string, error) {
	cmd := exec.Command("sh", "-c", "vcgencmd measure_temp")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Error()
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	return outStr, errStr, nil
}

// RaspModel is
func (s Service) RaspModel() (string, string, error) {
	cmd := exec.Command("cat", "/sys/firmware/devicetree/base/model")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Error()
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	return outStr, errStr, nil
}

// NetInfo is
func (s Service) NetInfo() ([]net.InterfaceStat, error) {
	netInfo, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	return netInfo, nil
}

// NetStats is
func (s Service) NetStats() ([]net.IOCountersStat, error) {
	netStats, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}
	return netStats, nil
}
