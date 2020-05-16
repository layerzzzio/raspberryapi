package rpi

// Process is
type Process struct {
	ID           int32   `json:"id"`
	Name         string  `json:"name"`
	Username     string  `json:"username,omitempty"`
	CommandLine  string  `json:"commandLine,omitempty"`
	Status       string  `json:"status,omitempty"`
	CreationTime int64   `json:"creationTime,omitempty"`
	Foreground   bool    `json:"foreground,omitempty"`
	Background   bool    `json:"background,omitempty"`
	IsRunning    bool    `json:"isRunning,omitempty"`
	CPUPercent   float64 `json:"cpuPercent"`
	MemPercent   float32 `json:"memPercent"`
	ParentP      int32   `json:"parentPID,omitempty"`
}
