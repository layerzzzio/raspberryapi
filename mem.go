package rpi

// MemoryUsage represents the current system swap memory usage.
type MemoryUsage struct {
	STotal       uint64  `json:"sMemTotal"`
	SUsed        uint64  `json:"sMemUsed"`
	SFree        uint64  `json:"sMemFree"`
	SUsedPercent float64 `json:"sMemUsedPercent"`
	VTotal       uint64  `json:"vMemTotal"`
	VAvailable   uint64  `json:"vMemAvailable"`
	VUsed        uint64  `json:"vMemUsed"`
	VUsedPercent float64 `json:"vMemUsedPercent"`
}
