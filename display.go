package rpi

// Display represents a current device display configuration
type Display struct {
	IsOverscan              bool `json:"isOverscan"`
	IsBlanking              bool `json:"isBlanking"`
	IsXscreenSaverInstalled bool `json:"isXscreenSaverInstalled"`
}
