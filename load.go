package rpi

// load represents the current host load stats.
type Load struct {
	Load1        float64 `json:"load1"`
	Load5        float64 `json:"load5"`
	Load15       float64 `json:"load15"`
	ProcsTotal   int     `json:"procsTotal"`
	ProcsRunning int     `json:"procsRunning"`
	ProcsBlocked int     `json:"procsBlocked"`
}
