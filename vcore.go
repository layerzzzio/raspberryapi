package rpi

// VCore represents a current system vCore.
type VCore struct {
	ID     int     `json:"id"`
	Used   float64 `json:"percentUsed"`
	User   float64 `json:"user"`
	System float64 `json:"system"`
	Idle   float64 `json:"idle"`
}
