package rpi

// Host represents the current system host specifications.
type Host struct {
	ID                 string  `json:"id"`
	Hostname           string  `json:"hostname"`
	Uptime             uint64  `json:"uptime"`
	BootTime           uint64  `json:"bootTime"`
	OS                 string  `json:"os"`
	Platform           string  `json:"platform"`
	PlatformFamily     string  `json:"platformFamily"`
	PlatformVersion    string  `json:"platformVersion"`
	KernelVersion      string  `json:"kernelVersion"`
	KernelArch         string  `json:"kernelArch"`
	CPU                uint8   `json:"cpus"`
	HyperThreading     bool    `json:"hyperThreading"`
	VCore              uint8   `json:"vcores"`
	CPUUsedPercent     float64 `json:"cpuPercent"`
	VUsedPercent       float64 `json:"virtMemUsedPercent"`
	SUsedPercent       float64 `json:"swapMemUsedPercent"`
	Processes          uint64  `json:"processes"`
	ActiveVirtualUsers uint16  `json:"activeVirtualUsers"`
	Temperature        float32 `json:"temperature"`
}
