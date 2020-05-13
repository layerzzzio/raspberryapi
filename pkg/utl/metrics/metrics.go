package metrics

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
)

// TODO: to test this method by simulating different OS scenarios in a Docker container (raspbian/strech)

// Service is
type Service struct{}

// DStats represents a pair of partition stats and mount point usage stats
type DStats struct {
	Partition  disk.PartitionStat
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
	CPUTimes     cpu.TimesStat
	Threads      int32
	MemPercent   float32
	MemInfo      process.MemoryInfoStat
	ParentP      process.Process
	ChildrenP    []*process.Process
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
				disk.PartitionStat{
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

//RunningProcess is
func (s Service) RunningProcess() (map[int32][]PInfo, error) {
	var pinfo map[int32][]PInfo

	ps, err := process.Processes()
	if err != nil {
		error.Error(err)
	}

	for i := range ps {
		p := ps[i]

		// ID
		id := p.Pid
		name, errN := p.Name()
		cmdline, errCL := p.Cmdline()
		username, errUN := p.Username()

		// Status
		status, errS := p.Status()
		ctime, errCR := p.CreateTime()
		backgrd, errBG := p.Background()
		foregrd, errFG := p.Foreground()
		isrunning, errIR := p.IsRunning()

		// Stats
		cpuper, errCP := p.CPUPercent()
		memper, errMP := p.MemoryPercent()
		meminfo, errMI := p.MemoryInfo()
		times, errT := p.Times()
		threads, errTH := p.NumThreads()

		// Structure
		parent, errP := p.Parent()
		children, errCH := p.Children()

		if errN != nil || errCL != nil || errUN != nil || errS != nil || errCR != nil || errBG != nil || errFG != nil || errIR != nil || errCP != nil || errMP != nil || errMI != nil || errT != nil || errTH != nil || errP != nil || errCH != nil {
			return nil, nil
		}

		pinfo[id] = append(
			pinfo[id],
			PInfo{
				ID:           id,
				Name:         name,
				Username:     username,
				CommandLine:  cmdline,
				Status:       status,
				CreationTime: ctime,
				Background:   backgrd,
				Foreground:   foregrd,
				IsRunning:    isrunning,
				CPUPercent:   cpuper,
				CPUTimes:     *times,
				Threads:      threads,
				MemPercent:   memper,
				MemInfo:      *meminfo,
				ParentP:      *parent,
				ChildrenP:    children,
			},
		)
	}
	return pinfo, nil
}
