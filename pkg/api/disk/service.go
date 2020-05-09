package disk

import (
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/disk"
)

// Service represents all disk application services.
type Service interface {
	List() ([]rpi.Disk, error)
	View(string) (rpi.Disk, error)
}

// Disk represents a disk application service.
type Disk struct {
	dsys DSYS
	m    Metrics
}

// DSYS represents a disk repository service.
type DSYS interface {
	List(map[string]disk.PartitionStat, map[string]*disk.UsageStat) ([]rpi.Disk, error)
	View(string, map[string]disk.PartitionStat, map[string]*disk.UsageStat) (rpi.Disk, error)
}

// Metrics represents the system metrics interface
type Metrics interface {
	DiskStats(bool) (map[string]disk.PartitionStat, map[string]*disk.UsageStat, error)
}

// New creates a Disk application service instance.
func New(dsys DSYS, m Metrics) *Disk {
	return &Disk{dsys: dsys, m: m}
}
