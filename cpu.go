package rpi

// CPU holds the current system CPU specifications.
type CPU struct {
	ID        int      `json:"id"`
	Cores     int32    `json:"cores"`
	ModelName string   `json:"modelName"`
	Mhz       float64  `json:"mhz"`
	Stats     CPUStats `json:"stats"`
}

// CPUStats is
type CPUStats struct {
	Used      float64 `json:"percentUsed"`
	User      float64 `json:"user"`
	System    float64 `json:"system"`
	Idle      float64 `json:"idle"`
	Nice      float64 `json:"nice"`
	Iowait    float64 `json:"iowait"`
	Irq       float64 `json:"irq"`
	Softirq   float64 `json:"softirq"`
	Steal     float64 `json:"steal"`
	Guest     float64 `json:"guest"`
	GuestNice float64 `json:"guestNice"`
}
