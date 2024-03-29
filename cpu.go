package rpi

// CPU represents a current host cpu.
type CPU struct {
	ID        int      `json:"id"`
	Cores     int32    `json:"cores"`
	ModelName string   `json:"modelName"`
	Mhz       float64  `json:"mhz"`
	Stats     CPUStats `json:"stats"`
}

// CPUStats represents a current host cpu stats.
type CPUStats struct {
	Used   float64 `json:"percentUsed"`
	User   float64 `json:"user"`
	System float64 `json:"system"`
	Idle   float64 `json:"idle"`
}
