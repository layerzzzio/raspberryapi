package rpi

// Mem represents the current host memory stats.
type Mem struct {
	STotal       uint64  `json:"swapMemTotal"`
	SUsed        uint64  `json:"swapMemUsed"`
	SFree        uint64  `json:"swapMemFree"`
	SUsedPercent float64 `json:"swapMemUsedPercent"`
	VTotal       uint64  `json:"virtMemTotal"`
	VAvailable   uint64  `json:"virtMemAvailable"`
	VUsed        uint64  `json:"virtMemUsed"`
	VUsedPercent float64 `json:"virtMemUsedPercent"`
}
