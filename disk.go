package rpi

// Disk represents the current system disk.
type Disk struct {
	ID          string       `json:"id"`
	Filesystem  string       `json:"filesystem"`
	Fstype      string       `json:"fstype"`
	Mountpoints []MountPoint `json:"mountpoints"`
}

// MountPoint is
type MountPoint struct {
	Mountpoint        string  `json:"mountpoint"`
	Fstype            string  `json:"fstype"`
	Opts              string  `json:"opts"`
	Total             uint64  `json:"total"`
	Free              uint64  `json:"free"`
	Used              uint64  `json:"used"`
	UsedPercent       float64 `json:"usedPercent"`
	InodesTotal       uint64  `json:"inodesTotal"`
	InodesUsed        uint64  `json:"inodesUsed"`
	InodesFree        uint64  `json:"inodesFree"`
	InodesUsedPercent float64 `json:"inodesUsedPercent"`
}
