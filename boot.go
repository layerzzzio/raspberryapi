package rpi

// Boot represents a current device boot behavior
type Boot struct {
	IsWaitForNetwork bool `json:"isWaitForNetwork"`
}
