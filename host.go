package rpi

// HostID represents the current system OS specifications.
type HostID struct {
	Hostname             string        `json:"hostname"`
	HostTimeUsage        HostTimeUsage `json:"hostTimeUsage"`
	OS                   string        `json:"os"`              // ex: freebsd, linux
	Platform             string        `json:"platform"`        // ex: ubuntu, linuxmint
	PlatformFamily       string        `json:"platformFamily"`  // ex: debian, rhel
	PlatformVersion      string        `json:"platformVersion"` // version of the complete OS
	KernelVersion        string        `json:"kernelVersion"`   // version of the OS kernel (if available)
	KernelArch           string        `json:"kernelArch"`      // native cpu architecture queried at runtime, as returned by `uname -m` or empty string in case of error
	VirtualizationSystem string        `json:"virtualizationSystem"`
	VirtualizationRole   string        `json:"virtualizationRole"` // guest or host
	HostID               string        `json:"hostid"`             // ex: uuid
}

// HostTimeUsage represents the current system usage time.
type HostTimeUsage struct {
	Uptime   uint64 `json:"uptime"`
	BootTime uint64 `json:"bootTime"`
}

// HostUserUsage represents a current system active terminals.
type HostUserUsage struct {
	User     string `json:"user"`
	Terminal string `json:"terminalType"`
	Host     string `json:"host"`
	Uptime   string `json:"uptime"`
}

// HostLoadAverage represents the current system average system load over a period of time:
// 1, 5, and 15 min.
type HostLoadAverage struct {
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}
