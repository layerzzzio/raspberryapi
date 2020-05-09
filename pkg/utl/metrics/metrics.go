package metrics

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// TODO: to test this method by simulating different OS scenarios in a Docker container (raspbian/strech)

// Service is
type Service struct{}

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
func (s Service) DiskStats(all bool) (map[string]disk.PartitionStat, map[string]*disk.UsageStat, error) {
	listDisk := make(map[string]disk.PartitionStat)
	stats := make(map[string]*disk.UsageStat)

	disks, err := disk.Partitions(all)
	if err != nil {
		return nil, nil, err
	}

	for i := range disks {
		listDisk[disks[i].Device] = disks[i]
	}

	for i := range disks {
		usage, err := disk.Usage(disks[i].Mountpoint)
		if err != nil {
			return nil, nil, err
		}
		stats[disks[i].Device] = usage
	}

	return listDisk, stats, nil
}
